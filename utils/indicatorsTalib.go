package utils

import (
	"code.afis.me/trading/indicator/talib"
	"github.com/juninhoitabh/marketanalyzer-talibservice/dto"
)

func Ema(data []float64, length int32, initial bool) []float64 {
	result := make([]float64, 0)

	if len(data) > 0 {
		emaResult := talib.Ema(data, length)

		lenResult := 0
		if len(emaResult) >= 50 && !initial {
			lenResult = len(emaResult) - 50
		}
		for i := lenResult; i < len(emaResult); i++ {
			result = append(result, emaResult[i])
		}
	}

	return result
}

func Sma(data []float64, length int32, initial bool) []float64 {
	result := make([]float64, 0)

	if len(data) > 0 {
		smaResult := talib.Sma(data, length)

		lenResult := 0
		if len(smaResult) >= 50 && !initial {
			lenResult = len(smaResult) - 50
		}
		for i := lenResult; i < len(smaResult); i++ {
			result = append(result, smaResult[i])
		}
	}

	return result
}

func Macd(data []float64, fastPeriod int32, slowPeriod int32, signalPeriod int32, initial bool) []dto.ReturnMacd {
	result := make([]dto.ReturnMacd, 0)

	if len(data) > 0 {
		slow, signal, fast := talib.Macd(data, fastPeriod, slowPeriod, signalPeriod)

		lenResult := 0
		if len(fast) >= 50 && !initial {
			lenResult = len(fast) - 50
		}
		for i := lenResult; i < len(fast); i++ {
			returnMacd := dto.ReturnMacd{
				FastLength:   fast[i],
				Macd:         slow[i],
				SignalLength: signal[i],
			}

			result = append(result, returnMacd)
		}
	}

	return result
}

func StochOscillator(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, smoothK int32, smoothD int32, mAUse int32, initial bool) []dto.ReturnStoch {
	result := make([]dto.ReturnStoch, 0)

	if len(dataHigh) > 0 {
		k, d := talib.Stoch(dataHigh, dataLow, dataClose, length, smoothK, mAUse, smoothD, 0)

		lenResult := 0
		if len(d) >= 50 && !initial {
			lenResult = len(d) - 50
		}
		for i := lenResult; i < len(d); i++ {
			returnStoch := dto.ReturnStoch{
				SmoothK: k[i],
				SmoothD: d[i],
			}

			result = append(result, returnStoch)
		}
	}

	return result
}

func Cci(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []float64 {
	result := make([]float64, 0)

	if len(dataHigh) > 0 {
		cciResult := talib.Cci(dataHigh, dataLow, dataClose, length)

		lenResult := 0
		if len(cciResult) >= 50 && !initial {
			lenResult = len(cciResult) - 50
		}
		for i := lenResult; i < len(cciResult); i++ {
			result = append(result, cciResult[i])
		}
	}

	return result
}

func Atr(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []float64 {
	result := make([]float64, 0)

	if len(dataHigh) > 0 {
		atrResult := talib.Atr(dataHigh, dataLow, dataClose, length)

		lenResult := 0
		if len(atrResult) >= 50 && !initial {
			lenResult = len(atrResult) - 50
		}
		for i := lenResult; i < len(atrResult); i++ {
			result = append(result, atrResult[i])
		}
	}

	return result
}

func Sar(dataHigh []float64, dataLow []float64, increment float64, maximum float64, initial bool) []float64 {
	result := make([]float64, 0)

	if len(dataHigh) > 0 {
		sarResult := talib.Sar(dataHigh, dataLow, increment, maximum)

		lenResult := 0
		if len(sarResult) >= 50 && !initial {
			lenResult = len(sarResult) - 50
		}
		for i := lenResult; i < len(sarResult); i++ {
			result = append(result, sarResult[i])
		}
	}

	return result
}

func Dmi(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []dto.ReturnDmi {
	result := make([]dto.ReturnDmi, 0)

	if len(dataHigh) > 0 {
		adxChan := make(chan []float64)
		plusDChan := make(chan []float64)
		minusDChan := make(chan []float64)

		go func() {
			adxChan <- talib.Adx(dataHigh, dataLow, dataClose, length)
		}()
		go func() {
			plusDChan <- talib.PlusDi(dataHigh, dataLow, dataClose, length)
		}()
		go func() {
			minusDChan <- talib.MinusDi(dataHigh, dataLow, dataClose, length)
		}()

		adx := <-adxChan
		plusD := <-plusDChan
		minusD := <-minusDChan

		lenResult := 0
		if len(adx) >= 50 && !initial {
			lenResult = len(adx) - 50
		}
		for i := lenResult; i < len(adx); i++ {
			ReturnDmi := dto.ReturnDmi{
				Adx:    adx[i],
				PlusD:  plusD[i],
				MinusD: minusD[i],
			}

			result = append(result, ReturnDmi)
		}
	}

	return result
}

func Adx(dataHigh []float64, dataLow []float64, dataClose []float64, length int32, initial bool) []float64 {
	result := make([]float64, 0)

	if len(dataHigh) > 0 {
		adxResult := talib.Adx(dataHigh, dataLow, dataClose, length)

		lenResult := 0
		if len(adxResult) >= 50 && !initial {
			lenResult = len(adxResult) - 50
		}
		for i := lenResult; i < len(adxResult); i++ {
			result = append(result, adxResult[i])
		}
	}

	return result
}
