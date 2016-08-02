// Code generated by protoc-gen-operatorcmd
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	bamboo "bread/bamboo"
	deploy "bread/deploy"
	pinger "bread/ping"
	"github.com/sr/operator"
	"golang.org/x/net/context"
)

const programName = "operator"

var cmd = operator.NewCommand(
	programName,
	[]operator.ServiceCommand{
		{
			Name:     "bamboo",
			Synopsis: `Undocumented.`,
			Methods: []operator.MethodCommand{
				{
					Name:     "list-builds",
					Synopsis: `Undocumented.`,
					Flags: []*flag.Flag{
						{
							Name:  "plan",
							Usage: "Undocumented.",
						},
					},
					Run: func(ctx *operator.CommandContext) (string, error) {
						plan := ctx.Flags.String("plan", "", "")
						if err := ctx.Flags.Parse(ctx.Args); err != nil {
							return "", err
						}
						conn, err := ctx.GetConn()
						if err != nil {
							return "", err
						}
						defer conn.Close()
						client := bamboo.NewBambooClient(conn)
						response, err := client.ListBuilds(
							context.Background(),
							&bamboo.ListBuildsRequest{
								Source: ctx.Source,
								Plan:   *plan,
							},
						)
						if err != nil {
							return "", err
						}
						return response.Output.PlainText, nil
					},
				},
			},
		},

		{
			Name:     "deploy",
			Synopsis: `Undocumented.`,
			Methods: []operator.MethodCommand{
				{
					Name:     "list-apps",
					Synopsis: `Undocumented.`,
					Flags:    []*flag.Flag{},
					Run: func(ctx *operator.CommandContext) (string, error) {
						if err := ctx.Flags.Parse(ctx.Args); err != nil {
							return "", err
						}
						conn, err := ctx.GetConn()
						if err != nil {
							return "", err
						}
						defer conn.Close()
						client := deploy.NewDeployClient(conn)
						response, err := client.ListApps(
							context.Background(),
							&deploy.ListAppsRequest{
								Source: ctx.Source,
							},
						)
						if err != nil {
							return "", err
						}
						return response.Output.PlainText, nil
					},
				},
				{
					Name:     "trigger",
					Synopsis: `Undocumented.`,
					Flags: []*flag.Flag{
						{
							Name:  "app",
							Usage: "Undocumented.",
						},
						{
							Name:  "build",
							Usage: "Undocumented.",
						},
					},
					Run: func(ctx *operator.CommandContext) (string, error) {
						app := ctx.Flags.String("app", "", "")
						build := ctx.Flags.String("build", "", "")
						if err := ctx.Flags.Parse(ctx.Args); err != nil {
							return "", err
						}
						conn, err := ctx.GetConn()
						if err != nil {
							return "", err
						}
						defer conn.Close()
						client := deploy.NewDeployClient(conn)
						response, err := client.Trigger(
							context.Background(),
							&deploy.TriggerRequest{
								Source: ctx.Source,
								App:    *app,
								Build:  *build,
							},
						)
						if err != nil {
							return "", err
						}
						return response.Output.PlainText, nil
					},
				},
			},
		},

		{
			Name:     "pinger",
			Synopsis: `Undocumented.`,
			Methods: []operator.MethodCommand{
				{
					Name:     "ping",
					Synopsis: `Undocumented.`,
					Flags:    []*flag.Flag{},
					Run: func(ctx *operator.CommandContext) (string, error) {
						if err := ctx.Flags.Parse(ctx.Args); err != nil {
							return "", err
						}
						conn, err := ctx.GetConn()
						if err != nil {
							return "", err
						}
						defer conn.Close()
						client := pinger.NewPingerClient(conn)
						response, err := client.Ping(
							context.Background(),
							&pinger.PingRequest{
								Source: ctx.Source,
							},
						)
						if err != nil {
							return "", err
						}
						return response.Output.PlainText, nil
					},
				},
			},
		},
	},
)

func main() {
	status, output := cmd.Run(os.Args)
	if status != 0 {
		if _, err := fmt.Fprintf(os.Stderr, "%s: %s\n", programName, output); err != nil {
			panic(err)
		}
	} else {
		if _, err := io.WriteString(os.Stdout, output+"\n"); err != nil {
			panic(err)
		}
	}
	os.Exit(status)
}
