package Primitive

type Number struct {
	Value 	interface{}
	Row   	int
	Col   	int
}

func NewNumber(value interface{}, row int, col int) Number {
	return Number{value,row,col}
}

func (p Number) GetValue() interface{} {
	return p.Value
}
