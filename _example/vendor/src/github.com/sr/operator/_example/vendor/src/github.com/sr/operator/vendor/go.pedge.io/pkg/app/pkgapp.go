package pkgapp // import "go.pedge.io/pkg/app"

import (
	"github.com/rcrowley/go-metrics"
	"go.pedge.io/env"
	"go.pedge.io/pkg/log"
	"go.pedge.io/pkg/metrics"
)

// AppEnv is the struct that represents the environment variables used by an app.
type AppEnv struct {
	// the app name.
	// default is app.
	AppName string `env:"APP_NAME,default=app"`
	// See pkglog for the log environment variables.
	LogEnv pkglog.Env
	// See pkgmetrics for the metrics environment variables.
	MetricsEnv pkgmetrics.Env
}

// GetAppEnv gets the AppEnv from the environment.
func GetAppEnv() (AppEnv, error) {
	appEnv := AppEnv{}
	if err := env.Populate(&appEnv); err != nil {
		return AppEnv{}, err
	}
	return appEnv, nil
}

// AppOptions are options for a new app.
type AppOptions struct {
	// The registry to use.
	// Can be nil
	MetricsRegistry metrics.Registry
}

// SetupAppEnv does the setup for AppEnv.
func SetupAppEnv(appEnv AppEnv) (AppOptions, error) {
	if err := pkglog.SetupLogging(appEnv.AppName, appEnv.LogEnv); err != nil {
		return AppOptions{}, err
	}
	registry, err := pkgmetrics.SetupMetrics(appEnv.AppName, appEnv.MetricsEnv)
	if err != nil {
		return AppOptions{}, err
	}
	return AppOptions{
		MetricsRegistry: registry,
	}, nil
}

// GetAndSetupAppEnv does GetAppEnv then SetupAppEnv.
func GetAndSetupAppEnv() (AppOptions, error) {
	appEnv, err := GetAppEnv()
	if err != nil {
		return AppOptions{}, err
	}
	return SetupAppEnv(appEnv)
}
