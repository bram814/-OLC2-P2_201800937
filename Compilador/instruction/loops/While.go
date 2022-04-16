package loops

import (
	"OLC2/Compilador/ast"
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type While struct {
	Condition     interfaces.Expression
	Instrucciones *arrayList.List
	Row           int
	Column        int
}

func NewWhile(cond interfaces.Expression, instruccion *arrayList.List, row int, column int) While {
	instr := While{cond, instruccion, row, column}
	return instr
}

func (p While) Compilar(env *interfaces.Environment, tree *ast.Arbol, gen *ast.Generator) interface{} {

	var cond interfaces.Value
	Linicio := gen.NewLabel()
	EV := gen.NewLabel()
	EF := gen.NewLabel()
	tree.AddDisplay(Linicio, EF, "-1", false) // Display
	gen.AddComment("While")
	gen.AddLabel(Linicio)

	cond = p.Condition.Compilar(env, tree, gen)

	if cond.Type == interfaces.EXCEPTION {
		return cond
	}

	if cond.Type == interfaces.BOOLEAN {

		var newTable interfaces.Environment
		newTable = interfaces.NewEnvironment(env)
		newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable)

		gen.AddIf(cond.Value, "1", "==", EV)
		gen.AddGoto(EF)

		gen.AddLabel(EV)
		for _, s := range p.Instrucciones.ToArray() {
			s.(interfaces.Instruction).Compilar(&newTable, tree, gen)

		}
		gen.AddGoto(Linicio)
		gen.AddLabel(EF)
		tree.RestPosDisplay()

	} else {
		excep := ast.NewException("Semantico", "Tipo de Dato no Booleano en While.", p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}
