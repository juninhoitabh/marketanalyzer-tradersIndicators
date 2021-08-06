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

func (u UseCaseIndicators) Ema(data []dto.PriceHistory, source string, length int32) ([]float64, error) {
	resultEmaString := u.IndicatorsRepository.Ema(data, source, length)

	return resultEmaString, nil
}
