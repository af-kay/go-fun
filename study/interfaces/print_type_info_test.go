package interfaces

import "testing"

type person struct {
	name string
}

func (p person) String() string {
	return "Person " + p.name
}

func TestPrintTypeInfo(t *testing.T) {
    var i64 int64 = 64

	PrintTypeInfo("Hello world")
	PrintTypeInfo(i64)
	PrintTypeInfo(20) // TODO: just "int" support
	PrintTypeInfo(person{"Lara"})
}
