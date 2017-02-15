// Code generated by protoc-gen-operatord
package main

import (
	"fmt"

	"github.com/sr/operator"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	breadpb "git.dev.pardot.com/Pardot/bread/pb"
)

func invoker(ctx context.Context, conn *grpc.ClientConn, req *operator.Request, pkg string) error {
	if req.Call.Service == fmt.Sprintf("%s.Deploy", pkg) {
		if req.Call.Method == "ListTargets" {
			client := breadpb.NewDeployClient(conn)
			_, err := client.ListTargets(
				ctx,
				&breadpb.ListTargetsRequest{
					Request: req,
				},
			)
			return err
		}
		if req.Call.Method == "ListBuilds" {
			client := breadpb.NewDeployClient(conn)
			_, err := client.ListBuilds(
				ctx,
				&breadpb.ListBuildsRequest{
					Request: req,
					Target:  req.Call.Args["target"],
					Branch:  req.Call.Args["branch"],
				},
			)
			return err
		}
		if req.Call.Method == "Trigger" {
			client := breadpb.NewDeployClient(conn)
			_, err := client.Trigger(
				ctx,
				&breadpb.TriggerRequest{
					Request: req,
					Target:  req.Call.Args["target"],
					Build:   req.Call.Args["build"],
					Branch:  req.Call.Args["branch"],
				},
			)
			return err
		}
	}
	if req.Call.Service == fmt.Sprintf("%s.Issues", pkg) {
		if req.Call.Method == "Mine" {
			client := breadpb.NewIssuesClient(conn)
			_, err := client.Mine(
				ctx,
				&breadpb.MyIssuesRequest{
					Request: req,
				},
			)
			return err
		}
	}
	if req.Call.Service == fmt.Sprintf("%s.Ping", pkg) {
		if req.Call.Method == "Ping" {
			client := breadpb.NewPingClient(conn)
			_, err := client.Ping(
				ctx,
				&breadpb.PingRequest{
					Request: req,
				},
			)
			return err
		}
		if req.Call.Method == "SlowLoris" {
			client := breadpb.NewPingClient(conn)
			_, err := client.SlowLoris(
				ctx,
				&breadpb.SlowLorisRequest{
					Request: req,
					Wait:    req.Call.Args["wait"],
				},
			)
			return err
		}
	}
	return fmt.Errorf("no such service: `%s %s`", req.Call.Service, req.Call.Method)
}
