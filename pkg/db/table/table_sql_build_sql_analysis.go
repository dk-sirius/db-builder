package table

type Analysis struct {
	TableName           string
	PrimaryFieldKey     []string
	FieldKey            []string
	UniqueIndexFieldKey []map[string]string
	IndexFieldKey       []map[string]string
}

func (t *TableSql) AnalysisTable() *Analysis {
	result := &Analysis{}
	result.TableName = t.Name()

	// fields
	for i, _ := range t.Def.FieldDef {
		if t.Def.FieldDef[i] != nil {
			result.FieldKey = append(result.FieldKey, t.Def.FieldDef[i].DefName)
		}
	}
	pks := t.PrimaryKeys()
	if pks != nil {
		result.PrimaryFieldKey = append(result.PrimaryFieldKey, pks...)
	}
	// index
	for i, _ := range t.Def.IndexDef {
		tmp := SqlTableIndexDef(*t.Def.IndexDef[i])
		desc := tmp.Index()
		if desc[IndexCursorClass] != "" {
			ui := make(map[string]string)
			ui[desc[IndexCursorName]] = desc[IndexCursorFields]
			if tmp.IsUnique(t.Def.IndexDef[i].DefClass) {
				result.UniqueIndexFieldKey = append(result.UniqueIndexFieldKey, ui)
			} else {
				result.IndexFieldKey = append(result.IndexFieldKey, ui)
			}
		}
	}
	return result
}
