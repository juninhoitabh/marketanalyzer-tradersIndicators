package service

import (
	"context"

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
	for _, v := range in.GetData() {
		priceHistoryDto := dto.PriceHistory{
			TimeHistory: v.GetTimeHistory(),
			Open:        v.GetOpen(),
			High:        v.GetHigh(),
			Low:         v.GetLow(),
			Close:       v.GetClose(),
			Volume:      v.GetVolume(),
		}
		p = append(p, priceHistoryDto)
	}
	resultEma, err := c.GenereteIndicatorsUseCase.Ema(p, in.GetSource(), in.GetLength())

	if err == nil {
		return &pb.StandardResult{
			Data: resultEma,
		}, nil
	}

	return &pb.StandardResult{
		Data: []float64{},
	}, err
}
