package Control


import (
	// "fmt"
	"OLC2/Optimizar/Interfaces"
)

type Return struct {
	Value 		Interfaces.Expression
	Row       	int
	Col       	int
}

func NewReturn(val Interfaces.Expression, row int, column int) Return {
	return Return{val, row, column}
}

func (p Return) Execute() interface{} {

	if p.Value != nil {
		return "return;"
		// return "return " + fmt.Sprintf("%v", p.Value.GetValue()) + ";"
	}

	return "return;"
}
	
