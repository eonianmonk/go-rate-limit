package responses

type GetRate struct {
	Rate int16 `json:"rate"`
}

func NewGetRateResponse(rate int16) *GetRate {
	return &GetRate{Rate: rate}
}
