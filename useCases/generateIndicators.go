package useCases

import (
	"encoding/json"
	"errors"

	"github.com/juninhoitabh/marketanalyzer-talibservice/domain"
	"github.com/juninhoitabh/marketanalyzer-talibservice/dto"
)

type UseCaseIndicators struct {
	IndicatorsRepository domain.Indicators
}

func NewUseCaseIndicators(indicatorsRepository domain.Indicators) UseCaseIndicators {
	return UseCaseIndicators{IndicatorsRepository: indicatorsRepository}
}

func (u UseCaseIndicators) Ema(data []dto.PriceHistory, source string, length int32) (string, error) {
	resultEmaString, err := json.Marshal(u.IndicatorsRepository.Ema(data, source, length))

	if err == nil {
		return string(resultEmaString), nil
	}

	return "", errors.New("Error converting Json")
}
