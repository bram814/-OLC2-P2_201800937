package Primitive

type Temp struct {
	Value interface{}
	Temp  string
	Row   int
	Col   int
}

func NewTemp(temp string, row int, column int) Temp {
	return Temp{Temp:temp, Row:row, Col:column }
}

func (t Temp) GetValue() interface{} {
	return t.Temp
}
