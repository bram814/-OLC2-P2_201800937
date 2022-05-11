package Interfaces

type Expression interface {
	GetValue() interface{}
}

type Instruction interface {
	Execute() interface{}
}
