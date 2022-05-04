package arrays

import (
	// "fmt"
	"OLC2/Compilador/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type DimUnoDatos struct {
	Datos    *arrayList.List
}

func NewDimUnoDatos(Datos  *arrayList.List) DimUnoDatos {
	instr := DimUnoDatos{Datos}
	return instr
}

func (p DimUnoDatos) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {
	return p
}