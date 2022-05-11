package Assignment

import (
	"fmt"
	"OLC2/Optimizar/Interfaces"
)

type AssignmentStack struct {
	Temp         	Interfaces.Expression
	Value        	Interfaces.Expression
	InstrDeleted 	bool
	Row          	int
	Col          	int
}

func GetStack(temp Interfaces.Expression, value Interfaces.Expression, row int, column int) AssignmentStack {
	return AssignmentStack{Value:value, Temp:temp, Row:row, Col:column}
}

func (a AssignmentStack) Execute() interface{} {

	pos  := a.Value.GetValue()
	temp := a.Temp.GetValue()
	
	return fmt.Sprintf("%v", temp) + " = stack[(int)" + fmt.Sprintf("%v", pos) + "];"
}
