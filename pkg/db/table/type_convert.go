package table

// TypeConvert define to adapter different db
type TypeConvert interface {
	ToType(ty FieldType) string
}
