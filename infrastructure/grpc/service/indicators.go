package service

import (
	"context"

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
	resultEma, err := c.GenereteIndicatorsUseCase.Ema(in.GetData(), in.GetLength(), in.GetInitial())

	if err == nil {
		return &pb.StandardResult{
			Data: resultEma,
		}, nil
	}

	return &pb.StandardResult{
		Data: []float64{},
	}, err
}

func (c *IndicatorsService) Sma(ctx context.Context, in *pb.StandardRequest) (*pb.StandardResult, error) {
	resultSma, err := c.GenereteIndicatorsUseCase.Sma(in.GetData(), in.GetLength(), in.GetInitial())

	if err == nil {
		return &pb.StandardResult{
			Data: resultSma,
		}, nil
	}

	return &pb.StandardResult{
		Data: []float64{},
	}, err
}

func (c *IndicatorsService) Macd(ctx context.Context, in *pb.MacdRequest) (*pb.MacdResult, error) {
	resultMacd, err := c.GenereteIndicatorsUseCase.Macd(in.GetData(), in.GetFastPeriod(), in.GetSlowPeriod(), in.GetSignalPeriod(), in.GetInitial())

	macdReturn := []*pb.MacdReturn{}
	if err == nil {
		for i := 0; i < len(resultMacd); i++ {
			mReturn := &pb.MacdReturn{
				FastLength:   resultMacd[i].FastLength,
				Macd:         resultMacd[i].Macd,
				SignalLength: resultMacd[i].SignalLength,
			}
			macdReturn = append(macdReturn, mReturn)
		}

		return &pb.MacdResult{
			Data: macdReturn,
		}, nil
	}

	return &pb.MacdResult{
		Data: macdReturn,
	}, err
}

func (c *IndicatorsService) StochOscillator(ctx context.Context, in *pb.StochRequest) (*pb.StochResult, error) {
	resultStoch, err := c.GenereteIndicatorsUseCase.StochOscillator(in.GetDataHigh(), in.GetDataLow(), in.GetDataClose(), in.GetLength(), in.GetSmoothK(), in.GetSmoothD(), in.GetMAUse(), in.GetInitial())

	stochReturn := []*pb.StochReturn{}
	if err == nil {
		for i := 0; i < len(resultStoch); i++ {
			sReturn := &pb.StochReturn{
				Smoothk: resultStoch[i].SmoothK,
				Smoothd: resultStoch[i].SmoothD,
			}
			stochReturn = append(stochReturn, sReturn)
		}

		return &pb.StochResult{
			Data: stochReturn,
		}, nil
	}

	return &pb.StochResult{
		Data: stochReturn,
	}, err
}

func (c *IndicatorsService) Cci(ctx context.Context, in *pb.StandardDatasRequest) (*pb.StandardResult, error) {
	resultCci, err := c.GenereteIndicatorsUseCase.Cci(in.GetDataHigh(), in.GetDataLow(), in.GetDataClose(), in.GetLength(), in.GetInitial())

	if err == nil {
		return &pb.StandardResult{
			Data: resultCci,
		}, nil
	}

	return &pb.StandardResult{
		Data: []float64{},
	}, err
}

func (c *IndicatorsService) Atr(ctx context.Context, in *pb.StandardDatasRequest) (*pb.StandardResult, error) {
	resultAtr, err := c.GenereteIndicatorsUseCase.Atr(in.GetDataHigh(), in.GetDataLow(), in.GetDataClose(), in.GetLength(), in.GetInitial())

	if err == nil {
		return &pb.StandardResult{
			Data: resultAtr,
		}, nil
	}

	return &pb.StandardResult{
		Data: []float64{},
	}, err
}

func (c *IndicatorsService) Sar(ctx context.Context, in *pb.StandardDatasOnlyRequest) (*pb.StandardResult, error) {
	resultSar, err := c.GenereteIndicatorsUseCase.Sar(in.GetDataHigh(), in.GetDataLow(), in.GetIncrement(), in.GetMaximum(), in.GetInitial())

	if err == nil {
		return &pb.StandardResult{
			Data: resultSar,
		}, nil
	}

	return &pb.StandardResult{
		Data: []float64{},
	}, err
}

func (c *IndicatorsService) Dmi(ctx context.Context, in *pb.StandardDatasRequest) (*pb.DmiResult, error) {
	resultDmi, err := c.GenereteIndicatorsUseCase.Dmi(in.GetDataHigh(), in.GetDataLow(), in.GetDataClose(), in.GetLength(), in.GetInitial())

	dmiReturn := []*pb.DmiReturn{}
	if err == nil {
		for i := 0; i < len(resultDmi); i++ {
			dReturn := &pb.DmiReturn{
				Adx:    resultDmi[i].Adx,
				Plusd:  resultDmi[i].PlusD,
				Minusd: resultDmi[i].MinusD,
			}
			dmiReturn = append(dmiReturn, dReturn)
		}

		return &pb.DmiResult{
			Data: dmiReturn,
		}, nil
	}

	return &pb.DmiResult{
		Data: dmiReturn,
	}, err
}

func (c *IndicatorsService) Adx(ctx context.Context, in *pb.StandardDatasRequest) (*pb.StandardResult, error) {
	resultAdx, err := c.GenereteIndicatorsUseCase.Adx(in.GetDataHigh(), in.GetDataLow(), in.GetDataClose(), in.GetLength(), in.GetInitial())

	if err == nil {
		return &pb.StandardResult{
			Data: resultAdx,
		}, nil
	}

	return &pb.StandardResult{
		Data: []float64{},
	}, err
}
