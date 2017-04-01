// Code generated by protoc-gen-bread. DO NOT EDIT.

package breadgen

import (
	"errors"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"git.dev.pardot.com/Pardot/infrastructure/bread/api"
	breadpb "git.dev.pardot.com/Pardot/infrastructure/bread/generated/pb"
)

func ChatCommandGRPCInvoker(ctx context.Context, conn *grpc.ClientConn, cmd *breadapi.ChatCommand) error {
	if conn == nil {
		return errors.New("required argument is nil: conn")
	}
	if cmd == nil {
		return errors.New("required argument is nil: cmd")
	}
	if cmd.Call == nil {
		return errors.New("required cmd struct field is nil: Call")
	}

	if cmd.Call.Package == "breadpb" && cmd.Call.Service == "Deploy" {
		if cmd.Call.Method == "ListTargets" {
			_, err := breadpb.NewDeployClient(conn).ListTargets(
				ctx,
				&breadpb.ListTargetsRequest{},
			)
			return err
		}
		if cmd.Call.Method == "ListBuilds" {
			_, err := breadpb.NewDeployClient(conn).ListBuilds(
				ctx,
				&breadpb.ListBuildsRequest{
					Target: cmd.Args["target"],
					Branch: cmd.Args["branch"],
				},
			)
			return err
		}
		if cmd.Call.Method == "Trigger" {
			_, err := breadpb.NewDeployClient(conn).Trigger(
				ctx,
				&breadpb.TriggerRequest{
					Target: cmd.Args["target"],
					Build:  cmd.Args["build"],
					Branch: cmd.Args["branch"],
				},
			)
			return err
		}
	}

	if cmd.Call.Package == "breadpb" && cmd.Call.Service == "Pinger" {
		if cmd.Call.Method == "Ping" {
			_, err := breadpb.NewPingerClient(conn).Ping(
				ctx,
				&breadpb.PingRequest{},
			)
			return err
		}
	}

	if cmd.Call.Package == "breadpb" && cmd.Call.Service == "Tickets" {
		if cmd.Call.Method == "Mine" {
			_, err := breadpb.NewTicketsClient(conn).Mine(
				ctx,
				&breadpb.TicketRequest{
					IncludeResolved: cmd.Args["include_resolved"],
				},
			)
			return err
		}
		if cmd.Call.Method == "SprintStatus" {
			_, err := breadpb.NewTicketsClient(conn).SprintStatus(
				ctx,
				&breadpb.TicketRequest{
					IncludeResolved: cmd.Args["include_resolved"],
				},
			)
			return err
		}
	}
	return fmt.Errorf("unhandleable command: %+v", cmd)
}
