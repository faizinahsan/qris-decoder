package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"faizinahsan/qris-decoder/application/usecase"
	"faizinahsan/qris-decoder/interfaces/http/dto"
)

type QRISHandler struct{}

func NewQRISHandler() *QRISHandler {
	return &QRISHandler{}
}

func (h *QRISHandler) Decode(c *gin.Context) {
	var req dto.DecodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "raw qris string is required"})
		return
	}

	result, err := usecase.DecodeQRIS(req.Raw)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, toResponse(result))
}

func toResponse(r usecase.DecodeResult) dto.DecodeResponse {
	return dto.DecodeResponse{
		MerchantName: r.MerchantName,
		MerchantCity: r.MerchantCity,
		MCC:          r.MCC,
		Currency:     r.Currency,
		Amount:       r.Amount,
		CountryCode:  r.CountryCode,
		Invoice:      r.Invoice,
		TerminalID:   r.TerminalID,
		CRC:          r.CRC,
		Merchant: dto.MerchantDTO{
			AcquirerGUI: r.Acquirer,
			PAN:         r.MerchantPAN,
			MerchantID:  r.MerchantID,
		},
		Validation: dto.ValidationDTO{
			Valid:         r.Validation.Valid,
			Errors:        r.Validation.Errors,
			IsCrossBorder: r.Validation.IsCrossBorder,
		},
	}
}
