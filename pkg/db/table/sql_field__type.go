package table

import (
	"fmt"
	"strconv"
)

// AstFieldType  字段类型
type AstFieldType uint8

const (
	fieldBeg AstFieldType = iota
	Indent
	Int8
	Uint8
	Int16
	Uint16
	Int32
	Uint32
	Int64
	Uint64
	Int
	Float32
	Float64
	Byte
	String
	fieldEnd
)

var (
	// fieldType golang type
	fieldType = [...]string{
		Int8:    "int8",    // 2 字节
		Int16:   "int16",   // 2 字节
		Uint16:  "uint16",  // 2 字节
		Int32:   "int32",   // 4个字节
		Uint32:  "uint32",  // 4个字节
		Int64:   "int64",   // 8个字节
		Uint64:  "uint64",  // 8个字节
		Int:     "int",     // 8个字节
		Float32: "float32", // 小数点前 131072 位；小数点后 16383 位
		Float64: "float64", // 小数点前 131072 位；小数点后 16383 位
		Uint8:   "uint8",   // 2 字节
		Byte:    "byte",    // 定长
		String:  "string",  // 变长，无长度限制
	}

	// postgresType postgres filed type
	postgresType = [...]string{
		Int8:    "smallint", // 2 字节
		Int16:   "smallint", // 2 字节
		Uint16:  "smallint", // 2 字节
		Int32:   "integer",  // 4个字节
		Uint32:  "integer",  // 4个字节
		Int64:   "bigint",   // 8个字节
		Uint64:  "bigint",   // 8个字节
		Int:     "bigint",   // 8个字节
		Float32: "decimal",  // 小数点前 131072 位；小数点后 16383 位
		Float64: "decimal",  // 小数点前 131072 位；小数点后 16383 位
		Uint8:   "char(1)",  // 2 字节
		Byte:    "char(1)",  // 定长
		String:  "text",     // 变长，无长度限制
	}
)

var fieldTypeMap map[string]AstFieldType

func init() {
	fieldTypeMap = make(map[string]AstFieldType)
	for i := fieldBeg + 1; i < fieldEnd; i++ {
		fieldTypeMap[fieldType[i]] = i
	}
}

func Lookup(ty string) AstFieldType {
	if tok, exist := fieldTypeMap[ty]; exist {
		return tok
	}
	return Indent
}

func (ft AstFieldType) String() string {
	s := ""
	if 1 <= ft && ft < AstFieldType(len(fieldType)) {
		s = fieldType[ft]
	}
	if s == "" {
		panic("AstFieldType(" + strconv.Itoa(int(ft)) + ")")
	}
	return s
}

func (AstFieldType) SqlAutoType() string {
	return "bigserial"
}

func (ft AstFieldType) SqlType() string {
	if ft > 0 && ft < AstFieldType(len(postgresType)) {
		return postgresType[ft]
	}
	return ""
}

func (ft AstFieldType) Varchar(size int64) string {
	if ft == String {
		return fmt.Sprintf("varchar(%d)", size)
	}
	return ft.String()
}
