package qris

// Merchant adalah Value Object: merepresentasikan data merchant dari subtag 26–51.
// Immutable setelah dibuat.
type Merchant struct {
	acquirerGUI string
	pan         string
	merchantID  string
	criteria    string
}

func NewMerchant(gui, pan, merchantID, criteria string) Merchant {
	return Merchant{
		acquirerGUI: gui,
		pan:         pan,
		merchantID:  merchantID,
		criteria:    criteria,
	}
}

func (m Merchant) AcquirerGUI() string { return m.acquirerGUI }
func (m Merchant) PAN() string         { return m.pan }
func (m Merchant) MerchantID() string  { return m.merchantID }
func (m Merchant) Criteria() string    { return m.criteria }

func (m Merchant) IsDomestic() bool {
	return len(m.acquirerGUI) >= 3 && m.acquirerGUI[:3] == "ID."
}

func (m Merchant) IsValid() bool {
	// PAN opsional di beberapa acquirer, GUI wajib ada
	return m.acquirerGUI != ""
}
