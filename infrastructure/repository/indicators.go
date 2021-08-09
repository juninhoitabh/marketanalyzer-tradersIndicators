package repository

import (
	"github.com/juninhoitabh/marketanalyzer-talibservice/dto"
	"github.com/juninhoitabh/marketanalyzer-talibservice/utils"
)

type IndicatorsRepository struct{}

/* func makeDataSource(data []dto.PriceHistory, source string) []float64 {
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
} */

func NewIndicatorsRepository() *IndicatorsRepository {
	return &IndicatorsRepository{}
}

func (i *IndicatorsRepository) Ema(data []float64, length int32, initial bool) []float64 {
	return utils.Ema(data, length, initial)
}

func (i *IndicatorsRepository) Sma(data []float64, length int32, initial bool) []float64 {
	return utils.Sma(data, length, initial)
}

func (i *IndicatorsRepository) Macd(data []float64, fastPeriod int32, slowPeriod int32, signalPeriod int32, initial bool) []dto.ReturnMacd {
	return utils.Macd(data, fastPeriod, slowPeriod, signalPeriod, initial)
}

func (i *IndicatorsRepository) StochOscillator(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, smoothK int32, smoothD int32, MAUse int32, initial bool) []dto.ReturnStoch {
	return utils.StochOscillator(dataHigh, dataLow, dataClose, length, smoothK, smoothD, MAUse, initial)
}

func (i *IndicatorsRepository) Cci(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []float64 {
	return utils.Cci(dataHigh, dataLow, dataClose, length, initial)
}

func (i *IndicatorsRepository) Atr(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []float64 {
	return utils.Atr(dataHigh, dataLow, dataClose, length, initial)
}

func (i *IndicatorsRepository) Sar(dataHigh []float64, dataLow []float64, increment float64, maximum float64, initial bool) []float64 {
	return utils.Sar(dataHigh, dataLow, increment, maximum, initial)
}

func (i *IndicatorsRepository) Dmi(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []dto.ReturnDmi {
	return utils.Dmi(dataHigh, dataLow, dataClose, length, initial)
}
func (i *IndicatorsRepository) Adx(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []float64 {
	return utils.Adx(dataHigh, dataLow, dataClose, length, initial)
}
