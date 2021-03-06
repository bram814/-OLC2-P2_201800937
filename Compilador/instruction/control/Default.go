package control

import (
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type Default struct {
	Instrucciones *arrayList.List
}

func NewDefault(instrucciones *arrayList.List) Default {
	instr := Default{instrucciones}
	return instr
}

func (p Default) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	var newTable interfaces.Environment
	newTable = interfaces.NewEnvironment(env)
	newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable)

	for _, i := range p.Instrucciones.ToArray() {
		i.(interfaces.Instruction).Compilar(&newTable, tree, gen)
	}

	return nil
}
