package gen_test

import (
	"fmt"
	"testing"

	"github.com/dk-sirius/db-builder/pkg/db/gen"
)

func TestNewExpr(t *testing.T) {
	//a := map[string]bool{
	//	"a": true,
	//	"b": true,
	//}
	//b := []string{"a", "b", "c"}
	c := "true"
	k := gen.NewExpr("mmmStr4", "Student", c)
	fmt.Println(k.Gen())
}

type Student struct {
	Name string
}

func (Student) mmm1() map[string]string {
	return map[string]string{"a": "true", "b": "true"}
}
func (Student) mmm2() map[string]bool {
	return map[string]bool{"a": true, "b": true}
}

func (Student) mmm() []string {
	return []string{"a", "b", "c"}
}

func (Student) mmmStr() map[string]bool {
	return map[string]bool{"a": true, "b": true}
}
func (Student) mmmStr1() []string {
	return []string{"a", "b", "c"}
}
func (Student) mmmStr2() bool {
	return true
}
func (Student) mmmStr4() string {
	return "true"
}
