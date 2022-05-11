package Assignment

import (
	"fmt"
	"OLC2/Optimizar/Interfaces"
)

type Assign struct {
	Value        	Interfaces.Expression
	Temp         	Interfaces.Expression
	InstrDeleted 	bool
	Row          	int
	Col          	int
}

func NewAssign(temp Interfaces.Expression, value Interfaces.Expression, row int, column int) Assign {
	return Assign{Value:value, Temp:temp, Row:row, Col:column}
}

func (a Assign) Execute() interface{} {

	value := a.Value.GetValue()
	temp := a.Temp.GetValue()

	if temp.(string) == "P" || temp.(string) == "H" {
		return temp.(string) + " = " + fmt.Sprintf("%v",value) + ";"
	}

	typeAssign := make(map[string]interface{})
	typeAssign["temp"] = temp
	typeAssign["val"] = value
	typeAssign["rep"] = false
	typeAssign["row"] = fmt.Sprintf("%v", a.Row)

	return typeAssign
}
