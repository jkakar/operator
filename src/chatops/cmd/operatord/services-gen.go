// Code generated by protoc-gen-operatord
package main

import (
	"chatops/ping"

	"github.com/sr/operator"
	"go.pedge.io/env"
	"google.golang.org/grpc"
)

func registerServices(
	server *grpc.Server,
	logger operator.Logger,
	instrumenter operator.Instrumenter,
	authorizer operator.Authorizer,
) {

	pingerConfig := &pinger.Env{}
	if err := env.Populate(pingerConfig); err != nil {
		logError(logger, "pinger", err)
	}
	pingerServer, err := pinger.NewAPIServer(pingerConfig)
	if err != nil {
		logError(logger, "pinger", err)
	}
	intercepted := &interceptedpingerPinger{
		authorizer,
		instrumenter,
		pingerServer,
	}
	pinger.RegisterPingerServer(server, intercepted)
	logger.Info(&operator.ServiceRegistered{&operator.Service{Name: "pinger"}})

}

func logError(logger operator.Logger, service string, err error) {
	logger.Error(&operator.ServiceStartupError{
		Service: &operator.Service{
			Name: service,
		},
		Message: err.Error(),
	})
}
