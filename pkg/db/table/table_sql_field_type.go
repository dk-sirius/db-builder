package table

import (
	"strconv"
)

// FieldType FieldType 字段类型
type FieldType uint8

const (
	fieldBeg FieldType = iota
	PIndent
	Pint8
	Puint8
	Pint16
	Puint16
	Pint32
	Puint32
	Pint64
	Puint64
	Pint
	Pfloat32
	Pfloat64
	Pbyte
	Pstring
	fieldEnd
)

// fieldType golang type
var fieldType = [...]string{
	Pint8:    "int8",    // 2 字节
	Pint16:   "int16",   // 2 字节
	Puint16:  "uint16",  // 2 字节
	Pint32:   "int32",   // 4个字节
	Puint32:  "uint32",  // 4个字节
	Pint64:   "int64",   // 8个字节
	Puint64:  "uint64",  // 8个字节
	Pint:     "int",     // 8个字节
	Pfloat32: "float32", // 小数点前 131072 位；小数点后 16383 位
	Pfloat64: "float64", // 小数点前 131072 位；小数点后 16383 位
	Puint8:   "uint8",   // 2 字节
	Pbyte:    "byte",    // 定长
	Pstring:  "string",  // 变长，无长度限制
}

var fType map[string]FieldType

func init() {
	fType = make(map[string]FieldType)
	for i := fieldBeg + 1; i < fieldEnd; i++ {
		fType[fieldType[i]] = i
	}
}

func Lookup(ty string) FieldType {
	if tok, isKeyword := fType[ty]; isKeyword {
		return tok
	}
	return PIndent
}

func (tok FieldType) String() string {
	s := ""
	if 1 <= tok && tok < FieldType(len(fieldType)) {
		//s = field.pgFieldType[tok]  TODO
	}
	if s == "" {
		s = "FieldType(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}
