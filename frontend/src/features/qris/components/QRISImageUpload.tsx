import { useRef, useState, useCallback } from 'react'
import { useQRIScan } from '../hooks/useQRIScan'

interface Props {
  onDecoded: (raw: string) => void
  loading: boolean
}

export const QRISImageUpload = ({ onDecoded, loading }: Props) => {
  const inputRef = useRef<HTMLInputElement>(null)
  const [preview, setPreview] = useState<string | null>(null)
  const [dragging, setDragging] = useState(false)
  const { scanning, error, scanFile } = useQRIScan(onDecoded)

  const handleFile = useCallback(
    (file: File) => {
      if (!file.type.startsWith('image/')) return
      setPreview(URL.createObjectURL(file))
      scanFile(file)
    },
    [scanFile],
  )

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (file) handleFile(file)
    e.target.value = ''
  }

  const handleDrop = (e: React.DragEvent) => {
    e.preventDefault()
    setDragging(false)
    const file = e.dataTransfer.files?.[0]
    if (file) handleFile(file)
  }

  const busy = scanning || loading

  return (
    <div style={styles.wrapper}>
      <div
        style={{ ...styles.dropzone, ...(dragging ? styles.dropzoneDragging : {}) }}
        onClick={() => !busy && inputRef.current?.click()}
        onDragOver={(e) => { e.preventDefault(); setDragging(true) }}
        onDragLeave={() => setDragging(false)}
        onDrop={handleDrop}
      >
        <input
          ref={inputRef}
          type="file"
          accept="image/*"
          style={{ display: 'none' }}
          onChange={handleChange}
        />

        {preview ? (
          <img src={preview} alt="preview" style={styles.preview} />
        ) : (
          <div style={styles.placeholder}>
            <span style={styles.icon}>📷</span>
            <span style={styles.hint}>Upload atau drag & drop foto QRIS</span>
            <span style={styles.hint2}>PNG, JPG, WEBP</span>
          </div>
        )}
      </div>

      {busy && <p style={styles.status}>🔍 Membaca QR code...</p>}
      {error && <p style={styles.error}>{error}</p>}
    </div>
  )
}

const styles: Record<string, React.CSSProperties> = {
  wrapper: { display: 'flex', flexDirection: 'column', gap: 8 },
  dropzone: {
    border: '2px dashed #334155',
    borderRadius: 10,
    padding: 24,
    cursor: 'pointer',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    minHeight: 140,
    transition: 'border-color 0.2s',
    background: '#0f172a',
  },
  dropzoneDragging: { borderColor: '#3b82f6', background: '#0f2040' },
  placeholder: { display: 'flex', flexDirection: 'column', alignItems: 'center', gap: 6 },
  icon: { fontSize: 32 },
  hint: { color: '#94a3b8', fontSize: 14 },
  hint2: { color: '#475569', fontSize: 12 },
  preview: { maxHeight: 200, maxWidth: '100%', borderRadius: 6, objectFit: 'contain' },
  status: { margin: 0, color: '#60a5fa', fontSize: 13 },
  error: { margin: 0, color: '#fca5a5', fontSize: 13 },
}
