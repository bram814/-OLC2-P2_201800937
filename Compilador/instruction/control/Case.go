package control

import (

	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/ast"
	arrayList "github.com/colegno/arraylist"
	
)

type Case struct {
	Condition 		interfaces.Expression
	ListaExpresion	*arrayList.List
	Instrucciones	*arrayList.List
}


func NewCase(cond interfaces.Expression, listaExpresion *arrayList.List, instrucciones *arrayList.List) Case {
	instr := Case{cond, listaExpresion, instrucciones}
	return instr
}	



func (p Case) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interfaces.Value {
	
	var result interfaces.Value
	if p.Condition != nil {
		result = p.Condition.Compilar(env, tree, gen)
		return result
	
	}
	return result
}