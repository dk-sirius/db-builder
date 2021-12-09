package table

import (
	"fmt"
	"strings"

	"github.com/dk-sirius/db-builder/pkg/db/token"
)

type SqlSchema struct {
	DBName    string
	tableName string
	Def       *DBTableDef
}

func NewSqlSchema(dbName, tableName string, dbdef *DBTableDef) *SqlSchema {
	return &SqlSchema{
		DBName:    dbName,
		tableName: tableName,
		Def:       dbdef,
	}
}

func (t *SqlSchema) Name() string {
	return fmt.Sprintf("t_%s", strings.ToLower(t.tableName))
}

func (t *SqlSchema) CreateTable() string {
	split := "\r\n\n"
	table := make([]string, 0)
	// create fields
	table = append(table, t.Fields())
	// alter fields
	table = append(table, t.AlterColumn())
	// create index
	table = append(table, t.Index())
	return strings.Join(table, split)
}

func (t *SqlSchema) Fields() string {
	fields := make([]string, 0)
	for i, _ := range t.Def.FieldDef {
		if t.Def.FieldDef[i] != nil {
			tmp := SqlFieldDef(*t.Def.FieldDef[i])
			fields = append(fields, tmp.String())
		}
	}
	// install primary constraint
	pk := t.PrimaryKey()
	if pk != "" {
		fields = append(fields, pk)
	}
	define := strings.Join(fields, ",")
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ( %s );", t.Name(), define)
}

// Index /**
// CREATE [ UNIQUE ] INDEX [ CONCURRENTLY ] [ name ] ON table [ USING method ]
//    ( { column | ( expression ) } [ COLLATE collation ] [ opclass ] [ ASC | DESC ] [ NULLS { FIRST | LAST } ] [, ...] )
//    [ WITH ( storage_parameter = value [, ... ] ) ]
//    [ TABLESPACE tablespace ]
//    [ WHERE predicate ]

func (t *SqlSchema) Index() string {
	indexs := make([]string, 0)
	for i, _ := range t.Def.IndexDef {
		if t.Def.IndexDef[i] != nil {
			tmp := SqlIndexDef(*t.Def.IndexDef[i])
			ints := strings.Split(tmp.String(), "$")
			// default method btree
			indef := fmt.Sprintf("CREATE %s IF NOT EXISTS %s ON %s (%s);", ints[IndexCursorClass], ints[IndexCursorName], t.Name(), ints[IndexCursorFields])
			indexs = append(indexs, indef)
		}
	}
	return strings.Join(indexs, "\n")
}

func (t *SqlSchema) AlterColumn() string {
	fields := make([]string, 0)
	for i, _ := range t.Def.FieldDef {
		if t.Def.FieldDef[i] != nil {
			tmp := SqlFieldDef(*t.Def.FieldDef[i])
			alert := fmt.Sprintf("ALTER TABLE %s ADD IF NOT EXISTS %s;", t.Name(), tmp.String())
			fields = append(fields, alert)
		}
	}
	return strings.Join(fields, "\n")
}

func (t *SqlSchema) PrimaryKey() string {
	if t.Def.ConstraintDef != nil && len(t.Def.ConstraintDef) > 0 {
		// parser constraint def
		for _, key := range t.Def.ConstraintDef {
			if key.PlaceHolderKey == token.Primary.String() {
				p := PrimaryDef(key.PlaceHolderValues)
				return p.Primary()
			}
		}

	}
	return ""
}

func (t *SqlSchema) PrimaryKeyValues() []string {
	if t.Def.ConstraintDef != nil && len(t.Def.ConstraintDef) > 0 {
		// parser constraint def
		for _, key := range t.Def.ConstraintDef {
			if key.PlaceHolderKey == token.Primary.String() {
				p := PrimaryDef(key.PlaceHolderValues)
				return p
			}
		}

	}
	return nil
}
