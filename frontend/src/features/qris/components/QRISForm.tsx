import { useState } from 'react'

interface Props {
  onSubmit: (raw: string) => void
  loading: boolean
}

const SAMPLE_QR =
  '00020101021151480018ID.CO.MINIMART.WWW0215ID10190020904070303UME5204517253033605802ID5910Minimarket6010Tanggerang61051290062070703A016304B38F'

export const QRISForm = ({ onSubmit, loading }: Props) => {
  const [value, setValue] = useState('')

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (value.trim()) onSubmit(value.trim())
  }

  return (
    <form onSubmit={handleSubmit} style={styles.form}>
      <textarea
        value={value}
        onChange={(e) => setValue(e.target.value)}
        placeholder="Paste QRIS string di sini..."
        rows={4}
        style={styles.textarea}
      />
      <div style={styles.actions}>
        <button
          type="button"
          onClick={() => setValue(SAMPLE_QR)}
          style={styles.sampleBtn}
        >
          Pakai Contoh
        </button>
        <button type="submit" disabled={loading || !value.trim()} style={styles.submitBtn}>
          {loading ? 'Memproses...' : 'Decode QRIS'}
        </button>
      </div>
    </form>
  )
}

const styles: Record<string, React.CSSProperties> = {
  form: { display: 'flex', flexDirection: 'column', gap: 12 },
  textarea: {
    padding: 12,
    borderRadius: 8,
    border: '1px solid #334155',
    background: '#0f172a',
    color: '#e2e8f0',
    fontSize: 13,
    fontFamily: 'monospace',
    resize: 'vertical',
  },
  actions: { display: 'flex', gap: 8, justifyContent: 'flex-end' },
  sampleBtn: {
    padding: '8px 16px',
    borderRadius: 6,
    border: '1px solid #475569',
    background: 'transparent',
    color: '#94a3b8',
    cursor: 'pointer',
    fontSize: 13,
  },
  submitBtn: {
    padding: '8px 20px',
    borderRadius: 6,
    border: 'none',
    background: '#3b82f6',
    color: '#fff',
    cursor: 'pointer',
    fontSize: 13,
    fontWeight: 600,
  },
}
