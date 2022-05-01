package ternario

import (
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
	// "reflect"
	// "fmt"
)

type Match struct {
	Condition    interfaces.Expression
	InstrCase    *arrayList.List
	InstrDefault interfaces.Expression
	Row          int
	Column       int
}

func NewMatch(cond interfaces.Expression, instrCase *arrayList.List, instrDefault interfaces.Expression, row int, column int) Match {
	instr := Match{cond, instrCase, instrDefault, row, column}
	return instr
}

func (p Match) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {

	var block2 string = "\n"
	var cond, newCond interfaces.Value
	var isType = interfaces.NULL

	cond = p.Condition.Compilar(env, tree, gen)

	Linicial := gen.NewLabel()
	newLabel := gen.NewLabel()
	Lfinal := gen.NewLabel()
	gen.AddComment("Ternario - Match")
	temp := gen.NewTemp()
	gen.AddGoto(Linicial)

	if p.InstrCase != nil {

		gen.AddComment("Match [CASE]")

		for _, s := range p.InstrCase.ToArray() {

			if s.(Case).Condition != nil {

				newCond = s.(Case).Condition.Compilar(env, tree, gen)
				EV := gen.NewLabel()
				block2 += "\t" + "if(" + cond.Value + " " + "==" + " " + newCond.Value + ") goto " + EV + ";\n"

			} else {
				for _, i := range s.(Case).ListaExpresion.ToArray() {
					newCond = i.(Case).Condition.Compilar(env, tree, gen)

					EV := gen.NewLabel()
					block2 += "\t" + "if(" + cond.Value + " " + "==" + " " + newCond.Value + ") goto " + EV + ";\n"

					gen.AddLabel(EV)

					newInstr := s.(Case).Instrucciones.Compilar(env, tree, gen)

					if isType != interfaces.NULL && isType != newInstr.Type {
						excep := interfaces.NewException("Semantico", "Tipos de Datos incorrectos en Match (ternario).", p.Row, p.Column)
						tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
						return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
					} else {
						isType = newInstr.Type
					}

					gen.AddExpression1(temp, newInstr.Value)
					gen.AddGoto(newLabel)

				}
			}
		}

		if p.InstrDefault != nil {
			gen.AddComment("Match [DEFAULT]")
			block2 += "\t" + "if(" + cond.Value + " " + "==" + " " + cond.Value + ") goto " + Lfinal + ";\n"
			gen.AddLabel(Lfinal)
			newInstr := p.InstrDefault.(Default).Compilar(env, tree, gen)

			if isType != interfaces.NULL && isType != newInstr.Type {
				excep := interfaces.NewException("Semantico", "Tipos de Datos incorrectos en Match (ternario).", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			} else {
				isType = newInstr.Type
			}

			gen.AddExpression1(temp, newInstr.Value)
			gen.AddGoto(newLabel)

		}

	} else {

		if p.InstrDefault != nil {
			gen.AddComment("Match [DEFAULT]")

			block2 += "\t" + "if(" + cond.Value + " " + "==" + " " + cond.Value + ") goto " + Lfinal + ";\n"
			gen.AddLabel(Lfinal)
			newInstr := p.InstrDefault.(Default).Compilar(env, tree, gen)

			if isType != interfaces.NULL && isType != newInstr.Type {
				excep := interfaces.NewException("Semantico", "Tipos de Datos incorrectos en Match (ternario).", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			} else {
				isType = newInstr.Type
			}

			gen.AddExpression1(temp, newInstr.Value)
			gen.AddGoto(newLabel)
		}
	}

	gen.AddLabel(Linicial)
	gen.AddConca(block2)
	gen.AddLabel(newLabel)

	return interfaces.Value{Value: temp, IsTemp: false, Type: isType, TrueLabel: "", FalseLabel: ""}
}
