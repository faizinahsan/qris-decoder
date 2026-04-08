package qris

// Field adalah Value Object: immutable, tidak punya identity unik.
// Merepresentasikan satu segmen TLV (Tag-Length-Value) dalam payload QRIS.
type Field struct {
	tag    string
	length int
	value  string
}

func NewField(tag string, length int, value string) Field {
	return Field{tag: tag, length: length, value: value}
}

func (f Field) Tag() string   { return f.tag }
func (f Field) Length() int   { return f.length }
func (f Field) Value() string { return f.value }
