package useCases

import (
	"github.com/juninhoitabh/marketanalyzer-talibservice/domain"
	"github.com/juninhoitabh/marketanalyzer-talibservice/dto"
)

type UseCaseIndicators struct {
	IndicatorsRepository domain.Indicators
}

func NewUseCaseIndicators(indicatorsRepository domain.Indicators) UseCaseIndicators {
	return UseCaseIndicators{IndicatorsRepository: indicatorsRepository}
}

func (u UseCaseIndicators) Ema(data []float64, length int32, initial bool) ([]float64, error) {
	resultEma := u.IndicatorsRepository.Ema(data, length, initial)

	return resultEma, nil
}

func (u UseCaseIndicators) Sma(data []float64, length int32, initial bool) ([]float64, error) {
	resultSma := u.IndicatorsRepository.Sma(data, length, initial)

	return resultSma, nil
}

func (u UseCaseIndicators) Macd(data []float64, fastPeriod int32, slowPeriod int32, signalPeriod int32, initial bool) ([]dto.ReturnMacd, error) {
	resultMacd := u.IndicatorsRepository.Macd(data, fastPeriod, slowPeriod, signalPeriod, initial)

	return resultMacd, nil
}

func (u UseCaseIndicators) StochOscillator(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, smoothK int32, smoothD int32, MAUse int32, initial bool) ([]dto.ReturnStoch, error) {
	resultStoch := u.IndicatorsRepository.StochOscillator(dataHigh, dataLow, dataClose, length, smoothK, smoothD, MAUse, initial)

	return resultStoch, nil
}

func (u UseCaseIndicators) Cci(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) ([]float64, error) {
	resultSma := u.IndicatorsRepository.Cci(dataHigh, dataLow, dataClose, length, initial)

	return resultSma, nil
}

func (u UseCaseIndicators) Atr(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) ([]float64, error) {
	resultAtr := u.IndicatorsRepository.Atr(dataHigh, dataLow, dataClose, length, initial)

	return resultAtr, nil
}

func (u UseCaseIndicators) Sar(dataHigh []float64, dataLow []float64, increment float64, maximum float64, initial bool) ([]float64, error) {
	resultSar := u.IndicatorsRepository.Sar(dataHigh, dataLow, increment, maximum, initial)

	return resultSar, nil
}

func (u UseCaseIndicators) Dmi(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) ([]dto.ReturnDmi, error) {
	resultDmi := u.IndicatorsRepository.Dmi(dataHigh, dataLow, dataClose, length, initial)

	return resultDmi, nil
}

func (u UseCaseIndicators) Adx(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) ([]float64, error) {
	resultAdx := u.IndicatorsRepository.Adx(dataHigh, dataLow, dataClose, length, initial)

	return resultAdx, nil
}
