package Printf

import (
	"fmt"
	"OLC2/Optimizar/Interfaces"
)

type Printf struct {
	Value  	Interfaces.Expression
	IsType  string
	Type   	string
	Row    	int
	Col    	int
}

func NewPrintf(value Interfaces.Expression, isType string, tipo string, row int, column int) Printf {
	return Printf{value, isType, tipo, row, column}
}


func (p Printf) Execute() interface{} {

	var val string

	if p.Type != "-1" {
		val = "printf(" + p.IsType + "," + "("+ p.Type + ")" + fmt.Sprintf("%v",p.Value.GetValue()) + ");"	

	} else {
		val ="printf(" + p.IsType + "," + fmt.Sprintf("%v",p.Value.GetValue()) + ");"	

	}

	return val
}
