package control

import (

	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/environment"
	"OLC2/Compilador/ast"
	arrayList "github.com/colegno/arraylist"
	// "reflect"
	// "fmt"
)


type If struct {

	Condicion	interfaces.Expression
	InstrIf		*arrayList.List
	InstrElse	*arrayList.List
	InstrElseIf *arrayList.List
	Row 		int
	Column 		int

}


func NewIf(condicion interfaces.Expression, instrIf *arrayList.List, instrElse *arrayList.List, instrElseIf *arrayList.List, row int, column int) If {
	instr := If{condicion, instrIf, instrElse, instrElseIf, row, column}
	return instr
}


func (p If) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interface{} {


	var cond interfaces.Value
	cond = p.Condicion.Compilar(env, tree, gen)


	if cond.Type == interfaces.EXCEPTION {
		return cond
	} 

	gen.AddComment("Control - If")

	if cond.Type == interfaces.BOOLEAN {
		var newTable environment.Environment
		newTable = environment.NewEnvironment(env.(environment.Environment))

		EV 		 := gen.NewLabel()
		EF 		 := gen.NewLabel()
		newLabel := gen.NewLabel()

		gen.AddIf(cond.Value,"1","==",EV)
		gen.AddGoto(EF)

		gen.AddLabel(EV)
		for _, s := range p.InstrIf.ToArray() {
			s.(interfaces.Instruction).Compilar(newTable, tree, gen)
		}
		gen.AddGoto(newLabel)


		gen.AddLabel(EF)

		if p.InstrElse != nil {
			gen.AddComment("Control - If (else)")
			newTable = environment.NewEnvironment(env.(environment.Environment))
			for _, s := range p.InstrElse.ToArray() {

				s.(interfaces.Instruction).Compilar(newTable, tree, gen)
				
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
		excep := ast.NewException("Semantico","Error en If. Tipo de Dato no Booleano.", p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}	