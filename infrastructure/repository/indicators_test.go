package repository

import (
	"reflect"
	"testing"

	"github.com/juninhoitabh/marketanalyzer-talibservice/dto"
)

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestEma(t *testing.T) {
	t.Parallel()
	valorEsperado := []float64{0.40584666666666663, 0.3449333333333333, 0.31473666666666666, 0.2974283333333333}
	data := []float64{0.53106, 0.40246, 0.28402, 0.28402, 0.28454, 0.28012}

	indicators := NewIndicatorsRepository()
	valor := indicators.Ema(data, 3, false)

	if !reflect.DeepEqual(valor, valorEsperado) {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestSma(t *testing.T) {
	t.Parallel()
	valorEsperado := []float64{0.40584666666666663, 0.32349999999999995, 0.2841933333333333, 0.2828933333333333}
	data := []float64{0.53106, 0.40246, 0.28402, 0.28402, 0.28454, 0.28012}

	indicators := NewIndicatorsRepository()
	valor := indicators.Sma(data, 3, false)

	if !reflect.DeepEqual(valor, valorEsperado) {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

func TestMacd(t *testing.T) {
	t.Parallel()
	valorEsperado := []dto.ReturnMacd{
		{
			FastLength:   -0.00028556676001980285,
			Macd:         0.0028480147280265378,
			SignalLength: 0.0031335814880463406,
		},
		{
			FastLength:   -0.0003093755301724222,
			Macd:         0.0027679558614789324,
			SignalLength: 0.0030773313916513546,
		},
		{
			FastLength:   -0.00027994107199454457,
			Macd:         0.002746491942930529,
			SignalLength: 0.0030264330149250736,
		},
		{
			FastLength:   -7.099870748817835e-05,
			Macd:         0.0029425254515299537,
			SignalLength: 0.003013524159018132,
		},
	}
	data := []float64{0.63022, 0.65428, 0.63118, 0.66266, 0.63521, 0.64039, 0.66142, 0.65684, 0.65138, 0.62829, 0.63738, 0.64784, 0.64848, 0.63498, 0.64495, 0.64665, 0.63646, 0.64373, 0.64629, 0.64259, 0.64138, 0.65168, 0.64367, 0.649, 0.64948, 0.65398, 0.65279, 0.64886, 0.65198, 0.6527, 0.6489, 0.64787, 0.65402, 0.65187, 0.64758, 0.6499, 0.65352, 0.65381, 0.65779, 0.65338, 0.65449, 0.65477, 0.65597, 0.65159, 0.65507, 0.65622, 0.6551, 0.64921, 0.6513, 0.64803, 0.64614, 0.6458, 0.65309, 0.64821, 0.6575, 0.65471, 0.65256, 0.65322, 0.65197, 0.6535, 0.65737, 0.65313, 0.65661, 0.65389, 0.65295, 0.65446, 0.65444, 0.65383, 0.6552, 0.66052}

	indicators := NewIndicatorsRepository()
	valor := indicators.Macd(data, 25, 56, 12, false)

	if !reflect.DeepEqual(valor, valorEsperado) {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
