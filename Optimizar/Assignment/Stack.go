package Assignment

import (
	"fmt"
	"OLC2/Optimizar/Interfaces"
)

type Stack struct {
	Value        	Interfaces.Expression
	Temp         	Interfaces.Expression
	InstrDeleted 	bool
	Row          	int
	Col          	int
}

func NewStack(temp Interfaces.Expression, value Interfaces.Expression, row int, column int) Stack {
	return Stack{Value:value, Temp:temp, Row:row, Col:column}
}

func (a Stack) Execute() interface{} {

	value  := a.Value.GetValue()
	temp := a.Temp.GetValue()
	

	return "stack[(int)" + fmt.Sprintf("%v",temp) + "] = " + fmt.Sprintf("%v", value) + ";"
}
