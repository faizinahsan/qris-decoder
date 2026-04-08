import type { DecodeResponse } from '../types'
import { ValidationBadge } from './ValidationBadge'
import { DataRow } from './DataRow'

interface Props {
  data: DecodeResponse
}

const CURRENCY_MAP: Record<string, string> = { '360': 'IDR' }
const MCC_LABEL: Record<string, string> = {
  '5172': 'Petroleum & Petroleum Products',
  '5411': 'Grocery Stores, Supermarkets',
  '5812': 'Eating Places, Restaurants',
  '5999': 'Miscellaneous Retail',
}

export const QRISResult = ({ data }: Props) => (
  <div style={styles.wrapper}>
    <ValidationBadge validation={data.validation} />

    <section style={styles.section}>
      <h3 style={styles.sectionTitle}>Merchant</h3>
      <DataRow label="Nama" value={data.merchant_name} />
      <DataRow label="Kota" value={data.merchant_city} />
      <DataRow label="Negara" value={data.country_code} />
      <DataRow label="MCC" value={data.mcc ? `${data.mcc} — ${MCC_LABEL[data.mcc] ?? 'Unknown'}` : undefined} />
    </section>

    <section style={styles.section}>
      <h3 style={styles.sectionTitle}>Transaksi</h3>
      <DataRow label="Currency" value={data.currency ? `${data.currency} (${CURRENCY_MAP[data.currency] ?? data.currency})` : undefined} />
      <DataRow label="Amount" value={data.amount ? formatAmount(data.amount) : 'Tidak ada (Static QR)'} />
      <DataRow label="Invoice" value={data.invoice} />
      <DataRow label="Terminal ID" value={data.terminal_id} />
    </section>

    <section style={styles.section}>
      <h3 style={styles.sectionTitle}>Acquirer / Merchant Account</h3>
      <DataRow label="Acquirer GUI" value={data.merchant.acquirer_gui} />
      <DataRow label="Merchant PAN" value={data.merchant.pan} />
      <DataRow label="Merchant ID" value={data.merchant.merchant_id} />
      <DataRow label="Criteria" value={data.merchant.criteria} />
    </section>

    <section style={styles.section}>
      <h3 style={styles.sectionTitle}>Integritas</h3>
      <DataRow label="CRC" value={data.crc} />
    </section>
  </div>
)

const formatAmount = (amount: string) => {
  const num = parseInt(amount, 10)
  if (isNaN(num)) return amount
  return `Rp ${num.toLocaleString('id-ID')}`
}

const styles: Record<string, React.CSSProperties> = {
  wrapper: { display: 'flex', flexDirection: 'column', gap: 20 },
  section: {
    background: '#0f172a',
    borderRadius: 8,
    padding: '12px 16px',
    border: '1px solid #1e293b',
  },
  sectionTitle: {
    margin: '0 0 8px',
    fontSize: 11,
    fontWeight: 700,
    color: '#3b82f6',
    textTransform: 'uppercase',
    letterSpacing: 1,
  },
}
