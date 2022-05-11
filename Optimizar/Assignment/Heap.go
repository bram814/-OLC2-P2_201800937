package Assignment

import (
	"fmt"
	"OLC2/Optimizar/Interfaces"
)

type Heap struct {
	Value        	Interfaces.Expression
	Temp         	Interfaces.Expression
	InstrDeleted 	bool
	Row          	int
	Col          	int
}

func NewHeap(temp Interfaces.Expression, value Interfaces.Expression, row int, column int) Heap {
	return Heap{Value:value, Temp:temp, Row:row, Col:column}
}

func (a Heap) Execute() interface{} {

	value  := a.Value.GetValue()
	temp := a.Temp.GetValue()
	

	return "heap[(int)" + fmt.Sprintf("%v",temp) + "] = " + fmt.Sprintf("%v", value) + ";"
}
