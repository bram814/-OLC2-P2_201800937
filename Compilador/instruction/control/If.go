package control

import (
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
	// "reflect"
	// "fmt"
)

type If struct {
	Condicion   interfaces.Expression
	InstrIf     *arrayList.List
	InstrElse   *arrayList.List
	InstrElseIf *arrayList.List
	Row         int
	Column      int
}

func NewIf(condicion interfaces.Expression, instrIf *arrayList.List, instrElse *arrayList.List, instrElseIf *arrayList.List, row int, column int) If {
	instr := If{condicion, instrIf, instrElse, instrElseIf, row, column}
	return instr
}

func (p If) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	var cond interfaces.Value
	cond = p.Condicion.Compilar(env, tree, gen)

	if cond.Type == interfaces.EXCEPTION {
		return cond
	}

	gen.AddComment("Control - If")

	if cond.Type == interfaces.BOOLEAN {

		var newTable interfaces.Environment
		newTable = interfaces.NewEnvironment(env)
		newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable)

		EV := gen.NewLabel()
		EF := gen.NewLabel()
		newLabel := gen.NewLabel()

		gen.AddIf(cond.Value, "1", "==", EV)
		gen.AddGoto(EF)

		gen.AddLabel(EV)
		for _, s := range p.InstrIf.ToArray() {
			s.(interfaces.Instruction).Compilar(&newTable, tree, gen)
		}
		gen.AddGoto(newLabel)

		gen.AddLabel(EF)

		if p.InstrElse != nil {
			gen.AddComment("Control - If (else)")
			newTable = interfaces.NewEnvironment(env)
			newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable)

			for _, s := range p.InstrElse.ToArray() {

				s.(interfaces.Instruction).Compilar(&newTable, tree, gen)

			}

		}

		if p.InstrElseIf != nil {
			gen.AddComment("Control - If (else if)")
			for _, s := range p.InstrElseIf.ToArray() {
				s.(interfaces.Instruction).Compilar(env, tree, gen)

			}
		}
		gen.AddGoto(newLabel)

		gen.AddLabel(newLabel)

	} else {
		excep := interfaces.NewException("Semantico", "Error en If. Tipo de Dato no Booleano.", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}
