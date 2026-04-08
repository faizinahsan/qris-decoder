interface Props {
  label: string
  value: string | undefined
}

export const DataRow = ({ label, value }: Props) => {
  if (!value) return null
  return (
    <div style={styles.row}>
      <span style={styles.label}>{label}</span>
      <span style={styles.value}>{value}</span>
    </div>
  )
}

const styles: Record<string, React.CSSProperties> = {
  row: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'flex-start',
    padding: '8px 0',
    borderBottom: '1px solid #1e293b',
    gap: 16,
  },
  label: { color: '#64748b', fontSize: 13, flexShrink: 0, minWidth: 140 },
  value: { color: '#e2e8f0', fontSize: 13, fontFamily: 'monospace', textAlign: 'right', wordBreak: 'break-all' },
}
