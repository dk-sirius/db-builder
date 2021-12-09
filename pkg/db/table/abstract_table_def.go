package table

/**
    // default each field value is not null

	// comments
    // @def constraintKey  constraintValues ...
    // @def index indexName RelativeField ....
    // @def unique_index indexName RelativeField ...
	type FieldDef struct {
		Field1 type `db:"fieldName,constraint(size=10\default=''...)"`   // see detail token.DbToken
		....
    }
*/

// DBTableDef database table define
type DBTableDef struct {
	// file name of picked
	Name string
	// all valid field define
	FieldDef []*FieldDef
	// all valid table index define
	IndexDef []*IndexDef
	// all valid doc constraint define
	ConstraintDef []*DocPlaceHolderDef
}
