package table

type TableSqlBuild struct {
	dbName    string // database name
	tableName string // target database table name
	filePath  string // database table description file path
}

func NewTableSqlBuild(dbName, tableName, filePath string) *TableSqlBuild {
	return &TableSqlBuild{
		dbName:    dbName,
		tableName: tableName,
		filePath:  filePath,
	}
}

func (t *TableSqlBuild) Build() *TableSql {
	return NewTableSql(t.dbName, t.tableName, Def(t.filePath))
}
