package expression

import (
	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/ast"
	"fmt"
)


type Primitivo struct {
	Value 	interface{}
	Type 	interfaces.TypeExpression
	Row 	int
	Column 	int
}

func NewPrimitivo(val interface{}, tipo interfaces.TypeExpression, row int, column int) Primitivo {
	exp := Primitivo{val, tipo, row, column}
	return exp
}

func (p Primitivo) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interfaces.Value {

	return interfaces.Value{
		Value:      fmt.Sprintf("%v", p.Value),
		IsTemp:     false,
		Type:       p.Type,
		TrueLabel:  "",
		FalseLabel: "",
	}
}

