package dto

type PriceHistory struct {
	TimeHistory int64   `json:"timeHistory"`
	Open        float64 `json:"open"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	Close       float64 `json:"close"`
	Volume      int64   `json:"volume"`
}

type ReturnMacd struct {
	FastLength   float64 `json:"fastLength"`
	Macd         float64 `json:"macd"`
	SignalLength float64 `json:"signalLength"`
}

type ReturnStoch struct {
	SmoothK float64 `json:"smoothK"`
	SmoothD float64 `json:"smoothD"`
}

type ReturnDmi struct {
	Adx    float64 `json:"adx"`
	PlusD  float64 `json:"plusD"`
	MinusD float64 `json:"minusD"`
}
