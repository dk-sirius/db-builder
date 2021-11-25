package token

type SchemaType uint8

const (
	SchemaTypePg SchemaType = iota
)

func (s SchemaType) String() string {
	switch s {
	case SchemaTypePg:
		return "postgres"
	}
	return ""
}
