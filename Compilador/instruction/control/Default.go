package control

import (

	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/environment"
	"OLC2/Compilador/ast"
	arrayList "github.com/colegno/arraylist"
	
)

type Default struct {
	Instrucciones	*arrayList.List
}


func NewDefault(instrucciones *arrayList.List) Default {
	instr := Default{instrucciones}
	return instr
}	



func (p Default) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interface{} {
	
	var newTable environment.Environment
	newTable = environment.NewEnvironment(env.(environment.Environment))
	
	for _, i := range p.Instrucciones.ToArray() {
		i.(interfaces.Instruction).Compilar(newTable, tree, gen)
	}


	return nil
}