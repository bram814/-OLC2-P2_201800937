package Control

type Label struct {
	Id  	string
	Row 	int
	Col 	int
}

func NewLabel(id string, row int, col int) Label {
	return Label{id, row, col}
}

func (p Label) Execute() interface{} {
	return p.Id + ":"
}
