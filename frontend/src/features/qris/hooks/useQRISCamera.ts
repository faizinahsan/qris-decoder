import { useRef, useState, useEffect, useCallback } from 'react'
import jsQR from 'jsqr'

interface UseQRISCameraState {
  visible: boolean  // apakah <video> sudah di-render di DOM
  active: boolean   // apakah stream sudah berjalan
  error: string | null
  scanning: boolean
}

export const useQRISCamera = (onDecoded: (raw: string) => void) => {
  const videoRef = useRef<HTMLVideoElement>(null)
  const streamRef = useRef<MediaStream | null>(null)
  const rafRef = useRef<number>(0)
  const [state, setState] = useState<UseQRISCameraState>({
    visible: false,
    active: false,
    error: null,
    scanning: false,
  })

  const stop = useCallback(() => {
    cancelAnimationFrame(rafRef.current)
    streamRef.current?.getTracks().forEach((t) => t.stop())
    streamRef.current = null
    if (videoRef.current) videoRef.current.srcObject = null
    setState({ visible: false, active: false, error: null, scanning: false })
  }, [])

  const start = useCallback(async () => {
    // Tampilkan <video> element dulu ke DOM, baru assign stream
    setState({ visible: true, active: false, error: null, scanning: true })
  }, [])

  // Effect ini jalan setelah visible=true, saat <video> sudah ada di DOM
  useEffect(() => {
    if (!state.visible || state.active || !state.scanning) return

    const run = async () => {
      try {
        const stream = await navigator.mediaDevices.getUserMedia({
          video: { facingMode: 'environment' },
        })
        streamRef.current = stream

        const video = videoRef.current
        if (!video) {
          stream.getTracks().forEach((t) => t.stop())
          setState({ visible: false, active: false, error: 'Video element tidak ditemukan.', scanning: false })
          return
        }

        video.srcObject = stream
        await video.play()
        setState({ visible: true, active: true, error: null, scanning: false })
      } catch {
        setState({ visible: false, active: false, error: 'Tidak dapat mengakses kamera. Pastikan izin kamera sudah diberikan.', scanning: false })
      }
    }

    run()
  }, [state.visible, state.active, state.scanning])

  // Scan frame realtime
  useEffect(() => {
    if (!state.active) return

    const canvas = document.createElement('canvas')
    const ctx = canvas.getContext('2d')!

    const tick = () => {
      const video = videoRef.current
      if (video && video.readyState === video.HAVE_ENOUGH_DATA) {
        canvas.width = video.videoWidth
        canvas.height = video.videoHeight
        ctx.drawImage(video, 0, 0)
        const imageData = ctx.getImageData(0, 0, canvas.width, canvas.height)
        const result = jsQR(imageData.data, imageData.width, imageData.height)
        if (result?.data) {
          stop()
          onDecoded(result.data)
          return
        }
      }
      rafRef.current = requestAnimationFrame(tick)
    }

    rafRef.current = requestAnimationFrame(tick)
    return () => cancelAnimationFrame(rafRef.current)
  }, [state.active, stop, onDecoded])

  useEffect(() => () => stop(), [stop])

  return { videoRef, ...state, start, stop }
}
