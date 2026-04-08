package main

import (
	"fmt"
	"os"

	"faizinahsan/qris-decoder/application/usecase"
	"faizinahsan/qris-decoder/domain/qris"
)

var rootTagName = map[string]string{
	"00": "Payload Format Indicator",
	"01": "Point of Initiation Method",
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

var merchantSubTagName = map[string]string{
	"00": "Globally Unique Identifier (Acquirer)",
	"01": "Merchant PAN",
	"02": "Merchant ID / Terminal ID",
	"03": "Merchant Criteria",
}

var additionalSubTagName = map[string]string{
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

func main() {
	raw := "00020101021151480018ID.CO.MINIMART.WWW0215ID10190020904070303UME5204517253033605802ID5910Minimarket6010Tanggerang61051290062070703A016304B38F"

	if len(os.Args) > 1 {
		raw = os.Args[1]
	}

	result, err := usecase.DecodeQRIS(raw)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	printSummary(result)
	printAllTags(result)
	printValidation(result)
}

func printSummary(r usecase.DecodeResult) {
	fmt.Println("===== AKSES DATA DARI MAP =====")
	fmt.Println("Merchant Name:", r.MerchantName)
	fmt.Println("Merchant City:", r.MerchantCity)
	fmt.Println("MCC          :", r.MCC)
	fmt.Println("Currency     :", r.Currency)
	if r.Amount != "" {
		fmt.Println("Amount       :", r.Amount)
	}
	fmt.Println("Acquirer     :", r.Acquirer)
	fmt.Println("Merchant PAN :", r.MerchantPAN)
	fmt.Println("Merchant ID  :", r.MerchantID)
	if r.Invoice != "" {
		fmt.Println("Invoice      :", r.Invoice)
	}
}

func printAllTags(r usecase.DecodeResult) {
	fmt.Println("\n===== QRIS TAG LIST =====")
	for tag, field := range r.AllFields {
		name := tagName(tag, field.Tag())
		fmt.Printf("\nTag %s (%s)\n", field.Tag(), name)
		fmt.Println("Length :", field.Length())
		fmt.Println("Value  :", field.Value())

		if subs, ok := r.AllSubFields[tag]; ok {
			fmt.Println(" Subtags:")
			printSubTags(tag, subs)
		}
	}
}

func printSubTags(parentTag int, subs map[int]qris.Field) {
	for _, sub := range subs {
		var name string
		if parentTag >= 26 && parentTag <= 51 {
			name = merchantSubTagName[sub.Tag()]
		} else if parentTag == 62 {
			name = additionalSubTagName[sub.Tag()]
		}
		if name == "" {
			name = "Network Specific"
		}
		fmt.Printf("  - %s (%s): %s\n", sub.Tag(), name, sub.Value())
	}
}

func printValidation(r usecase.DecodeResult) {
	fmt.Println("\n===== VALIDATION =====")
	v := r.Validation
	if v.Valid {
		fmt.Println("Status : VALID QRIS")
	} else {
		fmt.Println("Status : INVALID QRIS")
		for _, e := range v.Errors {
			fmt.Println("Error  :", e)
		}
	}
	fmt.Println("Acquirer GUI :", r.Acquirer)
	fmt.Println("Merchant PAN :", r.MerchantPAN)
	fmt.Println("Country      :", r.CountryCode)
	if v.IsCrossBorder {
		fmt.Println("QR Type      : CROSS BORDER")
	} else {
		fmt.Println("QR Type      : DOMESTIC")
	}
}

func tagName(tag int, tagStr string) string {
	if tag >= 26 && tag <= 45 {
		return "Merchant Account Information"
	}
	if tag >= 46 && tag <= 51 {
		return "Merchant Account Information (network)"
	}
	if name, ok := rootTagName[tagStr]; ok {
		return name
	}
	return "Unknown"
}
