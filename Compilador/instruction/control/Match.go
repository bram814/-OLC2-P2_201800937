package control

import (
	"OLC2/Compilador/ast"
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type Match struct {
	Condition    interfaces.Expression
	InstrCase    *arrayList.List
	InstrDefault *arrayList.List
	Row          int
	Column       int
}

func NewMatch(cond interfaces.Expression, instrCase *arrayList.List, instrDefault *arrayList.List, row int, column int) Match {
	instr := Match{cond, instrCase, instrDefault, row, column}
	return instr
}

func (p Match) Compilar(env *interfaces.Environment, tree *ast.Arbol, gen *ast.Generator) interface{} {

	var newTable interfaces.Environment
	newTable = interfaces.NewEnvironment(env)
	newTable.Posicion = tree.GetPos()

	var block2 string = "\n"
	var cond, newCond interfaces.Value

	cond = p.Condition.Compilar(env, tree, gen)

	Linicial := gen.NewLabel()
	newLabel := gen.NewLabel()
	Lfinal := gen.NewLabel()
	gen.AddComment("Match")
	gen.AddGoto(Linicial)

	if p.InstrCase != nil {

		gen.AddComment("Match [CASE]")

		for _, s := range p.InstrCase.ToArray() {

			if s.(Case).Condition != nil {
				newCond = s.(Case).Condition.Compilar(&newTable, tree, gen)
				EV := gen.NewLabel()
				block2 += "\t" + "if(" + cond.Value + " " + "==" + " " + newCond.Value + ") goto " + EV + ";\n"

			} else {
				for _, i := range s.(Case).ListaExpresion.ToArray() {
					newCond = i.(Case).Condition.Compilar(&newTable, tree, gen)
					EV := gen.NewLabel()
					block2 += "\t" + "if(" + cond.Value + " " + "==" + " " + newCond.Value + ") goto " + EV + ";\n"

					gen.AddLabel(EV)
					for _, j := range s.(Case).Instrucciones.ToArray() {
						j.(interfaces.Instruction).Compilar(&newTable, tree, gen)
					}
					gen.AddGoto(newLabel)

				}
			}
		}

		if p.InstrDefault != nil {
			gen.AddComment("Match [DEFAULT]")
			block2 += "\t" + "if(" + cond.Value + " " + "==" + " " + cond.Value + ") goto " + Lfinal + ";\n"
			gen.AddLabel(Lfinal)
			for _, s := range p.InstrDefault.ToArray() {
				s.(Default).Compilar(&newTable, tree, gen)
			}
			gen.AddGoto(newLabel)
		}

	} else {

		if p.InstrDefault != nil {
			gen.AddComment("Match [DEFAULT]")

			block2 += "\t" + "if(" + cond.Value + " " + "==" + " " + cond.Value + ") goto " + Lfinal + ";\n"
			gen.AddLabel(Lfinal)
			for _, s := range p.InstrDefault.ToArray() {
				s.(Default).Compilar(&newTable, tree, gen)
			}
			gen.AddGoto(newLabel)
		}
	}

	gen.AddLabel(Linicial)
	gen.AddConca(block2)
	gen.AddLabel(newLabel)

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}
