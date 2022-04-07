package instruction

import (

	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/ast"
)

type ListExpre struct {
	Expresion interfaces.Expression

}

func NewListExpre(val interfaces.Expression) ListExpre {
	exp := ListExpre{val}
	return exp
}

func (p ListExpre) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interfaces.Value {


	var newInstr interfaces.Value
	newInstr = p.Expresion.Compilar(env, tree, gen)
	
	return newInstr

}