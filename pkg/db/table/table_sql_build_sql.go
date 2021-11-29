package table

import (
	"fmt"
	"math"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/token"
)

type TableSql struct {
	DBName    string
	tableName string
	Def       *DBTableDef
}

func NewTableSql(dbName, tableName string, dbdef *DBTableDef) *TableSql {
	return &TableSql{
		DBName:    dbName,
		tableName: tableName,
		Def:       dbdef,
	}
}

func (t *TableSql) Name() string {
	return fmt.Sprintf("t_%s", strings.ToLower(t.tableName))
}

func (t *TableSql) CreateTable() string {
	split := "\r\n\n"
	table := make([]string, 0)
	var alterSequence string
	// create fields
	table = append(table, t.Fields(func(s string, s2 string) {
		seq, as := t.Sequence(s, s2)
		if seq != "" {
			table = append(table, seq)
		}
		alterSequence = as
	}))
	// alter seq
	if alterSequence != "" {
		table = append(table, alterSequence)
	}
	// alter fields
	table = append(table, t.AlterColumn())
	// create index
	table = append(table, t.Index())
	return strings.Join(table, split)
}

func (t *TableSql) Fields(seq func(string, string)) string {
	fields := make([]string, 0)
	for i, _ := range t.Def.FieldDef {
		if t.Def.FieldDef[i] != nil {
			tmp := &SqlTableFieldDef{
				t.Name(),
				t.Def.FieldDef[i],
			}
			if tmp.HasSequence() {
				seq(t.Def.FieldDef[i].DefName, tmp.SequenceName())
			}
			fields = append(fields, tmp.String())
		}
	}
	// install primary constraint
	pk := t.PrimaryKeyConstraint()
	if pk != "" {
		fields = append(fields, pk)
	}
	define := strings.Join(fields, ",")
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ( %s );", t.Name(), define)
}

func (t *TableSql) Sequence(fieldName, seqName string) (cs string, as string) {
	cs = fmt.Sprintf("CREATE SEQUENCE IF NOT EXISTS %s INCREMENT 1 MINVALUE 1 MAXVALUE %d START 1 CACHE 1;", seqName, math.MaxInt64)
	as = fmt.Sprintf("ALTER SEQUENCE %s OWNED BY %s.%s;", seqName, t.Name(), fieldName)
	return
}

// Index /**
// CREATE [ UNIQUE ] INDEX [ CONCURRENTLY ] [ name ] ON table [ USING method ]
//    ( { column | ( expression ) } [ COLLATE collation ] [ opclass ] [ ASC | DESC ] [ NULLS { FIRST | LAST } ] [, ...] )
//    [ WITH ( storage_parameter = value [, ... ] ) ]
//    [ TABLESPACE tablespace ]
//    [ WHERE predicate ]

func (t *TableSql) Index() string {
	indexs := make([]string, 0)
	for i, _ := range t.Def.IndexDef {
		if t.Def.IndexDef[i] != nil {
			tmp := SqlTableIndexDef(*t.Def.IndexDef[i])
			ints := strings.Split(tmp.String(), "$")
			// default method btree
			indef := fmt.Sprintf("CREATE %s IF NOT EXISTS %s ON %s (%s);", ints[IndexCursorClass], ints[IndexCursorName], t.Name(), ints[IndexCursorFields])
			indexs = append(indexs, indef)
		}
	}
	return strings.Join(indexs, "\n")
}

func (t *TableSql) AlterColumn() string {
	fields := make([]string, 0)
	for i, _ := range t.Def.FieldDef {
		if t.Def.FieldDef[i] != nil {
			tmp := &SqlTableFieldDef{
				t.Name(),
				t.Def.FieldDef[i],
			}
			alert := fmt.Sprintf("ALTER TABLE %s ADD IF NOT EXISTS %s;", t.Name(), tmp.String())
			fields = append(fields, alert)
		}
	}
	return strings.Join(fields, "\n")
}

func (t *TableSql) PrimaryKeyConstraint() string {
	if t.Def.ConstraintDef != nil && len(t.Def.ConstraintDef) > 0 {
		// parser constraint def
		for _, key := range t.Def.ConstraintDef {
			if key.ConstraintKey == token.Primary.String() {
				p := PrimaryDef(key.ConstraintValues)
				return p.Primary()
			}
		}

	}
	return ""
}

func (t *TableSql) PrimaryKeys() []string {
	if t.Def.ConstraintDef != nil && len(t.Def.ConstraintDef) > 0 {
		// parser constraint def
		for _, key := range t.Def.ConstraintDef {
			if key.ConstraintKey == token.Primary.String() {
				p := PrimaryDef(key.ConstraintValues)
				return p
			}
		}

	}
	return nil
}
