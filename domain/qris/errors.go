package qris

import "errors"

var (
	ErrInvalidFormat   = errors.New("invalid QRIS format")
	ErrCRCMismatch     = errors.New("CRC mismatch")
	ErrMissingTag      = errors.New("missing mandatory tag")
	ErrInvalidMerchant = errors.New("invalid merchant data")
)
