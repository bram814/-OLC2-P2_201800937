package function

import (
	"fmt"
	"OLC2/Compilador/interfaces"
	"OLC2/Compilador/instruction"
	arrayList "github.com/colegno/arraylist"
)

type Llamada struct {

	Id 				string
	Parametro 	    *arrayList.List
	Row				int
	Column			int

}

func NewLlamada(Id string, Parametro *arrayList.List, Row int, Column int) Llamada {
	instr := Llamada{Id, Parametro, Row, Column}
	return instr
}

func (p Llamada) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	gen.AddComment("Llamada de Funcion")
	temp := gen.NewTemp()
	symbol := env.GetFunction(p.Id)
	function := symbol.Value.(interfaces.Symbol).Value.(interfaces.SymbolFunction)

	if  p.Id != function.Id {
		excep := interfaces.NewException("Semantico", "Error en Llamada, No Existe esa Funcion "+ p.Id, p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	if function.Parametro != nil && p.Parametro != nil {
			

		if function.Parametro.Len() == p.Parametro.Len() {


			for s := 0; s < p.Parametro.Len(); s++ {
				instrCall := p.Parametro.GetValue(s).(instruction.ListExpre).Compilar(env, tree, gen)
				instrFunc := function.Parametro.GetValue(s).(ListExpre)

				if instrCall.Type == instrFunc.Type {
					if s == 0 {
						gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion+1), "+")
						gen.AddStack(temp, instrCall.Value)

					} else {
						gen.AddExpression(temp, temp, "1", "+")
						gen.AddStack(temp, instrCall.Value)

					}
					
				}
				

			}

			gen.AddExpression("P","P", fmt.Sprintf("%v", env.Posicion), "+")
			gen.PrintFunction(p.Id)
			gen.AddExpressionStack(temp,"P")
			gen.AddExpression("P","P", fmt.Sprintf("%v", env.Posicion), "-")


		} else {
			excep := interfaces.NewException("Semantico", "Error en Llamada, cantidad de Parametros incorrecto.", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}


	} else {

		if (function.Parametro != nil && p.Parametro == nil) || (function.Parametro == nil && p.Parametro != nil) {
			excep := interfaces.NewException("Semantico", "Error en Llamada, parametros incorrectos.", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		
		} else {
			gen.AddExpression("P","P", fmt.Sprintf("%v", env.Posicion), "+")
			gen.PrintFunction(p.Id)
			gen.AddExpressionStack(temp,"P")
			gen.AddExpression("P","P", fmt.Sprintf("%v", env.Posicion), "-")

		}


	}

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}
