package function

import (
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	"fmt"
)

type Function struct {

	Id 				string
	Parametro 	    *arrayList.List
	Instrucciones 	*arrayList.List
	Type			interfaces.TypeExpression
	Row				int
	Column			int

}

func NewFunction(Id string, Parametro *arrayList.List, Instrucciones *arrayList.List, tipo interfaces.TypeExpression, Row int, Column int) Function {
	instr := Function{Id, Parametro, Instrucciones, tipo, Row, Column}
	return instr
}

func (p Function) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	var newTable interfaces.Environment
	newTable = interfaces.NewEnvironment(env)
	newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable)
	env.NewPos()

	if p.Parametro != nil {

		for _,s := range p.Parametro.ToArray() {

			symbol := env.SearchSymbol(s.(ListExpre).Id)

			if symbol.Type != interfaces.NULL {

				excep := interfaces.NewException("Semantico", "Ya Existe ese Id "+s.(ListExpre).Id, p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			gen.AddComment("Identificador")
			temp0 := gen.NewTemp()
			temp1 := gen.NewTemp()
			gen.AddExpression(temp0, "P", fmt.Sprintf("%v", env.Posicion), "+")
			gen.AddExpressionStack(temp1, temp0)
			
			result := interfaces.Value{Value: temp1, IsTemp: true, Type: s.(ListExpre).Type, TrueLabel: "", FalseLabel: ""}
			
			env.AddSymbol(s.(ListExpre).Id, result, s.(ListExpre).Type, true, env.Posicion, &newTable)
			env.NewPos()


		}
	}
		

	if p.Instrucciones != nil {
		for _, s := range p.Instrucciones.ToArray() {
			
			if reflect.TypeOf(s).String() == "transferencia.Break" 	 { 
				excep := interfaces.NewException("Semantico","Sentencia Break fuera de Ciclo(Func).", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				return excep
			}
			if reflect.TypeOf(s).String() == "transferencia.Continue" { 
				excep := interfaces.NewException("Semantico","Sentencia Continue fuera de Ciclo (Func).", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				return excep
			}
			
			s.(interfaces.Instruction).Compilar(&newTable, tree, gen)
			
			
			

		}
	}

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}

}