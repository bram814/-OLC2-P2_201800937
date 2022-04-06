package instruction


import (
	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/environment"
	"OLC2/Compilador/ast"
	arrayList "github.com/colegno/arraylist"
	// "reflect"
	// "fmt"

)


type Main struct {
	Instrucciones   *arrayList.List
	Row 			int
	Column			int
}


func NewMain(instrucciones *arrayList.List, row int, column int) Main {
	instr := Main{instrucciones, row, column}
	return instr
}


func (p Main) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interface{} {

	var newTable environment.Environment
	newTable = environment.NewEnvironment(env.(environment.Environment))

	if p.Instrucciones != nil {
		gen.AddFunction("void","main()")
		gen.StachHeap()
		for _, s := range p.Instrucciones.ToArray() {
			s.(interfaces.Instruction).Compilar(newTable, tree, gen)

		}

		gen.AddFunctionEnd()
	}
	
	return nil
}