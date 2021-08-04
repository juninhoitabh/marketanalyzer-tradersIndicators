package useCases

import (
	"reflect"

	"code.afis.me/trading/indicator/talib"
)

type PriceHistory struct {
	TimeHistory int64   `json:"timeHistory"`
	Open        float64 `json:"open"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	Close       float64 `json:"close"`
	Volume      int64   `json:"volume"`
}

func makeDataSource(data []PriceHistory, source string) []float64 {
	newData := make([]float64, 0)
	for _, v := range data {
		priceHistory := reflect.ValueOf(v)
		field := reflect.Indirect(priceHistory).FieldByName(source)
		if field.Kind().String() == "float64" {
			value := float64(field.Float())
			newData = append(newData, value)
		}
	}

	return newData
}

func Ema(data []PriceHistory, source string, length int32) []float64 {
	emaResult := talib.Ema(makeDataSource(data, source), length)

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
