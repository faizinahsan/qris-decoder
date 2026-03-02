package main

import (
	"fmt"
	"strconv"
)

type TLV struct {
	Tag    string
	Length int
	Value  string
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

func main() {
	qr := "00020101021126670015ID.CO.JALIN.WWW011893600916237846693302151995432187654340303UMI51450015ID.CO.JALIN.WWW0215ID10190021351360303UMI5204581253033605802ID5918Merchant Jalin UAT6015Jakarta Selatan61051287262320303777070783218430810V4L1D4T1N663040CC5"

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

		// Merchant Account Info
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

		// Additional Data
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
}
