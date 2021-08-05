package server

import (
	"log"
	"net"

	"github.com/juninhoitabh/marketanalyzer-talibservice/infrastructure/grpc/pb"
	"github.com/juninhoitabh/marketanalyzer-talibservice/infrastructure/grpc/service"
	"github.com/juninhoitabh/marketanalyzer-talibservice/useCases"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	ProcessIndicators useCases.UseCaseIndicators
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (g GRPCServer) Serve() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Could not listen tcp port")
	}
	indicatorsService := service.NewIndicatorsService()
	indicatorsService.GenereteIndicatorsUseCase = g.ProcessIndicators
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterIndicatorsServiceServer(grpcServer, indicatorsService)

	grpcServer.Serve(lis)
}
