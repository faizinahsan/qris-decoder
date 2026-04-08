package dto

type DecodeRequest struct {
	Raw string `json:"raw" binding:"required"`
}

type MerchantDTO struct {
	AcquirerGUI string `json:"acquirer_gui"`
	PAN         string `json:"pan"`
	MerchantID  string `json:"merchant_id"`
	Criteria    string `json:"criteria"`
}

type ValidationDTO struct {
	Valid         bool     `json:"valid"`
	Errors        []string `json:"errors"`
	IsCrossBorder bool     `json:"is_cross_border"`
}

type DecodeResponse struct {
	MerchantName string        `json:"merchant_name"`
	MerchantCity string        `json:"merchant_city"`
	MCC          string        `json:"mcc"`
	Currency     string        `json:"currency"`
	Amount       string        `json:"amount"`
	CountryCode  string        `json:"country_code"`
	Invoice      string        `json:"invoice"`
	TerminalID   string        `json:"terminal_id"`
	CRC          string        `json:"crc"`
	Merchant     MerchantDTO   `json:"merchant"`
	Validation   ValidationDTO `json:"validation"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
