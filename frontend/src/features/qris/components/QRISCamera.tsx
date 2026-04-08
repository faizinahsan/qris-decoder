import './QRISCamera.css'
import { useQRISCamera } from '../hooks/useQRISCamera'

interface Props {
  onDecoded: (raw: string) => void
  loading: boolean
}

export const QRISCamera = ({ onDecoded, loading }: Props) => {
  const { videoRef, visible, active, error, scanning, start, stop } = useQRISCamera(onDecoded)

  return (
    <div style={styles.wrapper}>
      <div style={styles.viewfinder}>
        {/* video selalu di-render saat visible=true agar ref tersedia di DOM */}
        <video
          ref={videoRef}
          style={{ ...styles.video, display: visible ? 'block' : 'none' }}
          muted
          playsInline
        />

        {visible && (
          <div style={styles.overlay}>
            <div style={styles.scanBox}>
              <span className="qris-corner tl" />
              <span className="qris-corner tr" />
              <span className="qris-corner bl" />
              <span className="qris-corner br" />
            </div>
            <p style={styles.hint}>
              {active ? 'Arahkan kamera ke QR code' : 'Membuka kamera...'}
            </p>
          </div>
        )}

        {!visible && (
          <div style={styles.placeholder}>
            <span style={styles.icon}>📷</span>
            <span style={styles.placeholderText}>Kamera belum aktif</span>
          </div>
        )}
      </div>

      <div style={styles.actions}>
        {!visible ? (
          <button
            onClick={start}
            disabled={scanning || loading}
            style={styles.startBtn}
          >
            {scanning ? 'Membuka kamera...' : 'Aktifkan Kamera'}
          </button>
        ) : (
          <button onClick={stop} style={styles.stopBtn}>
            Matikan Kamera
          </button>
        )}
      </div>

      {error && <p style={styles.error}>{error}</p>}
      {active && <p style={styles.status}>🔍 Scanning otomatis...</p>}
    </div>
  )
}

const styles: Record<string, React.CSSProperties> = {
  wrapper: { display: 'flex', flexDirection: 'column', gap: 12 },
  viewfinder: {
    position: 'relative',
    width: '100%',
    aspectRatio: '1',
    borderRadius: 10,
    overflow: 'hidden',
    background: '#0f172a',
    border: '1px solid #1e293b',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
  video: { width: '100%', height: '100%', objectFit: 'cover' },
  overlay: {
    position: 'absolute',
    inset: 0,
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    gap: 16,
  },
  scanBox: {
    position: 'relative',
    width: '60%',
    aspectRatio: '1',
  },
  corner: {},
  hint: { color: '#94a3b8', fontSize: 12, margin: 0, textAlign: 'center' },
  placeholder: { display: 'flex', flexDirection: 'column', alignItems: 'center', gap: 8 },
  icon: { fontSize: 40 },
  placeholderText: { color: '#475569', fontSize: 14 },
  actions: { display: 'flex', justifyContent: 'center' },
  startBtn: {
    padding: '10px 28px',
    borderRadius: 8,
    border: 'none',
    background: '#3b82f6',
    color: '#fff',
    fontSize: 14,
    fontWeight: 600,
    cursor: 'pointer',
  },
  stopBtn: {
    padding: '10px 28px',
    borderRadius: 8,
    border: '1px solid #475569',
    background: 'transparent',
    color: '#94a3b8',
    fontSize: 14,
    cursor: 'pointer',
  },
  status: { margin: 0, color: '#60a5fa', fontSize: 13, textAlign: 'center' },
  error: { margin: 0, color: '#fca5a5', fontSize: 13, textAlign: 'center' },
}
