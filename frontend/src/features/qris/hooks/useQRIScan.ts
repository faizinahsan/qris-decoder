import { useState } from 'react'
import jsQR from 'jsqr'

interface UseQRIScanState {
  scanning: boolean
  error: string | null
}

export const useQRIScan = (onDecoded: (raw: string) => void) => {
  const [state, setState] = useState<UseQRIScanState>({ scanning: false, error: null })

  const scanFile = (file: File) => {
    setState({ scanning: true, error: null })

    const img = new Image()
    const url = URL.createObjectURL(file)
    img.src = url

    img.onload = () => {
      const canvas = document.createElement('canvas')
      canvas.width = img.width
      canvas.height = img.height

      const ctx = canvas.getContext('2d')!
      ctx.drawImage(img, 0, 0)

      const imageData = ctx.getImageData(0, 0, canvas.width, canvas.height)
      const result = jsQR(imageData.data, imageData.width, imageData.height)

      URL.revokeObjectURL(url)

      if (!result) {
        setState({ scanning: false, error: 'QR code tidak ditemukan dalam gambar.' })
        return
      }

      setState({ scanning: false, error: null })
      onDecoded(result.data)
    }

    img.onerror = () => {
      URL.revokeObjectURL(url)
      setState({ scanning: false, error: 'Gagal membaca file gambar.' })
    }
  }

  return { ...state, scanFile }
}
