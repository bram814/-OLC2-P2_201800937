package loops

import (
	"OLC2/Compilador/interfaces"
	"fmt"
	// "reflect"
	arrayList "github.com/colegno/arraylist"
)

type For struct {
	Id            string
	Type          interfaces.TypeExpression
	Left          interfaces.Expression
	Right         interfaces.Expression
	Instrucciones *arrayList.List
	Row           int
	Column        int
}

func NewFor(id string, tipo interfaces.TypeExpression, left interfaces.Expression, right interfaces.Expression, instrucciones *arrayList.List, row int, column int) For {
	instr := For{id, tipo, left, right, instrucciones, row, column}
	return instr
}

func (p For) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	var newTable interfaces.Environment
	newTable = interfaces.NewEnvironment(env)
	newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable)

	if p.Type == interfaces.INTEGER {

		gen.AddComment("For - Integer")
		left := p.Left.Compilar(&newTable, tree, gen)
		if left.Type == interfaces.EXCEPTION {
			return left
		}

		right := p.Right.Compilar(&newTable, tree, gen)
		if right.Type == interfaces.EXCEPTION {
			return right
		}

		symbol := newTable.GetSymbol(p.Id)

		if symbol.Type == interfaces.NULL {

			gen.AddComment("Declaracion")

			temp := gen.NewTemp()
			gen.AddExpression(temp, "P", fmt.Sprintf("%v", newTable.Posicion), "+")
			gen.AddStack(temp, left.Value)
			newTable.AddSymbol(p.Id, left, left.Type, true, newTable.Posicion, &newTable)
			tree.AddTableSymbol(*interfaces.NewTableSymbol(p.Id,"Variable - for","Local", p.Row, p.Column, "--", fmt.Sprintf("%v", newTable.Posicion)))
			newTable.NewPos()

		}

		Linicio := gen.NewLabel()
		gen.AddLabel(Linicio)
		symbol = newTable.GetSymbol(p.Id)
		gen.AddComment("Identificador")

		temp := gen.NewTemp()
		gen.AddExpressionStack(temp, fmt.Sprintf("%v", symbol.Posicion))

		EV := gen.NewLabel()
		Lfinal := gen.NewLabel()
		Lincre := gen.NewLabel()
		tree.AddDisplay(Lincre, Lfinal, "-1", false) // Display
		gen.AddComment("Relacional <")
		gen.AddIf(temp, right.Value, "<", EV)
		gen.AddGoto(Lfinal)

		gen.AddLabel(EV)
		
		// var newTable interfaces.Environment
		// newTable = interfaces.NewEnvironment(env)
		
		
		var SecondTable interfaces.Environment
		SecondTable = interfaces.NewEnvironment(&newTable)
		SecondTable.UpdatePos(tree.GetPos(), newTable.Posicion, newTable.Posicion != 0, &SecondTable)

		for _, s := range p.Instrucciones.ToArray() {
			s.(interfaces.Instruction).Compilar(&SecondTable, tree, gen)
		}

		pos := fmt.Sprintf("%v", tree.PosDisplay-1)
		display := tree.GetDisplay(pos)
		if display.IsTemp {
			excep := interfaces.NewException("Semantico", "Error en For, Sentencia de Control incorrecta Break.", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		}

		gen.AddComment("Incremento For")
		gen.AddLabel(Lincre)
		symbol = newTable.GetSymbol(p.Id)
		gen.AddComment("Identificador")
		temp = gen.NewTemp()
		gen.AddExpressionStack(temp, fmt.Sprintf("%v", symbol.Posicion))

		gen.AddComment("Asignacion")
		symbol.IsMut = true
		gen.AddExpression(temp, temp, "1", "+")
		left.Value = temp
		newTable.SetSymbol(p.Id, left, true, symbol.Posicion)
		gen.AddStack(fmt.Sprintf("%v", symbol.Posicion), temp)

		gen.AddGoto(Linicio)
		// gen.AddIf()

		gen.AddLabel(Lfinal)
		tree.RestPosDisplay()

	} else if p.Type == interfaces.STRING {

		gen.AddComment("For - String")

		left := p.Left.Compilar(&newTable, tree, gen)
		if left.Type == interfaces.EXCEPTION {
			return left
		}

		symbol := newTable.GetSymbol(p.Id)
		temp := gen.NewTemp()
		secondTemp := gen.NewTemp()

		if symbol.Type == interfaces.NULL {

			gen.AddComment("Identificador")
			gen.AddExpression(temp, "P", fmt.Sprintf("%v", newTable.Posicion), "+")
			gen.AddStack(temp, left.Value)
			left.Type = interfaces.CHAR
			newTable.AddSymbol(p.Id, left, interfaces.CHAR, true, newTable.Posicion, &newTable)
			tree.AddTableSymbol(*interfaces.NewTableSymbol(p.Id,"Variable - for","Local", p.Row, p.Column, "--", fmt.Sprintf("%v", newTable.Posicion)))
			newTable.NewPos()

			gen.AddExpressionStack(secondTemp, temp)
		}

		Linicio := gen.NewLabel()
		gen.AddLabel(Linicio)

		thirdTemp := gen.NewTemp()

		gen.AddExpressionHeap(thirdTemp, secondTemp)
		gen.AddStack(temp, secondTemp)
		gen.AddExpression(secondTemp, secondTemp, "1", "+")

		gen.AddComment("Add If")

		EV := gen.NewLabel()
		tree.AddDisplay(Linicio, EV, "-1", false)
		gen.AddIf(thirdTemp, "-1", "==", EV)

		for _, s := range p.Instrucciones.ToArray() {
			s.(interfaces.Instruction).Compilar(&newTable, tree, gen)

		}
		pos := fmt.Sprintf("%v", tree.PosDisplay-1)
		display := tree.GetDisplay(pos)
		
		if display.IsTemp {
			excep := interfaces.NewException("Semantico", "Error en For, Sentencia de Control incorrecta Break.", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		}


		gen.AddGoto(Linicio)
		gen.AddLabel(EV)
		tree.RestPosDisplay()

	}

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}
