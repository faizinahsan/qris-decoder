package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TLV struct {
	Tag    string
	Length int
	Value  string
}

type QRISData struct {
	Tags    map[int]string         // Root tags: key = tag, value = value
	SubTags map[int]map[int]string // Subtags: key = parent tag, value = map of subtag
}

var rootTagName = map[string]string{
	"00": "Payload Format Indicator",
	"01": "Point of Initiation Method",
	"26": "Merchant Account Information",
	"27": "Merchant Account Information",
	"28": "Merchant Account Information",
	"29": "Merchant Account Information",
	"30": "Merchant Account Information",
	"31": "Merchant Account Information",
	"32": "Merchant Account Information",
	"33": "Merchant Account Information",
	"34": "Merchant Account Information",
	"35": "Merchant Account Information",
	"36": "Merchant Account Information",
	"37": "Merchant Account Information",
	"38": "Merchant Account Information",
	"39": "Merchant Account Information",
	"40": "Merchant Account Information",
	"41": "Merchant Account Information",
	"42": "Merchant Account Information",
	"43": "Merchant Account Information",
	"44": "Merchant Account Information",
	"45": "Merchant Account Information",
	"46": "Merchant Account Information (network)",
	"47": "Merchant Account Information (network)",
	"48": "Merchant Account Information (network)",
	"49": "Merchant Account Information (network)",
	"50": "Merchant Account Information (network)",
	"51": "Merchant Account Information (network)",
	"52": "Merchant Category Code",
	"53": "Transaction Currency",
	"54": "Transaction Amount",
	"55": "Tip Indicator",
	"56": "Convenience Fee Fixed",
	"57": "Convenience Fee Percentage",
	"58": "Country Code",
	"59": "Merchant Name",
	"60": "Merchant City",
	"61": "Postal Code",
	"62": "Additional Data Field Template",
	"63": "CRC",
}

var merchantSubTag = map[string]string{
	"00": "Globally Unique Identifier (Acquirer)",
	"01": "Merchant PAN",
	"02": "Merchant ID / Terminal ID",
	"03": "Merchant Criteria",
	"04": "Network Specific",
}

var additionalSubTag = map[string]string{
	"01": "Bill Number",
	"02": "Mobile Number",
	"03": "Store Label",
	"04": "Loyalty Number",
	"05": "Reference Label",
	"06": "Customer Label",
	"07": "Terminal Label",
	"08": "Purpose of Transaction",
	"09": "Additional Consumer Data",
}

func parseTLV(data string) ([]TLV, error) {
	var result []TLV
	i := 0

	for i < len(data) {
		if i+4 > len(data) {
			break
		}

		tag := data[i : i+2]
		lengthStr := data[i+2 : i+4]

		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return nil, err
		}

		start := i + 4
		end := start + length
		if end > len(data) {
			break
		}

		value := data[start:end]

		result = append(result, TLV{
			Tag:    tag,
			Length: length,
			Value:  value,
		})

		i = end
	}

	return result, nil
}

func parseQRISToMap(qr string) (*QRISData, error) {
	tlvs, err := parseTLV(qr)
	if err != nil {
		return nil, err
	}

	data := &QRISData{
		Tags:    make(map[int]string),
		SubTags: make(map[int]map[int]string),
	}

	for _, tlv := range tlvs {
		tag, _ := strconv.Atoi(tlv.Tag)
		data.Tags[tag] = tlv.Value

		if isMerchantAccountTag(tlv.Tag) || tlv.Tag == "62" {
			sub := parseSubTLV(tlv.Value)
			data.SubTags[tag] = make(map[int]string)
			for _, s := range sub {
				subTag, _ := strconv.Atoi(s.Tag)
				data.SubTags[tag][subTag] = s.Value
			}
		}
	}

	return data, nil
}

func parseSubTLV(data string) []TLV {
	var result []TLV
	i := 0

	for i < len(data) {
		if i+4 > len(data) {
			break
		}

		tag := data[i : i+2]
		lengthStr := data[i+2 : i+4]
		length, _ := strconv.Atoi(lengthStr)

		start := i + 4
		end := start + length
		if end > len(data) {
			break
		}

		value := data[start:end]

		result = append(result, TLV{
			Tag:    tag,
			Length: length,
			Value:  value,
		})

		i = end
	}

	return result
}

func isMerchantAccountTag(tag string) bool {
	t, _ := strconv.Atoi(tag)
	return t >= 26 && t <= 51
}

var mandatoryTags = map[string]bool{
	"00": true,
	"52": true,
	"53": true,
	"58": true,
	"59": true,
	"60": true,
	"63": true,
}

func crc16CCITT(data string) uint16 {
	var crc uint16 = 0xFFFF

	for i := 0; i < len(data); i++ {
		crc ^= uint16(data[i]) << 8
		for j := 0; j < 8; j++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}

	return crc
}

func validateMandatory(tags map[string]bool) []string {
	var errors []string
	for tag := range mandatoryTags {
		if !tags[tag] {
			errors = append(errors, "missing mandatory tag "+tag)
		}
	}
	return errors
}

func main() {
	qr := "00020101021126670015ID.CO.JALIN.WWW011893600916237846693302151995432187654340303UMI51450015ID.CO.JALIN.WWW0215ID10190021351360303UMI5204581253033605802ID5918Merchant Jalin UAT6015Jakarta Selatan61051287262320303777070783218430810V4L1D4T1N663040CC5"

	// Parse ke map
	qrisData, err := parseQRISToMap(qr)
	if err != nil {
		panic(err)
	}

	// Contoh akses data langsung dari map
	fmt.Println("===== AKSES DATA DARI MAP =====")
	fmt.Println("Merchant Name:", qrisData.Tags[59])
	fmt.Println("Merchant City:", qrisData.Tags[60])
	fmt.Println("MCC          :", qrisData.Tags[52])
	fmt.Println("Currency     :", qrisData.Tags[53])
	if amount, ok := qrisData.Tags[54]; ok {
		fmt.Println("Amount       :", amount)
	}
	if subtags, ok := qrisData.SubTags[26]; ok {
		fmt.Println("Acquirer     :", subtags[0])
		fmt.Println("Merchant PAN :", subtags[1])
		fmt.Println("Merchant ID  :", subtags[2])
	}
	if subtags, ok := qrisData.SubTags[62]; ok {
		if invoice, ok := subtags[5]; ok {
			fmt.Println("Invoice      :", invoice)
		}
	}

	tlvs, err := parseTLV(qr)
	if err != nil {
		panic(err)
	}

	fmt.Println("===== QRIS TAG LIST =====")

	for _, tlv := range tlvs {
		name := rootTagName[tlv.Tag]
		if name == "" {
			name = "Unknown"
		}

		fmt.Printf("\nTag %s (%s)\n", tlv.Tag, name)
		fmt.Println("Length :", tlv.Length)
		fmt.Println("Value  :", tlv.Value)

		if isMerchantAccountTag(tlv.Tag) {
			sub := parseSubTLV(tlv.Value)
			fmt.Println(" Subtags:")
			for _, s := range sub {
				subName := merchantSubTag[s.Tag]
				if subName == "" {
					subName = "Network Specific"
				}
				fmt.Printf("  - %s (%s): %s\n", s.Tag, subName, s.Value)
			}
		}

		if tlv.Tag == "62" {
			sub := parseSubTLV(tlv.Value)
			fmt.Println(" Subtags:")
			for _, s := range sub {
				subName := additionalSubTag[s.Tag]
				if subName == "" {
					subName = "Additional Field"
				}
				fmt.Printf("  - %s (%s): %s\n", s.Tag, subName, s.Value)
			}
		}
	}

	// VALIDATION
	validateQRIS(qr, tlvs)
}

func validateQRIS(qr string, tlvs []TLV) {
	fmt.Println("\n===== VALIDATION =====")

	found := map[string]TLV{}

	var gui string
	var merchantPAN string
	var country string
	var crc string
	var tipIndicator string

	for _, t := range tlvs {
		found[t.Tag] = t

		if isMerchantAccountTag(t.Tag) {
			sub := parseSubTLV(t.Value)
			for _, s := range sub {
				if s.Tag == "00" {
					gui = s.Value
				}
				if s.Tag == "01" {
					merchantPAN = s.Value
				}
			}
		}

		if t.Tag == "58" {
			country = t.Value
		}

		if t.Tag == "55" {
			tipIndicator = t.Value
		}

		if t.Tag == "63" {
			crc = t.Value
		}
	}

	var errors []string

	// Mandatory tag
	mandatory := []string{"00", "01", "52", "53", "58", "59", "60", "63"}
	for _, m := range mandatory {
		if _, ok := found[m]; !ok {
			errors = append(errors, "missing mandatory tag "+m)
		}
	}

	// CRC validation
	idx := strings.Index(qr, "6304")
	if idx == -1 {
		errors = append(errors, "CRC tag not found")
	} else {
		payload := qr[:idx+4]
		calculated := fmt.Sprintf("%04X", crc16CCITT(payload))
		if calculated != crc {
			errors = append(errors, "CRC mismatch")
		}
	}

	// Merchant PAN validation
	if merchantPAN == "" {
		errors = append(errors, "merchant PAN not found")
	}

	if len(merchantPAN) < 10 {
		errors = append(errors, "merchant PAN too short")
	}

	// GUI validation
	if gui == "" {
		errors = append(errors, "acquirer GUI not found")
	}

	// Country validation
	if country == "" {
		errors = append(errors, "country code missing")
	}

	// Cross Border
	crossBorder := false
	if country != "ID" {
		crossBorder = true
	}
	if !strings.HasPrefix(gui, "ID.") {
		crossBorder = true
	}

	// Result
	if len(errors) == 0 {
		fmt.Println("Status : VALID QRIS")
	} else {
		fmt.Println("Status : INVALID QRIS")
		for _, e := range errors {
			fmt.Println("Error  :", e)
		}
	}

	fmt.Println("Acquirer GUI :", gui)
	fmt.Println("Merchant PAN :", merchantPAN)
	fmt.Println("Country      :", country)

	if tipIndicator != "" {
		fmt.Println("Tip Indicator:", tipIndicator)
	}

	if crossBorder {
		fmt.Println("QR Type      : CROSS BORDER")
	} else {
		fmt.Println("QR Type      : DOMESTIC")
	}
}
