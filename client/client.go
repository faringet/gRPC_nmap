package client

import (
	"context"
	"fmt"
	"gRPC_nmap/config"
	pb "gRPC_nmap/gen/netvuln"
	"google.golang.org/grpc"
)

func SendRequest(config config.Config) error {
	conn, err := grpc.Dial(config.Addr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("не смогли соедениться: %w", err)
	}

	clent := pb.NewNetVulnServiceClient(conn)

	response, err := clent.CheckVuln(context.Background(), &pb.CheckVulnRequest{})
	if err != nil {
		return fmt.Errorf("не смогли сделать запрос: %w", err)
	}
	fmt.Printf("%#v\n", response.Results[0])
	return nil

}
