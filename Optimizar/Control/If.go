package Control

import (
	"OLC2/Optimizar/Interfaces"
)

type If struct {
	Condition Interfaces.Expression
	label     string
	Row       int
	Col       int
}

func NewIf(condition Interfaces.Expression, label string, row int, col int) If {
	return If{condition, label, row, col}
}

func (i If) Execute() interface{} {
	return "if (" + i.Condition.GetValue().(string) + ") goto " + i.label + ";"
}
