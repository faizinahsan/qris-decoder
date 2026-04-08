import type { ValidationDTO } from '../types'

interface Props {
  validation: ValidationDTO
}

export const ValidationBadge = ({ validation }: Props) => (
  <div style={styles.wrapper}>
    <div style={{ ...styles.badge, background: validation.valid ? '#166534' : '#7f1d1d' }}>
      {validation.valid ? '✓ VALID QRIS' : '✗ INVALID QRIS'}
    </div>
    <div style={styles.badge2}>
      {validation.is_cross_border ? '🌐 Cross Border' : '🇮🇩 Domestic'}
    </div>
    {!validation.valid && validation.errors && (
      <ul style={styles.errors}>
        {validation.errors.map((e, i) => (
          <li key={i} style={styles.errorItem}>
            {e}
          </li>
        ))}
      </ul>
    )}
  </div>
)

const styles: Record<string, React.CSSProperties> = {
  wrapper: { display: 'flex', flexWrap: 'wrap', gap: 8, alignItems: 'center' },
  badge: {
    padding: '4px 12px',
    borderRadius: 20,
    fontSize: 12,
    fontWeight: 700,
    color: '#fff',
    letterSpacing: 0.5,
  },
  badge2: {
    padding: '4px 12px',
    borderRadius: 20,
    fontSize: 12,
    background: '#1e293b',
    color: '#94a3b8',
  },
  errors: { width: '100%', margin: '4px 0 0', padding: '0 0 0 16px' },
  errorItem: { color: '#fca5a5', fontSize: 12, marginBottom: 2 },
}
