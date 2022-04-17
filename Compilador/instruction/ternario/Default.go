package ternario

import (
	"OLC2/Compilador/interfaces"
)

type Default struct {
	Instrucciones interfaces.Expression
}

func NewDefault(instrucciones interfaces.Expression) Default {
	instr := Default{instrucciones}
	return instr
}

func (p Default) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {

	result := p.Instrucciones.Compilar(env, tree, gen)

	return result
}
