export interface MerchantDTO {
  acquirer_gui: string
  pan: string
  merchant_id: string
  criteria: string
}

export interface ValidationDTO {
  valid: boolean
  errors: string[] | null
  is_cross_border: boolean
}

export interface DecodeResponse {
  merchant_name: string
  merchant_city: string
  mcc: string
  currency: string
  amount: string
  country_code: string
  invoice: string
  terminal_id: string
  crc: string
  merchant: MerchantDTO
  validation: ValidationDTO
}

export interface ErrorResponse {
  message: string
}
