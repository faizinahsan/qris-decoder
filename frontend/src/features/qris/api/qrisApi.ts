import axios from 'axios'
import type { DecodeResponse } from '../types'

const client = axios.create({
  baseURL: '/api/v1',
  headers: { 'Content-Type': 'application/json' },
})

export const decodeQRIS = async (raw: string): Promise<DecodeResponse> => {
  const { data } = await client.post<DecodeResponse>('/qris/decode', { raw })
  return data
}
