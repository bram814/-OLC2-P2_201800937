package Function

import (
	"fmt"
)

type Call struct {
	Id 		  	string
	Row       	int
	Col       	int
}

func NewCall(id string, row int, column int) Call {
	return Call{id, row, column}
}

func (p Call) Execute() interface{} {
	return fmt.Sprintf("%v", p.Id) + "();"
}
