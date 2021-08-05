package repository

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/juninhoitabh/marketanalyzer-talibservice/dto"
)

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestEma(t *testing.T) {
	t.Parallel()
	valorEsperado := []float64{0.40584666666666663, 0.3449333333333333, 0.31473666666666666, 0.2974283333333333}

	p := []dto.PriceHistory{}
	jsonString := `[{"timeHistory":1617235200000,"open":0.61216,"high":0.65021,"low":0.34998,"close":0.53106,"volume":16978421}, 
									{"timeHistory":1619827200000,"open":0.53146,"high":0.79935,"low":0.19615,"close":0.40246,"volume":74437486}, 
									{"timeHistory":1622505600000,"open":0.40218,"high":0.44974,"low":0.19856,"close":0.28402,"volume":39368811}, 
									{"timeHistory":1622505600000,"open":0.40218,"high":0.44974,"low":0.19856,"close":0.28402,"volume":39368811}, 
									{"timeHistory":1625097600000,"open":0.28418,"high":0.28977,"low":0.19815,"close":0.28454,"volume":28918161}, 
									{"timeHistory":1627776000000,"open":0.28437,"high":0.30102,"low":0.26391,"close":0.28012,"volume":4712859}]`
	json.Unmarshal([]byte(jsonString), &p)

	indicators := NewIndicatorsRepository()

	valor := indicators.Ema(p, "Close", 3)

	if !reflect.DeepEqual(valor, valorEsperado) {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
