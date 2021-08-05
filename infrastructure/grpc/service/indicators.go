package service

import (
	"context"
	"encoding/json"

	"github.com/juninhoitabh/marketanalyzer-talibservice/dto"
	"github.com/juninhoitabh/marketanalyzer-talibservice/infrastructure/grpc/pb"
	"github.com/juninhoitabh/marketanalyzer-talibservice/useCases"
)

type IndicatorsService struct {
	GenereteIndicatorsUseCase useCases.UseCaseIndicators
	pb.UnimplementedIndicatorsServiceServer
}

func NewIndicatorsService() *IndicatorsService {
	return &IndicatorsService{}
}

func (c *IndicatorsService) Ema(ctx context.Context, in *pb.StandardRequest) (*pb.StandardResult, error) {
	p := []dto.PriceHistory{}
	json.Unmarshal([]byte(in.GetData()), &p)
	resultEma, err := c.GenereteIndicatorsUseCase.Ema(p, in.GetSource(), in.GetLength())

	if err == nil {
		return &pb.StandardResult{
			Data: resultEma,
		}, nil
	}

	return &pb.StandardResult{
		Data: "",
	}, err
}
