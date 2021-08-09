package domain

import "github.com/juninhoitabh/marketanalyzer-talibservice/dto"

type Indicators interface {
	Ema(data []float64, length int32, initial bool) []float64
	Sma(data []float64, length int32, initial bool) []float64
	Macd(data []float64, fastPeriod int32, slowPeriod int32, signalPeriod int32, initial bool) []dto.ReturnMacd
	StochOscillator(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, smoothK int32, smoothD int32, MAUse int32, initial bool) []dto.ReturnStoch
	Cci(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []float64
	Atr(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []float64
	Sar(dataHigh []float64, dataLow []float64, increment float64, maximum float64, initial bool) []float64
	Dmi(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []dto.ReturnDmi
	Adx(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []float64
}
