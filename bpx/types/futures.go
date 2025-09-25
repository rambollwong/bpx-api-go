package types

type Positon struct {
	UserId                   int32    `json:"userId"`
	Symbol                   string   `json:"symbol"`
	PositonId                string   `json:"positonId"`
	BreakEvenPrice           string   `json:"breakEvenPrice"`
	EntryPrice               string   `json:"entryPrice"`
	EstLiquidationPrice      string   `json:"estLiquidationPrice"`
	Imf                      string   `json:"imf"`
	ImfFunction              Function `json:"imfFunction"`
	MarkPrice                string   `json:"markPrice"`
	Mmf                      string   `json:"mmf"`
	MmfFunction              Function `json:"mmfFunction"`
	NetCost                  string   `json:"netCost"`
	NetQuantity              string   `json:"netQuantity"`
	NetExposureQuantity      string   `json:"netExposureQuantity"`
	NetExposureNotional      string   `json:"netExposureNotional"`
	PnlRealized              string   `json:"pnlRealized"`
	PnlUnrealized            string   `json:"pnlUnrealized"`
	CumulativeFundingPayment string   `json:"cumulativeFundingPayment"`
	SubaccountId             uint16   `json:"subaccountId"`
	CumulativeInterest       string   `json:"cumulativeInterest"`
}

type Function struct {
	Type   string `json:"type"`
	Base   string `json:"base"`
	Factor string `json:"factor"`
}
