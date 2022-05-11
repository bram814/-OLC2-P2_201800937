package Assignment

import (
	"fmt"
	"OLC2/Optimizar/Interfaces"
)

type AssignmentHeap struct {
	Temp         	Interfaces.Expression
	Value        	Interfaces.Expression
	InstrDeleted 	bool
	Row          	int
	Col          	int
}

func GetHeap(temp Interfaces.Expression, value Interfaces.Expression, row int, column int) AssignmentHeap {
	return AssignmentHeap{Value:value, Temp:temp, Row:row, Col:column}
}

func (a AssignmentHeap) Execute() interface{} {

	pos  := a.Value.GetValue()
	temp := a.Temp.GetValue()
	
	return fmt.Sprintf("%v", temp) + " = heap[(int)" + fmt.Sprintf("%v",pos) + "];"
}
