package Primitive

import (
	"fmt"
	"OLC2/Optimizar/Interfaces"
)

type Operation struct {
	OpLeft   Interfaces.Expression
	Operator string
	OpRight  Interfaces.Expression
	Row      int
	Col      int
}

func NewOperation(left Interfaces.Expression, operator string, right Interfaces.Expression, row int, column int) Operation {
	return Operation{left, operator, right, row, column}
}

func (p Operation) GetValue() interface{} {
	return fmt.Sprintf("%v", p.OpLeft.GetValue()) + " " + p.Operator + " " + fmt.Sprintf("%v", p.OpRight.GetValue())
}
