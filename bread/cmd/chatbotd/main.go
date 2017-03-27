// Command chatbotd is the HTTP server that receives and handles chat messages.
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	heroku "github.com/cyberdelia/heroku-go/v3"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sr/operator/hipchat"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"git.dev.pardot.com/Pardot/infrastructure/bread"
	"git.dev.pardot.com/Pardot/infrastructure/bread/hipchat"
	"git.dev.pardot.com/Pardot/infrastructure/bread/pb"
)

const (
	grpcAddress         = ":8443"
	grpcDialTimeout     = 10 * time.Second
	defaultProtoPackage = "bread"
)

var (
	port               = flag.String("port", "", "Listening port of the HTTP server. Defaults to $PORT.")
	timeout            = flag.Duration("command-timeout", 10*time.Minute, "")
	hipchatOAuthID     = flag.String("hipchat-oauth-id", "", "")
	hipchatOAuthSecret = flag.String("hipchat-oauth-secret", "", "")
	herokuAPIUsername  = flag.String("heroku-api-username", "", "")
	herokuAPIPassword  = flag.String("heroku-api-password", "", "")
	canoeAPIURL        = flag.String("canoe-api-url", "https://canoe.dev.pardot.com", "")
	canoeAPIKey        = flag.String("canoe-api-key", "", "Canoe API key")

	hipchatAddonConfig = &breadhipchat.AddonConfig{}
	ldapConfig         = &bread.LDAPConfig{}
)

func init() {
	flag.StringVar(&hipchatAddonConfig.Name, "hipchat-addon-name", "", "")
	flag.StringVar(&hipchatAddonConfig.Key, "hipchat-addon-key", "", "")
	flag.StringVar(&hipchatAddonConfig.Homepage, "hipchat-addon-homepage", "", "")
	flag.StringVar(&hipchatAddonConfig.URL, "hipchat-addon-url", "", "")
	flag.StringVar(&hipchatAddonConfig.WebhookURL, "hipchat-addon-webhook-url", "", "")
	flag.StringVar(&hipchatAddonConfig.AvatarURL, "hipchat-addon-avatar-url", "", "")
	flag.StringVar(&hipchatAddonConfig.HerokuApp, "hipchat-addon-heroku-app", "", "")
	flag.StringVar(&ldapConfig.Addr, "ldap-address", "localhost:389", "Address of the LDAP server used to authenticate and authorize commands")
	flag.StringVar(&ldapConfig.Base, "ldap-base", bread.LDAPBase, "LDAP Base DN")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "chatbotd: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	flag.VisitAll(func(f *flag.Flag) {
		k := strings.ToUpper(strings.Replace(f.Name, "-", "_", -1))
		if v := os.Getenv(k); v != "" {
			if err := f.Value.Set(v); err != nil {
				panic(err)
			}
		}
	})
	flag.Parse()
	if *port == "" {
		return errors.New("required flag missing: port")
	}
	if *canoeAPIURL == "" {
		return errors.New("required flag missing: canoe-api-url")
	}
	parsedCanoeAPIURL, err := url.Parse(*canoeAPIURL)
	if err != nil {
		return fmt.Errorf("required flag canoe-url is invalid: %s", err)
	}
	if *canoeAPIKey == "" {
		return errors.New("required flag missing: canoe-api-key")
	}
	if v, ok := os.LookupEnv("LDAP_CA_CERT"); ok {
		ldapConfig.CACert = []byte(v)
	}

	logger := log.New(os.Stderr, "chatbotd: ", 0)

	client, err := operatorhipchat.NewClient(
		context.TODO(),
		&operatorhipchat.ClientConfig{
			Hostname: bread.HipchatHost,
			Scopes:   breadhipchat.DefaultScopes,
			Credentials: &operatorhipchat.ClientCredentials{
				ID:     *hipchatOAuthID,
				Secret: *hipchatOAuthSecret,
			},
		},
	)
	if err != nil {
		return err
	}

	authorizer, err := bread.NewAuthorizer(
		ldapConfig,
		bread.NewCanoeClient(parsedCanoeAPIURL, *canoeAPIKey),
		bread.ACL,
	)
	if err != nil {
		return err
	}
	interceptor := &bread.Interceptor{Authorizer: authorizer}

	errC := make(chan error, 1)
	// Start the gRPC server and establish a client connection to it.
	grpcListener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.UnaryServerInterceptor))
	breadpb.RegisterPingerServer(grpcServer, &pingerServer{client})
	go func() {
		errC <- grpcServer.Serve(grpcListener)
	}()

	conn, err := grpc.Dial(
		grpcAddress,
		grpc.WithBlock(),
		grpc.WithTimeout(grpcDialTimeout),
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	// Channel used to by the event HTTP server to send chat commands to the workers.
	chatCommands := make(chan *bread.ChatCommand, 100)

	// All chat messages are processed through these functions
	messageHandlers := []bread.ChatMessageHandler{
		bread.LogHandler(logger),
		bread.ChatCommandHandler(defaultProtoPackage, chatCommands),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/_ping", bread.NewPingHandler(nil))

	// Either mount the Hipchat addon server or event servers, depending on
	// whether OAuth client credentials are available or not.
	if *hipchatOAuthID == "" && *hipchatOAuthSecret == "" {
		if *herokuAPIUsername == "" {
			return fmt.Errorf("required flag missing: heroku-api-username")
		}
		if *herokuAPIPassword == "" {
			return fmt.Errorf("required flag missing: heroku-api-password")
		}

		handler, err := breadhipchat.AddonHandler(
			heroku.NewService(&http.Client{
				Transport: &heroku.Transport{
					Username: *herokuAPIUsername,
					Password: *herokuAPIPassword,
				},
			}),
			hipchatAddonConfig,
		)
		if err != nil {
			return err
		}
		mux.HandleFunc("/hipchat/addon", handler)
	} else {
		webhookURL, err := url.Parse(hipchatAddonConfig.WebhookURL)
		if err != nil {
			return err
		}

		handler, err := breadhipchat.EventHandler(
			client,
			&clientcredentials.Config{
				ClientID:     *hipchatOAuthID,
				ClientSecret: *hipchatOAuthSecret,
				TokenURL:     fmt.Sprintf("%s/v2/oauth/token", bread.HipchatHost),
				Scopes:       breadhipchat.DefaultScopes,
			},
			func(msg *bread.ChatMessage) error {
				var fail bool
				for _, handler := range messageHandlers {
					if err := handler(msg); err != nil {
						fail = true
						logger.Printf("error handling message: %s", err)
					}
				}
				if fail {
					return errors.New("one of the handler failed. check the logs for details")
				}
				return nil
			},
		)
		if err != nil {
			return err
		}
		mux.HandleFunc(webhookURL.Path, handler)
	}

	// Start chat command workers.
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for cmd := range chatCommands {
				if err := bread.HandleChatRPCCommand(client, Invoker, *timeout, conn, cmd); err != nil {
					logger.Printf("command handler error: %s", err)
				}
			}
			wg.Done()
		}()
	}

	// Start the HTTP server.
	httpServer := &http.Server{
		Addr:     fmt.Sprintf(":%s", *port),
		ErrorLog: logger,
		Handler:  mux,
	}
	go func() {
		errC <- httpServer.ListenAndServe()
	}()

	var services []string
	for s := range grpcServer.GetServiceInfo() {
		services = append(services, s)
	}
	logger.Printf("server listening on %s. registered gRPC services: %s", httpServer.Addr, strings.Join(services, " "))

	// Gracefully shutdown the HTTP and gRPC servers, and chat command handler
	// goroutines on SIGINT.
	stopC := make(chan os.Signal)
	signal.Notify(stopC, os.Interrupt)

	select {
	case <-stopC:
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			return err
		}
		close(chatCommands)
		wg.Wait()
		// This blocks until all RPCs are finished.
		grpcServer.GracefulStop()
		return nil
	case err := <-errC:
		return err
	}
}

// Invoker is a generated function TODO(sr)
func Invoker(ctx context.Context, conn *grpc.ClientConn, cmd *bread.ChatCommand) error {
	if cmd.Package == "bread" {
		if cmd.Service == "Pinger" {
			if cmd.Method == "Ping" {
				_, err := breadpb.NewPingerClient(conn).Ping(ctx, &empty.Empty{})
				return err
			}
		}
	}
	return fmt.Errorf("unhandleable command: %+v", cmd)
}

// TODO(sr) Move this out of the main.
type pingerServer struct {
	hipchat operatorhipchat.Client
}

func (s *pingerServer) Ping(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	var (
		email  string
		roomID int64
	)
	// TODO(sr) Abstract this into a chatbot.Reply(context.Context) function or similar
	if md, ok := metadata.FromContext(ctx); ok {
		if val, ok := md["user_email"]; ok {
			if len(val) == 1 {
				email = val[0]
			}
		}
		if val, ok := md["hipchat_room_id"]; ok {
			if len(val) == 1 {
				if i, err := strconv.Atoi(val[0]); err == nil {
					roomID = int64(i)
				}
			}
		}
	}
	if email == "" {
		return nil, errors.New("no user email found in request")
	}
	if roomID == 0 {
		return nil, errors.New("no hipchat room ID found in request")
	}
	return &empty.Empty{},
		s.hipchat.SendRoomNotification(ctx, &operatorhipchat.RoomNotification{
			Message:       fmt.Sprintf(`PONG <a href="mailto:%s">%s</a>`, email, email),
			RoomID:        roomID,
			MessageFormat: "html",
			MessageOptions: &operatorhipchat.MessageOptions{
				Color: "gray",
				From:  "pinger.Ping",
			},
		})
}