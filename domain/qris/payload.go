package qris

// QRISPayload adalah Aggregate Root.
// Semua akses ke data QRIS harus melalui struct ini.
// Tidak boleh ada logic bisnis di luar domain layer yang memanipulasi field ini langsung.
type QRISPayload struct {
	raw       string
	fields    map[int]Field
	subFields map[int]map[int]Field
	merchant  Merchant
	crc       string
}

func NewQRISPayload(
	raw string,
	fields map[int]Field,
	subFields map[int]map[int]Field,
	merchant Merchant,
	crc string,
) QRISPayload {
	return QRISPayload{
		raw:       raw,
		fields:    fields,
		subFields: subFields,
		merchant:  merchant,
		crc:       crc,
	}
}

func (q QRISPayload) Raw() string                      { return q.raw }
func (q QRISPayload) Merchant() Merchant               { return q.merchant }
func (q QRISPayload) CRC() string                      { return q.crc }
func (q QRISPayload) Fields() map[int]Field            { return q.fields }
func (q QRISPayload) SubFields() map[int]map[int]Field { return q.subFields }

func (q QRISPayload) TagValue(tag int) (string, bool) {
	f, ok := q.fields[tag]
	return f.Value(), ok
}

func (q QRISPayload) SubTagValue(parent, sub int) (string, bool) {
	subs, ok := q.subFields[parent]
	if !ok {
		return "", false
	}
	f, ok := subs[sub]
	return f.Value(), ok
}

// MandatoryTags mendefinisikan tag yang wajib ada di setiap QRIS payload.
// Ini adalah business rule, tempatnya di domain.
var MandatoryTags = []int{0, 1, 52, 53, 58, 59, 60, 63}
