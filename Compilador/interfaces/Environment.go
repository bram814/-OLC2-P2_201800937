package interfaces

import (
	"fmt"
	// "reflect"
)

type Environment struct {
	anterior 	*Environment
	variable 	map[string]Symbol
	Function 	map[string]Symbol
	// structs  map[string]interfaces.Symbol
	Posicion 	int

}

func NewEnvironment(anterior *Environment) Environment {

	env := Environment{anterior, make(map[string]Symbol), make(map[string]Symbol), 0}
	return env
}

func (env *Environment) IsAmbit() bool {

	var tmpEnv *Environment
	tmpEnv = env

	cont := 0

	for {

		if tmpEnv.anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.anterior
			cont++
		}
	}

	return 0 < cont
}

func (env *Environment) NewPos() {
	env.Posicion = env.Posicion + 1
}

func (env *Environment) UpdatePos(treePos int, envPos int, isBool bool , newTable *Environment) {

	
	newTable.Posicion = envPos
		
	

}

func (env *Environment) AddSymbol(id string, value Value, tipo TypeExpression, isMut bool, pos int, newTable *Environment) {

	if variable, ok := newTable.variable[id]; ok {
		fmt.Println("[ADD SYMBOL] La variable " + variable.Id + " ya existe")
		return
	}
	newTable.variable[id] = Symbol{Id: id, Type: tipo, Value: value, IsMut: isMut, Posicion: pos}
}

func (env *Environment) SearchSymbol(id string) Symbol {
	var tmpEnv *Environment
	tmpEnv = env

	for {

		if tmpEnv.anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.anterior
		}
	}
	if variable, ok := tmpEnv.variable[id]; ok {
		return variable
	}
	return Symbol{Id: "", Type: NULL, Value: Symbol{Id: "", Type: NULL, Value: 0}}
}

func (env *Environment) GetSymbol(id string) Symbol {

	var tmpEnv *Environment
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.variable[id]; ok {
			return variable
		}

		if tmpEnv.anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.anterior
		}
	}


	return Symbol{Id: "", Type: NULL, Value: Symbol{Id: "", Type: NULL, Value: 0}}
}

func (env *Environment) SetSymbol(id string, value Value, mut bool, pos int) Symbol {

	var tmpEnv *Environment
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.variable[id]; ok {
			tmpEnv.variable[id] = Symbol{Id: id, Type: variable.Type, Value: value, IsMut: mut, Posicion: pos}
			return variable
		}

		if tmpEnv.anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.anterior
		}
	}

	return Symbol{Id: "", Type: NULL, Value: Symbol{Id: "", Type: NULL, Value: 0}}
}

func (env *Environment) AddFunction(id string, value Symbol, tipo TypeExpression) {
	
	if variable, ok := env.Function[id]; ok {
		fmt.Println("[ADD SYMBOL] La variable " + variable.Id + " ya existe")
		return
	}
	env.Function[id] = Symbol{Id: id, Type: tipo, Value: value, IsMut: true, Posicion: 0}

}

func (env *Environment) GetFunction(id string) Symbol {

	var tmpEnv *Environment
	tmpEnv = env

	for {

		if variable, ok := tmpEnv.Function[id]; ok {
			return variable
		} 
		

		if tmpEnv.anterior == nil {
			break

		} else {
			tmpEnv = tmpEnv.anterior
		}
	}

	
	return Symbol{Id: "", Type: NULL, Value: Symbol{Id: "", Type: NULL, Value: 0}}

}
