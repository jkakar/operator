// Code generated by protoc-gen-operatord
package main

import (
	"time"

	"github.com/sr/operator"
	"golang.org/x/net/context"

	servicepkg "chatops/ping"
)

type interceptedpingerPinger struct {
	authorizer   operator.Authorizer
	instrumenter operator.Instrumenter
	server       servicepkg.PingerServer
}

// Ping intercepts the Pinger.Ping method.
func (a *interceptedpingerPinger) Ping(
	ctx context.Context,
	request *servicepkg.PingRequest,
) (response *servicepkg.PingResponse, err error) {
	defer func(start time.Time) {
		a.instrumenter.Instrument(
			operator.NewRequest(
				request.Source,
				"pinger",
				"Ping",
				"PingRequest",
				"PingResponse",
				err,
				start,
			),
		)
	}(time.Now())
	if err := a.authorizer.Authorize(request.Source); err != nil {
		return nil, err
	}
	return a.server.Ping(ctx, request)
}
