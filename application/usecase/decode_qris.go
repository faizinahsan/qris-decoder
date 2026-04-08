package usecase

import (
	"faizinahsan/qris-decoder/domain/qris"
	"faizinahsan/qris-decoder/infrastructure/parser"
)

// DecodeResult adalah output DTO dari use case ini.
// Ini yang akan dikonsumsi oleh interface layer (CLI, HTTP handler, dll).
type DecodeResult struct {
	MerchantName string
	MerchantCity string
	MCC          string
	Currency     string
	Amount       string
	CountryCode  string
	Acquirer     string
	MerchantPAN  string
	MerchantID   string
	Invoice      string
	TerminalID   string
	CRC          string
	Validation   qris.ValidationResult
	AllFields    map[int]qris.Field
	AllSubFields map[int]map[int]qris.Field
}

// DecodeQRIS adalah use case utama: parse + validate QRIS string.
func DecodeQRIS(raw string) (DecodeResult, error) {
	payload, err := parser.Parse(raw)
	if err != nil {
		return DecodeResult{}, err
	}

	validation := qris.Validate(payload)

	invoice, _ := payload.SubTagValue(62, 5)
	if invoice == "" {
		invoice, _ = payload.SubTagValue(62, 1)
	}
	terminalID, _ := payload.SubTagValue(62, 7)

	amount, _ := payload.TagValue(54)
	merchantName, _ := payload.TagValue(59)
	merchantCity, _ := payload.TagValue(60)
	mcc, _ := payload.TagValue(52)
	currency, _ := payload.TagValue(53)
	country, _ := payload.TagValue(58)

	return DecodeResult{
		MerchantName: merchantName,
		MerchantCity: merchantCity,
		MCC:          mcc,
		Currency:     currency,
		Amount:       amount,
		CountryCode:  country,
		Acquirer:     payload.Merchant().AcquirerGUI(),
		MerchantPAN:  payload.Merchant().PAN(),
		MerchantID:   payload.Merchant().MerchantID(),
		Invoice:      invoice,
		TerminalID:   terminalID,
		CRC:          payload.CRC(),
		Validation:   validation,
		AllFields:    payload.Fields(),
		AllSubFields: payload.SubFields(),
	}, nil
}
