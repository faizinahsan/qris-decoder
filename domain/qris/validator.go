package qris

import (
	"fmt"
	"strings"
)

// ValidationResult adalah output dari Domain Service Validator.
type ValidationResult struct {
	Valid         bool
	Errors        []string
	IsCrossBorder bool
}

// Validate adalah Domain Service.
// Menerima aggregate dan menjalankan semua business rule validasi.
func Validate(payload QRISPayload) ValidationResult {
	var errs []string

	// Rule 1: mandatory tags
	for _, tag := range MandatoryTags {
		if _, ok := payload.TagValue(tag); !ok {
			errs = append(errs, fmt.Sprintf("missing mandatory tag %02d", tag))
		}
	}

	// Rule 2: CRC harus valid
	if err := validateCRC(payload.Raw(), payload.CRC()); err != nil {
		errs = append(errs, err.Error())
	}

	// Rule 3: merchant harus valid
	if !payload.Merchant().IsValid() {
		errs = append(errs, ErrInvalidMerchant.Error())
	}

	isCrossBorder := isCrossBorder(payload)

	return ValidationResult{
		Valid:         len(errs) == 0,
		Errors:        errs,
		IsCrossBorder: isCrossBorder,
	}
}

func validateCRC(raw, crc string) error {
	idx := strings.Index(raw, "6304")
	if idx == -1 {
		return ErrCRCMismatch
	}
	payload := raw[:idx+4]
	calculated := fmt.Sprintf("%04X", crc16CCITT(payload))
	if calculated != crc {
		return ErrCRCMismatch
	}
	return nil
}

func isCrossBorder(payload QRISPayload) bool {
	country, _ := payload.TagValue(58)
	return country != "ID" || !payload.Merchant().IsDomestic()
}

func crc16CCITT(data string) uint16 {
	var crc uint16 = 0xFFFF
	for i := 0; i < len(data); i++ {
		crc ^= uint16(data[i]) << 8
		for j := 0; j < 8; j++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}
	return crc
}
