package parser

import (
	"strconv"

	"faizinahsan/qris-decoder/domain/qris"
)

// Parse mengubah raw QRIS string menjadi QRISPayload aggregate.
// Ini adalah satu-satunya tempat yang boleh tahu tentang format TLV.
func Parse(raw string) (qris.QRISPayload, error) {
	fields, err := parseTLV(raw)
	if err != nil {
		return qris.QRISPayload{}, qris.ErrInvalidFormat
	}

	fieldMap := make(map[int]qris.Field)
	subFieldMap := make(map[int]map[int]qris.Field)

	for _, f := range fields {
		tag, _ := strconv.Atoi(f.Tag())
		fieldMap[tag] = f

		if isMerchantTag(tag) || tag == 62 {
			subs, err := parseTLV(f.Value())
			if err != nil {
				continue
			}
			subFieldMap[tag] = make(map[int]qris.Field)
			for _, s := range subs {
				subTag, _ := strconv.Atoi(s.Tag())
				subFieldMap[tag][subTag] = s
			}
		}
	}

	merchant := buildMerchant(subFieldMap)
	crc, _ := fieldMap[63]

	return qris.NewQRISPayload(raw, fieldMap, subFieldMap, merchant, crc.Value()), nil
}

func parseTLV(data string) ([]qris.Field, error) {
	var result []qris.Field
	i := 0
	for i < len(data) {
		if i+4 > len(data) {
			break
		}
		tag := data[i : i+2]
		length, err := strconv.Atoi(data[i+2 : i+4])
		if err != nil {
			return nil, err
		}
		end := i + 4 + length
		if end > len(data) {
			break
		}
		result = append(result, qris.NewField(tag, length, data[i+4:end]))
		i = end
	}
	return result, nil
}

func buildMerchant(subFields map[int]map[int]qris.Field) qris.Merchant {
	// Cari merchant account dari tag 26–51
	for tag := 26; tag <= 51; tag++ {
		subs, ok := subFields[tag]
		if !ok {
			continue
		}
		gui := subFieldValue(subs, 0)
		pan := subFieldValue(subs, 1)
		mid := subFieldValue(subs, 2)
		criteria := subFieldValue(subs, 3)
		if gui != "" {
			return qris.NewMerchant(gui, pan, mid, criteria)
		}
	}
	return qris.NewMerchant("", "", "", "")
}

func subFieldValue(subs map[int]qris.Field, tag int) string {
	if f, ok := subs[tag]; ok {
		return f.Value()
	}
	return ""
}

func isMerchantTag(tag int) bool {
	return tag >= 26 && tag <= 51
}
