package interfaces

import (
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"time"
)

type Arbol struct {
	Code         *arrayList.List
	StackGlobal  int
	IsPrimitive  bool
	IsCocant     bool
	IsCompareStr bool
	_Exception   *arrayList.List
	Display      map[string]Display
	PosDisplay   int
}

type Exception struct {
	Tipo        string
	Descripcion string
	Row         int
	Column      int
	Time        string
}

type Display struct {
	Temp    string
	IsTemp  bool
	LInicio string
	LFinal  string
	Type 	TypeExpression
}

func NewArbol() *Arbol {
	tree := Arbol{
		Code:         arrayList.New(),
		StackGlobal:  0,
		IsPrimitive:  false,
		IsCocant:     false,
		IsCompareStr: false,
		_Exception:   arrayList.New(),
		Display:      make(map[string]Display),
		PosDisplay:   0}
	return &tree
}

func NewException(tipo string, descripcion string, row int, column int) *Exception {
	t := time.Now()
	hora := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	e := Exception{Tipo: tipo, Descripcion: descripcion, Row: row, Column: column, Time: hora}
	return &e
}

/* Add text to Code */
func (a *Arbol) AddCode(input string) {
	a.Code.Add(input)
}

/* Get text to Code */
func (a Arbol) GetCode() *arrayList.List {
	return a.Code
}

/* Add Exception */
func (a *Arbol) AddException(e Exception) {
	t := time.Now()
	hora := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d",
		t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute(), t.Second())
	e.Time = hora
	a._Exception.Add(e)
}

/* Get Exception */
func (a Arbol) GetException() *arrayList.List {
	return a._Exception
}

/* POS GLOBAL */
func (a Arbol) GetPos() int {
	return a.StackGlobal
}

func (a *Arbol) NewPos() {
	a.StackGlobal = a.StackGlobal + 1
}

/* *************************  DISPLAY ************************* */

func (a *Arbol) AddDisplay(labelInicio string, labelFinal string, temp string, isTemp bool) {

	pos := fmt.Sprintf("%v", a.PosDisplay)
	a.Display[pos] = Display{Temp: temp, IsTemp: isTemp, LInicio: labelInicio, LFinal: labelFinal, Type: NULL}
	a.PosDisplay = a.PosDisplay + 1
}

func (a *Arbol) SetDisplayTemp(pos string, temp string, isTemp bool, isType TypeExpression) {

	if display, ok := a.Display[pos]; ok {

		a.Display[pos] = Display{Temp: temp, IsTemp: isTemp, LInicio: display.LInicio, LFinal: display.LFinal, Type: isType}
	}

}

func (a *Arbol) GetDisplay(pos string) Display {

	if display, ok := a.Display[pos]; ok {

		return display
	}

	return Display{Temp: "-1", IsTemp: false, LInicio: "-1", LFinal: "-1"}

}

func (a *Arbol) RestPosDisplay() {
	a.PosDisplay = a.PosDisplay - 1
}
