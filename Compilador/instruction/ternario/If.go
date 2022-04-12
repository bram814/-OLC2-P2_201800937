package ternario

import (

	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/ast"
	arrayList "github.com/colegno/arraylist"
)


type If struct {

	Condicion	interfaces.Expression
	InstrIf		interfaces.Expression
	InstrElse	interfaces.Expression
	InstrElseIf *arrayList.List
	Row 		int
	Column 		int
}


func NewIf(condicion interfaces.Expression, instrIf interfaces.Expression, instrElse interfaces.Expression, instrElseIf *arrayList.List, row int, column int) If {
	instr := If{condicion, instrIf, instrElse, instrElseIf, row, column}
	return instr
}


func (p If) Compilar(env interface{}, tree *ast.Arbol, gen *ast.Generator) interfaces.Value {


	gen.AddComment("Ternario - If")
	var cond, result interfaces.Value
	cond = p.Condicion.Compilar(env, tree, gen)


	if cond.Type == interfaces.EXCEPTION {
		return cond
	} 


	if cond.Type == interfaces.BOOLEAN {

		EV 		 := gen.NewLabel()
		EF 		 := gen.NewLabel()
		newLabel := gen.NewLabel()
		temp     := gen.NewTemp()

		gen.AddExpression(temp, "0","1","-")
		

		var isType interfaces.TypeExpression = interfaces.NULL

		gen.AddIf(cond.Value,"1","==",EV)
		gen.AddGoto(EF)

		gen.AddLabel(EV)
		result = p.InstrIf.Compilar(env, tree, gen)
		isType = result.Type
		gen.AddExpression(temp, result.Value,"","")
		
		gen.AddGoto(newLabel)


		gen.AddLabel(EF)

		if p.InstrElse != nil {
			gen.AddComment("Ternario - If (else)")
			result = p.InstrElse.Compilar(env, tree, gen)

			gen.AddExpression(temp, result.Value,"","")

			if isType != result.Type {
				excep := ast.NewException("Semantico","Tipos de Datos incorrectos en If (ternario).", p.Row, p.Column)
				tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}
		}

		if p.InstrElseIf != nil {
			for _, s := range p.InstrElseIf.ToArray() {
				gen.AddComment("Ternario - If (else if)")
				newInstr := s.(If).Compilar(env, tree, gen)
				if newInstr.Type != result.Type {
					excep := ast.NewException("Semantico","Tipos Diferentes en If (ternario).", p.Row, p.Column)
					tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
				}
				gen.AddComment("Ternario - Retorno de Temp")
				newLabel1 := gen.NewLabel()
				newLabel2 := gen.NewLabel()

				gen.AddIf(newInstr.Value,"-1","!=",newLabel1)
				gen.AddGoto(newLabel2)

				gen.AddLabel(newLabel1)
				gen.AddExpression(temp,newInstr.Value,"","") 
				gen.AddGoto(newLabel2)
				gen.AddLabel(newLabel2)
				
			}
		}
		gen.AddGoto(newLabel)


		gen.AddLabel(newLabel)
		return interfaces.Value{Value: temp, IsTemp: false, Type: result.Type, TrueLabel: "", FalseLabel: ""}

	} else {
		excep := ast.NewException("Semantico","Tipo de Dato no Booleano en If (ternario).", p.Row, p.Column)
		tree.AddException(ast.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}	