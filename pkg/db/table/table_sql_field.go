package table

import (
	"fmt"
	"strings"
)

type SqlTableFieldDef struct {
	TableName string
	*TableFieldDef
}

type FieldCursor uint8

type SqlSequenceDef struct {
	RelativeField string
	SequenceName  string
}

type Field struct {
	Define string
	SqlSequenceDef
}

const (
	FieldCursorName       FieldCursor = iota // 字段名称
	FieldCursorType                          // 字段类型
	FieldCursorConstraint                    // 字段约束
)

func (s SqlTableFieldDef) String() string {
	ty, ct := s.detect()
	field := [...]string{
		FieldCursorName:       s.DefName,
		FieldCursorType:       ty,
		FieldCursorConstraint: ct,
	}
	return strings.Join(field[0:], " ")
}

func (s *SqlTableFieldDef) HasSequence() bool {
	ct := Constraint(s.DefDesc)
	return ct.IncludeSequence()
}

func (s *SqlTableFieldDef) detect() (string, string) {
	ct := Constraint(s.DefDesc)
	ty := s.detectVarchar(ct)
	xct := ct.String()
	if ct.IncludeSequence() {
		xct = s.attachNextVal(xct)
	}
	return ty, xct
}

func (f *SqlTableFieldDef) attachNextVal(ck string) string {
	return fmt.Sprintf("%s DEFAULT nextval('%s'::regclass)", ck, f.SequenceName())
}

func (f *SqlTableFieldDef) SequenceName() string {
	return fmt.Sprintf("%s_%s_seq", f.TableName, f.DefName)
}

func (s *SqlTableFieldDef) detectVarchar(ct Constraint) string {
	// convert golang type to postgresql type
	// golang type
	srcType := Lookup(s.DefType)
	// simple database type
	targetType := new(PGType).ToType(srcType)
	varchar := ct.HasSizeConstraint()
	if varchar != "" && srcType == Pstring {
		// change database text to varchar
		targetType = varchar
	}
	return targetType
}
