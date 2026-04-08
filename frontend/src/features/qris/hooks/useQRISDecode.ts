import { useState } from 'react'
import { decodeQRIS } from '../api/qrisApi'
import type { DecodeResponse } from '../types'

interface UseQRISDecodeState {
  result: DecodeResponse | null
  loading: boolean
  error: string | null
}

export const useQRISDecode = () => {
  const [state, setState] = useState<UseQRISDecodeState>({
    result: null,
    loading: false,
    error: null,
  })

  const decode = async (raw: string) => {
    setState({ result: null, loading: true, error: null })
    try {
      const result = await decodeQRIS(raw)
      setState({ result, loading: false, error: null })
    } catch (err: unknown) {
      const message =
        axios_error(err) ?? 'Gagal menghubungi server, pastikan backend berjalan.'
      setState({ result: null, loading: false, error: message })
    }
  }

  return { ...state, decode }
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const axios_error = (err: any): string | null =>
  err?.response?.data?.message ?? err?.message ?? null
