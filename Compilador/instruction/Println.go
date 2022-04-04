package instruction

import (
	"fmt"
	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/ast"
)

type Println struct {
	Expression 	interfaces.Expression
	Row 		int
	Column 		int
}

func NewPrintln(val interfaces.Expression, row int, column int) Println {
	exp := Println{val, row, column}
	return exp
}

func (p Println) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interface{} {

	var result interfaces.Value

	result = p.Expression.Compilar(env, tree, gen)

	if result.Type == interfaces.BOOLEAN {
		newLabel := gen.NewLabel()
		gen.AddLabel(result.TrueLabel)
		gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", 1))
		gen.AddGoto(newLabel)
		gen.AddLabel(result.FalseLabel)
		gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", 0))
		gen.AddLabel(newLabel)

	} else {
		gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))
	}

	return result.Value
}
