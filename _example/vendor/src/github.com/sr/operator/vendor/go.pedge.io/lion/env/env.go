/*
Package envlion provides simple utilities to setup lion from the environment.
*/
package envlion // import "go.pedge.io/lion/env"

import (
	"fmt"
	"log/syslog"
	"os"
	"path/filepath"
	"strings"

	"go.pedge.io/env"
	"go.pedge.io/lion"
	"go.pedge.io/lion/syslog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Env defines a struct for environment variables that can be parsed with go.pedge.io/env.
type Env struct {
	// The log app name, will default to app if not set.
	LogAppName string `env:"LOG_APP_NAME,default=app"`
	// The level to log at, must be one of DEBUG, INFO, WARN, ERROR, FATAL, PANIC.
	LogLevel string `env:"LOG_LEVEL"`
	// LogDisableStderr says to disable logging to stderr.
	LogDisableStderr bool `env:"LOG_DISABLE_STDERR"`
	// The directory to write rotating logs to.
	// If not set and SyslogNetwork and SyslogAddress not set, logs will be to stderr.
	LogDirPath string `env:"LOG_DIR_PATH"`
	// The syslog network, either udp or tcp.
	// Must be set with SyslogAddress.
	// If not set and LogDir not set, logs will be to stderr.
	SyslogNetwork string `env:"SYSLOG_NETWORK"`
	// The syslog host:port.
	// Must be set with SyslogNetwork.
	// If not set and LogDir not set, logs will be to stderr.
	SyslogAddress string `env:"SYSLOG_ADDRESS"`
}

// Setup gets the Env from the environment, and then calls SetupEnv.
func Setup() error {
	appEnv := Env{}
	if err := env.Populate(&appEnv); err != nil {
		return err
	}
	return SetupEnv(appEnv)
}

// SetupEnv sets up logging for the given Env.
func SetupEnv(env Env) error {
	var pushers []lion.Pusher
	logAppName := env.LogAppName
	if logAppName == "" {
		logAppName = "app"
	}
	if !env.LogDisableStderr {
		pushers = append(pushers, lion.NewTextWritePusher(os.Stderr))
	}
	if env.LogDirPath != "" {
		pushers = append(pushers, newLogDirPusher(env.LogDirPath, logAppName))
	}
	if env.SyslogNetwork != "" && env.SyslogAddress != "" {
		pusher, err := newSyslogPusher(env.SyslogNetwork, env.SyslogAddress, logAppName)
		if err != nil {
			return err
		}
		pushers = append(pushers, pusher)
	}
	switch len(pushers) {
	case 0:
		lion.SetLogger(lion.DiscardLogger)
	case 1:
		lion.SetLogger(lion.NewLogger(pushers[0]))
	default:
		lion.SetLogger(lion.NewLogger(lion.NewMultiPusher(pushers...)))
	}
	lion.RedirectStdLogger()
	if env.LogLevel != "" {
		level, err := lion.NameToLevel(strings.ToUpper(env.LogLevel))
		if err != nil {
			return err
		}
		lion.SetLevel(level)
	}
	return nil
}

func newLogDirPusher(logDirPath string, logAppName string) lion.Pusher {
	return lion.NewTextWritePusher(
		&lumberjack.Logger{
			Filename:   filepath.Join(logDirPath, fmt.Sprintf("%s.log", logAppName)),
			MaxBackups: 3,
		},
	)
}

func newSyslogPusher(syslogNetwork string, syslogAddress string, logAppName string) (lion.Pusher, error) {
	writer, err := syslog.Dial(
		syslogNetwork,
		syslogAddress,
		syslog.LOG_INFO,
		logAppName,
	)
	if err != nil {
		return nil, err
	}
	return sysloglion.NewPusher(writer), nil
}
