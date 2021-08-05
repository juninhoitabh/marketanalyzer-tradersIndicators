package utils

import (
	"code.afis.me/trading/indicator/talib"
)

func Ema(data []float64, source string, length int32) []float64 {
	emaResult := talib.Ema(data, length)

	result := make([]float64, 0)

	lenResult := 0
	if len(emaResult) >= 50 {
		lenResult = len(emaResult) - 50
	}
	for i := lenResult; i < len(emaResult); i++ {
		result = append(result, emaResult[i])
	}

	return result
}
