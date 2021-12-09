package table

type SqlBuilder struct {
	dbName    string // database name
	tableName string // target database table name
	filePath  string // database table description file path
}

func NewSqlBuilder(dbName, tableName, filePath string) *SqlBuilder {
	return &SqlBuilder{
		dbName:    dbName,
		tableName: tableName,
		filePath:  filePath,
	}
}

func (t *SqlBuilder) Build() *SqlSchema {
	return NewSqlSchema(t.dbName, t.tableName, Def(t.filePath))
}
