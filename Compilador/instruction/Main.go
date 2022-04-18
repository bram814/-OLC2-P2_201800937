package instruction

import (
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	// "fmt"
)

type Main struct {
	Instrucciones *arrayList.List
	Row           int
	Column        int
}

func NewMain(instrucciones *arrayList.List, row int, column int) Main {
	instr := Main{instrucciones, row, column}
	return instr
}

func (p Main) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	var newTable interfaces.Environment
	newTable = interfaces.NewEnvironment(env)
	newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable)

	if p.Instrucciones != nil {
		// gen.AddFunction("void","main()")
		// gen.StachHeap()
		for _, s := range p.Instrucciones.ToArray() {


			if reflect.TypeOf(s).String() == "transferencia.Break" 	 { 
				excep := interfaces.NewException("Semantico","Sentencia Break fuera de Ciclo.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				return excep
			}
			if reflect.TypeOf(s).String() == "transferencia.Continue" { 
				excep := interfaces.NewException("Semantico","Sentencia Continue fuera de Ciclo.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				return excep
			}
			if reflect.TypeOf(s).String() == "transferencia.Return"   { 
				excep := interfaces.NewException("Semantico","Sentencia Return fuera de Ciclo.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo:excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
				return excep 
			}
			
			s.(interfaces.Instruction).Compilar(&newTable, tree, gen)

		}

		// gen.AddFunctionEnd()
	}

	return nil
}
