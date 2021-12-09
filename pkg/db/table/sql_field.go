package table

import (
	"strings"
)

type SqlFieldDef FieldDef

type FieldCursor uint8

const (
	FieldCursorName       FieldCursor = iota // 字段名称
	FieldCursorType                          // 字段类型
	FieldCursorConstraint                    // 字段约束
)

func (s SqlFieldDef) String() string {
	attachment := NewFieldCheck(s.DefDesc).Handle()
	field := [...]string{
		FieldCursorName:       s.DefName,
		FieldCursorType:       s.sqlType(attachment),
		FieldCursorConstraint: attachment.CheckType,
	}
	return strings.Join(field[0:], " ")
}

func (s SqlFieldDef) IsAuto() bool {
	return NewFieldCheck(s.DefDesc).Handle().IsAuto
}

func (s *SqlFieldDef) sqlType(att *FieldCheck) string {
	aft := Lookup(s.DefType)
	switch aft {
	case String:
		if att.Size != 0 {
			return aft.Varchar(att.Size)
		}
	case Int, Int32, Int64, Uint64, Uint32:
		if att.IsAuto {
			return aft.SqlAutoType()
		}
	}
	return aft.SqlType()
}
