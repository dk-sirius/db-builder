package gen

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"
)

type ExprBody struct {
	Data interface{}
	Keep bool
}

func (ex *ExprBody) Type(value interface{}) interface{} {
	if reflect.TypeOf(value).Kind() == reflect.String && !ex.Keep {
		return fmt.Sprintf("\"%s\"", value)
	}
	return value
}
func (ex *ExprBody) MapExpr() string {
	ret := reflect.TypeOf(ex.Data).String()
	elements := `{{- range $k,$v:=.Data}}"{{$k}}":{{$.Type $v}},{{- end}}`
	return fmt.Sprintf("%s{ %s }", ret, elements)
}
func (ex *ExprBody) SliceExpr() string {
	ret := reflect.TypeOf(ex.Data).String()
	elements := `{{- range $index,$v:=.Data}}{{$.Type $v}},{{- end}}`
	return fmt.Sprintf("%s{ %s }", ret, elements)
}

func (ex *ExprBody) BasicExpr() string {
	return `{{.Type .Data}}`
}

func (ex *ExprBody) Tmpl() string {
	switch reflect.TypeOf(ex.Data).Kind() {
	case reflect.Map:
		return ex.MapExpr()
	case reflect.Slice:
		return ex.SliceExpr()
	case reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Float32,
		reflect.Float64,
		reflect.String:
		return ex.BasicExpr()
	default:
		panic("not support type")
	}
	return ""
}

func (ex *ExprBody) Exec() string {
	tm := ex.Tmpl()
	tmpl, err := template.New("body").Parse(tm)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, ex)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
