package table

type PGType struct {
}

// pgType postgres filed type
var pgType = [...]string{
	Pint8:    "smallint", // 2 字节
	Pint16:   "smallint", // 2 字节
	Puint16:  "smallint", // 2 字节
	Pint32:   "integer",  // 4个字节
	Puint32:  "integer",  // 4个字节
	Pint64:   "bigint",   // 8个字节
	Puint64:  "bigint",   // 8个字节
	Pint:     "bigint",   // 8个字节
	Pfloat32: "decimal",  // 小数点前 131072 位；小数点后 16383 位
	Pfloat64: "decimal",  // 小数点前 131072 位；小数点后 16383 位
	Puint8:   "char(1)",  // 2 字节
	Pbyte:    "char(1)",  // 定长
	Pstring:  "text",     // 变长，无长度限制
}

func (pg *PGType) ToType(ty FieldType) string {
	if ty > 0 && ty < FieldType(len(pgType)) {
		return pgType[ty]
	}
	return ""
}
