package domain

import "github.com/juninhoitabh/marketanalyzer-talibservice/dto"

type Indicators interface {
	Ema(data []dto.PriceHistory, source string, length int32) []float64
}
