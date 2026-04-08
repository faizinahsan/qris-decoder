# qris-decoder

Decoder untuk QRIS, menampilkan data tag dan namanya.

## Menjalankan dengan Docker

Pastikan Docker dan Docker Compose sudah terinstall.

```bash
# Jalankan semua service (backend + frontend)
docker compose up -d
```

Aplikasi tersedia di **http://localhost**

```bash
# Rebuild setelah ada perubahan code
docker compose up -d --build

# Lihat logs semua service
docker compose logs -f

# Lihat logs per service
docker compose logs -f backend
docker compose logs -f frontend

# Stop semua service
docker compose down
```

## Menjalankan Lokal (tanpa Docker)

```bash
# Backend (Go)
go run main.go

# Frontend (React) — di terminal terpisah
cd frontend && npm install && npm run dev
```

Backend berjalan di `localhost:8080`, frontend di `localhost:5173`.

## Arsitektur

```
Browser → nginx:80
           ├── /       → React static files
           └── /api/   → proxy → backend:8080
```


## Mapping Lengkap Tag QRIS (Production Reference)

### 1. Payload Level (Root TLV)
   | Tag   | Field Name                                        | Wajib    | Contoh          | Keterangan                |
   | ----- | ------------------------------------------------- | -------- | --------------- | ------------------------- |
   | 00    | Payload Format Indicator                          | Ya       | 01              | Versi format QR           |
   | 01    | Point of Initiation Method                        | Opsional | 11              | 11 = Dynamic, 12 = Static |
   | 02–25 | Reserved                                          | -        | -               | Tidak dipakai di QRIS     |
   | 26–45 | Merchant Account Information                      | Ya       | ID.CO.JALIN.WWW | Data acquirer / merchant  |
   | 46–51 | Merchant Account Information (additional network) | Opsional | ID.CO.QRIS.WWW  | Network tambahan          |
   | 52    | Merchant Category Code (MCC)                      | Ya       | 5812            | Kategori merchant         |
   | 53    | Transaction Currency                              | Ya       | 360             | 360 = IDR                 |
   | 54    | Transaction Amount                                | Dynamic  | 10000           | Nominal transaksi         |
   | 55    | Tip Indicator                                     | Opsional | 01              | Pengaturan tip            |
   | 56    | Convenience Fee Fixed                             | Opsional | 2000            | Tip fixed                 |
   | 57    | Convenience Fee Percentage                        | Opsional | 10              | Tip persen                |
   | 58    | Country Code                                      | Ya       | ID              | Negara                    |
   | 59    | Merchant Name                                     | Ya       | TOKO ABC        | Nama merchant             |
   | 60    | Merchant City                                     | Ya       | JAKARTA         | Kota merchant             |
   | 61    | Postal Code                                       | Opsional | 40123           | Kode pos                  |
   | 62    | Additional Data Field Template                    | Opsional | -               | Data tambahan transaksi   |
   | 63    | CRC                                               | Ya       | 4 hex           | Validasi QR               |

### 2. Merchant Account Information (Tag 26–51)

Ini bagian paling penting untuk routing transaksi.

| Subtag | Field                            | Contoh           | Keterangan              |
| ------ | -------------------------------- | ---------------- | ----------------------- |
| 00     | Globally Unique Identifier (GUI) | ID.CO.JALIN.WWW  | Acquirer / switch       |
| 01     | Merchant PAN                     | 9360xxxxxxxxxxxx | Identifier merchant     |
| 02     | Merchant ID / Terminal ID        | internal id      | ID merchant di acquirer |
| 03     | Merchant Criteria                | UMI / UKE / UBM  | Kategori usaha          |
| 04+    | Network specific                 | -                | Tergantung acquirer     |

### 3. Tip Indicator (Tag 55)

| Value | Arti           |
| ----- | -------------- |
| 01    | Tip optional   |
| 02    | Tip fixed      |
| 03    | Tip percentage |

### 4. Additional Data Field Template (Tag 62)

Sering dipakai untuk invoice / reference transaksi.

| Subtag | Field                    | Keterangan           |
| ------ | ------------------------ | -------------------- |
| 01     | Bill Number              | Nomor tagihan        |
| 02     | Mobile Number            | Nomor HP             |
| 03     | Store Label              | ID toko              |
| 04     | Loyalty Number           | Member ID            |
| 05     | Reference Label          | Reference transaksi  |
| 06     | Customer Label           | Customer ID          |
| 07     | Terminal Label           | ID terminal          |
| 08     | Purpose of Transaction   | Keterangan transaksi |
| 09     | Additional Consumer Data | Tambahan data        |
### 5. CRC (Tag 63)

Format:
```
63 04 XXXX
```

Menggunakan:
- **CRC-16/CCITT-FALSE**
- Polynomial: `0x1021`
- Initial: `0xFFFF`

Ini wajib valid agar QR bisa diproses oleh bank / wallet.

### 6. Field Penting untuk Backend Payment System

Biasanya sistem production hanya extract ini:

| Field        | Sumber Tag    |
| ------------ | ------------- |
| Acquirer     | 26.00         |
| MerchantPAN  | 26.01         |
| MerchantID   | 26.02         |
| MerchantType | 26.03         |
| MCC          | 52            |
| Currency     | 53            |
| Amount       | 54            |
| MerchantName | 59            |
| MerchantCity | 60            |
| Invoice      | 62.01 / 62.05 |
| TerminalID   | 62.07         |
| CRC          | 63            |
