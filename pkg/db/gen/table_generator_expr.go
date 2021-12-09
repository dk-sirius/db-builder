package gen

import (
	"bytes"
	"reflect"
	"text/template"
)

type MethodExpr struct {
	MethodName   string
	ReceiverName string
	Data         interface{}
	Keep         bool
}

func NewExpr(methodName, receiverName string, data interface{}) *MethodExpr {
	return &MethodExpr{
		MethodName:   methodName,
		ReceiverName: receiverName,
		Data:         data,
	}
}

func (m *MethodExpr) ToKeep(flag bool) *MethodExpr {
	m.Keep = flag
	return m
}

func (m *MethodExpr) Body() interface{} {
	body := &ExprBody{
		Data: m.Data,
		Keep: m.Keep,
	}
	re := body.Exec()
	return re
}

func (m *MethodExpr) Ret() string {
	return reflect.TypeOf(m.Data).String()
}

func (m *MethodExpr) Method() string {
	return `func ({{.ReceiverName}}){{.MethodName}}(){{.Ret}}{
    	return {{.Body}}
	}`
}

func (m *MethodExpr) Gen() string {
	tmpl, err := template.New("gen").Parse(m.Method())
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, m)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
