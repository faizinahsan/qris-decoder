import { useState } from 'react'
import { useQRISDecode } from './features/qris/hooks/useQRISDecode'
import { QRISForm } from './features/qris/components/QRISForm'
import { QRISImageUpload } from './features/qris/components/QRISImageUpload'
import { QRISResult } from './features/qris/components/QRISResult'

type Tab = 'manual' | 'image'

export default function App() {
  const [tab, setTab] = useState<Tab>('manual')
  const { result, loading, error, decode } = useQRISDecode()

  return (
    <div style={styles.page}>
      <div style={styles.container}>
        <header style={styles.header}>
          <h1 style={styles.title}>QRIS Decoder</h1>
          <p style={styles.subtitle}>Parse dan validasi QRIS string sesuai standar BI</p>
        </header>

        <div style={styles.tabs}>
          <button
            style={{ ...styles.tab, ...(tab === 'manual' ? styles.tabActive : {}) }}
            onClick={() => setTab('manual')}
          >
            ✏️ Input Manual
          </button>
          <button
            style={{ ...styles.tab, ...(tab === 'image' ? styles.tabActive : {}) }}
            onClick={() => setTab('image')}
          >
            📷 Upload Foto
          </button>
        </div>

        {tab === 'manual' && <QRISForm onSubmit={decode} loading={loading} />}
        {tab === 'image' && <QRISImageUpload onDecoded={decode} loading={loading} />}

        {error && <div style={styles.errorBox}>{error}</div>}
        {result && <QRISResult data={result} />}
      </div>
    </div>
  )
}

const styles: Record<string, React.CSSProperties> = {
  page: {
    minHeight: '100vh',
    background: '#020617',
    color: '#e2e8f0',
    padding: '40px 16px',
    fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif',
  },
  container: {
    maxWidth: 640,
    margin: '0 auto',
    display: 'flex',
    flexDirection: 'column',
    gap: 24,
  },
  header: { textAlign: 'center' },
  title: { margin: 0, fontSize: 28, fontWeight: 700, color: '#f8fafc' },
  subtitle: { margin: '8px 0 0', color: '#64748b', fontSize: 14 },
  tabs: {
    display: 'flex',
    gap: 8,
    background: '#0f172a',
    padding: 4,
    borderRadius: 10,
    border: '1px solid #1e293b',
  },
  tab: {
    flex: 1,
    padding: '8px 0',
    borderRadius: 7,
    border: 'none',
    background: 'transparent',
    color: '#64748b',
    fontSize: 13,
    fontWeight: 500,
    cursor: 'pointer',
    transition: 'all 0.15s',
  },
  tabActive: {
    background: '#1e293b',
    color: '#f1f5f9',
  },
  errorBox: {
    padding: '12px 16px',
    borderRadius: 8,
    background: '#450a0a',
    border: '1px solid #7f1d1d',
    color: '#fca5a5',
    fontSize: 13,
  },
}
