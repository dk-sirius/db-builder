package table

type Analysis struct {
	TableName           string
	PrimaryFieldKey     []SqlFieldName
	FieldKey            []SqlFieldName
	UniqueIndexFieldKey []map[IndexName]IndexMapFields
	IndexFieldKey       []map[IndexName]IndexMapFields
	AutoFieldKey        map[SqlFieldName]bool
}

func (t *SqlSchema) AnalysisTable() *Analysis {
	result := &Analysis{}
	result.TableName = t.Name()
	result.AutoFieldKey = make(map[SqlFieldName]bool)
	// fields
	for i, _ := range t.Def.FieldDef {
		if t.Def.FieldDef[i] != nil {
			result.FieldKey = append(result.FieldKey, t.Def.FieldDef[i].DefName)
			tmp := SqlFieldDef(*t.Def.FieldDef[i])
			if tmp.IsAuto() {
				result.AutoFieldKey[t.Def.FieldDef[i].DefName] = true
			}
		}
	}
	pks := t.PrimaryKeyValues()
	if pks != nil {
		result.PrimaryFieldKey = append(result.PrimaryFieldKey, pks...)
	}
	// index
	for i, _ := range t.Def.IndexDef {
		tmp := SqlIndexDef(*t.Def.IndexDef[i])
		desc := tmp.Index()
		if desc[IndexCursorClass] != "" {
			ui := make(map[IndexName]IndexMapFields)
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
