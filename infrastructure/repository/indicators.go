package repository

import (
	"reflect"

	"github.com/juninhoitabh/marketanalyzer-talibservice/dto"
	"github.com/juninhoitabh/marketanalyzer-talibservice/utils"
)

type IndicatorsRepository struct{}

func makeDataSource(data []dto.PriceHistory, source string) []float64 {
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

func NewIndicatorsRepository() *IndicatorsRepository {
	return &IndicatorsRepository{}
}

func (i *IndicatorsRepository) Ema(data []dto.PriceHistory, source string, length int32) []float64 {
	return utils.Ema(makeDataSource(data, source), source, length)
}
