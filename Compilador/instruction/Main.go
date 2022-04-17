package instruction

import (
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
	// "reflect"
	// "fmt"
)

type Main struct {
	Instrucciones *arrayList.List
	Row           int
	Column        int
}

func NewMain(instrucciones *arrayList.List, row int, column int) Main {
	instr := Main{instrucciones, row, column}
	return instr
}

func (p Main) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	var newTable interfaces.Environment
	newTable = interfaces.NewEnvironment(env)
	newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable)

	if p.Instrucciones != nil {
		// gen.AddFunction("void","main()")
		// gen.StachHeap()
		for _, s := range p.Instrucciones.ToArray() {
			s.(interfaces.Instruction).Compilar(&newTable, tree, gen)

		}

		// gen.AddFunctionEnd()
	}

	return nil
}
