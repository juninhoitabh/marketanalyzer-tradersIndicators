package main

import (
	"fmt"

	"github.com/juninhoitabh/marketanalyzer-talibservice/infrastructure/grpc/server"
	"github.com/juninhoitabh/marketanalyzer-talibservice/infrastructure/repository"
	"github.com/juninhoitabh/marketanalyzer-talibservice/useCases"
)

func main() {
	processIndicatorsUseCase := useCases.NewUseCaseIndicators(repository.NewIndicatorsRepository())
	serveGrpc(processIndicatorsUseCase)
}

func serveGrpc(processIndicatorsUseCase useCases.UseCaseIndicators) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessIndicators = processIndicatorsUseCase
	fmt.Println("Run gRPC Server")
	grpcServer.Serve()
}
