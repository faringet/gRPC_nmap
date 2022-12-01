package server

import (
	"context"
	"fmt"
	"gRPC_nmap/config"
	pb "gRPC_nmap/gen/netvuln"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.NetVulnServiceServer
}

func (s *server) CheckVuln(ctx context.Context, req *pb.CheckVulnRequest) (*pb.CheckVulnResponse, error) {
	return &pb.CheckVulnResponse{
		Results: []*pb.TargetResult{
			{Target: "test"},
		},
	}, nil

}

func Start(config config.Config) error {
	listener, err := net.Listen("tcp", config.Addr)
	if err != nil {
		return fmt.Errorf("не смогли забиндить порт: %w", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterNetVulnServiceServer(grpcServer, &server{})

	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("не смогли запустить сервер: %w", err)
	}
	return nil
}
