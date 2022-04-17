package ternario

import (
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
	"fmt"
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

func (p Loop) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {

	Linicio := gen.NewLabel()
	Lfinal := gen.NewLabel()
	temp := gen.NewTemp()
	var isType interfaces.TypeExpression
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
	gen.AddComment("Break temp")
	pos := fmt.Sprintf("%v", tree.PosDisplay-1)
	display := tree.GetDisplay(pos)
	if display.IsTemp {
		gen.AddExpression1(temp, display.Temp)
		isType = display.Type
	} else {

		excep := interfaces.NewException("Semantico", "Error en Loop, Break no trea una Expresion.", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return 
	}
	tree.RestPosDisplay()

	return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
}
