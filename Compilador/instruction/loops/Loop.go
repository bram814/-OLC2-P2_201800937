package loops

import (
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type Loop struct {
	Instrucciones *arrayList.List
	Row           int
	Column        int
}

func NewLoop(instruccion *arrayList.List, row int, column int) Loop {
	instr := Loop{instruccion, row, column}
	return instr
}

func (p Loop) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	Linicio := gen.NewLabel()
	Lfinal := gen.NewLabel()
	tree.AddDisplay(Linicio, Lfinal, "-1", false) // Display

	gen.AddComment("Loop")
	gen.AddLabel(Linicio)

	var newTable interfaces.Environment
	newTable = interfaces.NewEnvironment(env)
	newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable)

	for _, s := range p.Instrucciones.ToArray() {
		s.(interfaces.Instruction).Compilar(&newTable, tree, gen)

	}
	gen.AddGoto(Linicio)
	gen.AddLabel(Lfinal)
	tree.RestPosDisplay()

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}
