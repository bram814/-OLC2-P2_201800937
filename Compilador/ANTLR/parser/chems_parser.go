// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2/Compilador/interfaces"
import "OLC2/Compilador/expression"
import "OLC2/Compilador/instruction"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 63, 143,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 25,
	10, 3, 13, 3, 14, 3, 26, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 34, 10, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 43, 10, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 61, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 88, 10, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 110, 10, 9, 12, 9, 14, 9, 113, 11,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 141, 10, 10, 3, 10, 2, 3,
	16, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 7, 4, 2, 50, 50, 60, 60, 3,
	2, 46, 48, 3, 2, 49, 50, 4, 2, 38, 39, 42, 45, 3, 2, 57, 58, 2, 151, 2,
	20, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 33, 3, 2, 2, 2, 8, 42, 3, 2, 2, 2,
	10, 60, 3, 2, 2, 2, 12, 62, 3, 2, 2, 2, 14, 71, 3, 2, 2, 2, 16, 87, 3,
	2, 2, 2, 18, 140, 3, 2, 2, 2, 20, 21, 5, 4, 3, 2, 21, 22, 8, 2, 1, 2, 22,
	3, 3, 2, 2, 2, 23, 25, 5, 8, 5, 2, 24, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2,
	2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 29,
	8, 3, 1, 2, 29, 5, 3, 2, 2, 2, 30, 31, 7, 34, 2, 2, 31, 34, 8, 4, 1, 2,
	32, 34, 8, 4, 1, 2, 33, 30, 3, 2, 2, 2, 33, 32, 3, 2, 2, 2, 34, 7, 3, 2,
	2, 2, 35, 36, 5, 10, 6, 2, 36, 37, 5, 6, 4, 2, 37, 38, 8, 5, 1, 2, 38,
	43, 3, 2, 2, 2, 39, 40, 5, 12, 7, 2, 40, 41, 8, 5, 1, 2, 41, 43, 3, 2,
	2, 2, 42, 35, 3, 2, 2, 2, 42, 39, 3, 2, 2, 2, 43, 9, 3, 2, 2, 2, 44, 45,
	7, 4, 2, 2, 45, 46, 7, 51, 2, 2, 46, 47, 5, 14, 8, 2, 47, 48, 7, 52, 2,
	2, 48, 49, 7, 34, 2, 2, 49, 50, 8, 6, 1, 2, 50, 61, 3, 2, 2, 2, 51, 52,
	7, 4, 2, 2, 52, 53, 7, 51, 2, 2, 53, 54, 7, 29, 2, 2, 54, 55, 7, 35, 2,
	2, 55, 56, 5, 14, 8, 2, 56, 57, 7, 52, 2, 2, 57, 58, 7, 34, 2, 2, 58, 59,
	8, 6, 1, 2, 59, 61, 3, 2, 2, 2, 60, 44, 3, 2, 2, 2, 60, 51, 3, 2, 2, 2,
	61, 11, 3, 2, 2, 2, 62, 63, 7, 19, 2, 2, 63, 64, 7, 18, 2, 2, 64, 65, 7,
	51, 2, 2, 65, 66, 7, 52, 2, 2, 66, 67, 7, 53, 2, 2, 67, 68, 5, 4, 3, 2,
	68, 69, 7, 54, 2, 2, 69, 70, 8, 7, 1, 2, 70, 13, 3, 2, 2, 2, 71, 72, 5,
	16, 9, 2, 72, 73, 8, 8, 1, 2, 73, 15, 3, 2, 2, 2, 74, 75, 8, 9, 1, 2, 75,
	76, 9, 2, 2, 2, 76, 77, 5, 16, 9, 5, 77, 78, 8, 9, 1, 2, 78, 88, 3, 2,
	2, 2, 79, 80, 5, 18, 10, 2, 80, 81, 8, 9, 1, 2, 81, 88, 3, 2, 2, 2, 82,
	83, 7, 51, 2, 2, 83, 84, 5, 14, 8, 2, 84, 85, 7, 52, 2, 2, 85, 86, 8, 9,
	1, 2, 86, 88, 3, 2, 2, 2, 87, 74, 3, 2, 2, 2, 87, 79, 3, 2, 2, 2, 87, 82,
	3, 2, 2, 2, 88, 111, 3, 2, 2, 2, 89, 90, 12, 9, 2, 2, 90, 91, 9, 3, 2,
	2, 91, 92, 5, 16, 9, 10, 92, 93, 8, 9, 1, 2, 93, 110, 3, 2, 2, 2, 94, 95,
	12, 8, 2, 2, 95, 96, 9, 4, 2, 2, 96, 97, 5, 16, 9, 9, 97, 98, 8, 9, 1,
	2, 98, 110, 3, 2, 2, 2, 99, 100, 12, 7, 2, 2, 100, 101, 9, 5, 2, 2, 101,
	102, 5, 16, 9, 8, 102, 103, 8, 9, 1, 2, 103, 110, 3, 2, 2, 2, 104, 105,
	12, 6, 2, 2, 105, 106, 9, 6, 2, 2, 106, 107, 5, 16, 9, 7, 107, 108, 8,
	9, 1, 2, 108, 110, 3, 2, 2, 2, 109, 89, 3, 2, 2, 2, 109, 94, 3, 2, 2, 2,
	109, 99, 3, 2, 2, 2, 109, 104, 3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109,
	3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 17, 3, 2, 2, 2, 113, 111, 3, 2,
	2, 2, 114, 115, 7, 26, 2, 2, 115, 141, 8, 10, 1, 2, 116, 117, 7, 27, 2,
	2, 117, 141, 8, 10, 1, 2, 118, 119, 7, 29, 2, 2, 119, 141, 8, 10, 1, 2,
	120, 121, 7, 30, 2, 2, 121, 141, 8, 10, 1, 2, 122, 123, 7, 28, 2, 2, 123,
	141, 8, 10, 1, 2, 124, 125, 7, 26, 2, 2, 125, 126, 7, 20, 2, 2, 126, 127,
	7, 21, 2, 2, 127, 141, 8, 10, 1, 2, 128, 129, 7, 26, 2, 2, 129, 130, 7,
	20, 2, 2, 130, 131, 7, 22, 2, 2, 131, 141, 8, 10, 1, 2, 132, 133, 7, 27,
	2, 2, 133, 134, 7, 20, 2, 2, 134, 135, 7, 21, 2, 2, 135, 141, 8, 10, 1,
	2, 136, 137, 7, 27, 2, 2, 137, 138, 7, 20, 2, 2, 138, 139, 7, 22, 2, 2,
	139, 141, 8, 10, 1, 2, 140, 114, 3, 2, 2, 2, 140, 116, 3, 2, 2, 2, 140,
	118, 3, 2, 2, 2, 140, 120, 3, 2, 2, 2, 140, 122, 3, 2, 2, 2, 140, 124,
	3, 2, 2, 2, 140, 128, 3, 2, 2, 2, 140, 132, 3, 2, 2, 2, 140, 136, 3, 2,
	2, 2, 141, 19, 3, 2, 2, 2, 10, 26, 33, 42, 60, 87, 109, 111, 140,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'print!'", "'println!'", "'number'", "'let'", "'mut'", "'if'", "'else'",
	"'match'", "'while'", "'break'", "'continue'", "'return'", "'loop'", "'for'",
	"'in'", "'main'", "'fn'", "'as'", "'i64'", "'f64'", "'String'", "'bool'",
	"'&str'", "", "", "", "", "", "", "'..'", "'.'", "';'", "','", "':'", "'='",
	"'=='", "'>='", "'=>'", "'->'", "'<='", "'!='", "'>'", "'<'", "'*'", "'/'",
	"'%'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'&&'",
	"'||'", "'|'", "'!'",
}
var symbolicNames = []string{
	"", "R_PRINT", "R_PRINTLN", "P_NUMBER", "R_LET", "R_MUT", "R_IF", "R_ELSE",
	"R_MATCH", "R_WHILE", "R_BREAK", "R_CONTINUE", "R_RETURN", "R_LOOP", "R_FOR",
	"R_IN", "R_MAIN", "R_FUNCTION", "R_AS", "R_INT", "R_FLOAT", "R_STRING",
	"R_BOOL", "R_STR", "NUMBER", "DOUBLE", "CHAR", "STRING", "BOOLEAN", "ID",
	"TK_DOBLEPUNTO", "TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA", "TK_DOSPUNTOS",
	"TK_IGUAL", "TK_IGUALIGUAL", "TK_MAYORIGUAL", "TK_IGUALMAYOR", "TK_MENOSMAYOR",
	"TK_MENORIGUAL", "TK_DIFIGUAL", "TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV",
	"TK_MODULO", "TK_MAS", "TK_MENOS", "TK_PARA", "TK_PARC", "TK_LLAVEA", "TK_LLAVEC",
	"TK_CORA", "TK_CORC", "TK_AND", "TK_OR", "TK_BARRA", "TK_NOT", "WHITESPACE",
	"TK_MULTI", "TK_LINE",
}

var ruleNames = []string{
	"start", "instrucciones", "end_instr", "instruccion", "instr_println",
	"instr_main", "expressions", "instr_expre", "primitivo",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Chems struct {
	*antlr.BaseParser
}

func NewChems(input antlr.TokenStream) *Chems {
	this := new(Chems)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Chems.g4"

	return this
}

// Chems tokens.
const (
	ChemsEOF           = antlr.TokenEOF
	ChemsR_PRINT       = 1
	ChemsR_PRINTLN     = 2
	ChemsP_NUMBER      = 3
	ChemsR_LET         = 4
	ChemsR_MUT         = 5
	ChemsR_IF          = 6
	ChemsR_ELSE        = 7
	ChemsR_MATCH       = 8
	ChemsR_WHILE       = 9
	ChemsR_BREAK       = 10
	ChemsR_CONTINUE    = 11
	ChemsR_RETURN      = 12
	ChemsR_LOOP        = 13
	ChemsR_FOR         = 14
	ChemsR_IN          = 15
	ChemsR_MAIN        = 16
	ChemsR_FUNCTION    = 17
	ChemsR_AS          = 18
	ChemsR_INT         = 19
	ChemsR_FLOAT       = 20
	ChemsR_STRING      = 21
	ChemsR_BOOL        = 22
	ChemsR_STR         = 23
	ChemsNUMBER        = 24
	ChemsDOUBLE        = 25
	ChemsCHAR          = 26
	ChemsSTRING        = 27
	ChemsBOOLEAN       = 28
	ChemsID            = 29
	ChemsTK_DOBLEPUNTO = 30
	ChemsTK_PUNTO      = 31
	ChemsTK_PUNTOCOMA  = 32
	ChemsTK_COMA       = 33
	ChemsTK_DOSPUNTOS  = 34
	ChemsTK_IGUAL      = 35
	ChemsTK_IGUALIGUAL = 36
	ChemsTK_MAYORIGUAL = 37
	ChemsTK_IGUALMAYOR = 38
	ChemsTK_MENOSMAYOR = 39
	ChemsTK_MENORIGUAL = 40
	ChemsTK_DIFIGUAL   = 41
	ChemsTK_MAYOR      = 42
	ChemsTK_MENOR      = 43
	ChemsTK_MULT       = 44
	ChemsTK_DIV        = 45
	ChemsTK_MODULO     = 46
	ChemsTK_MAS        = 47
	ChemsTK_MENOS      = 48
	ChemsTK_PARA       = 49
	ChemsTK_PARC       = 50
	ChemsTK_LLAVEA     = 51
	ChemsTK_LLAVEC     = 52
	ChemsTK_CORA       = 53
	ChemsTK_CORC       = 54
	ChemsTK_AND        = 55
	ChemsTK_OR         = 56
	ChemsTK_BARRA      = 57
	ChemsTK_NOT        = 58
	ChemsWHITESPACE    = 59
	ChemsTK_MULTI      = 60
	ChemsTK_LINE       = 61
)

// Chems rules.
const (
	ChemsRULE_start         = 0
	ChemsRULE_instrucciones = 1
	ChemsRULE_end_instr     = 2
	ChemsRULE_instruccion   = 3
	ChemsRULE_instr_println = 4
	ChemsRULE_instr_main    = 5
	ChemsRULE_expressions   = 6
	ChemsRULE_instr_expre   = 7
	ChemsRULE_primitivo     = 8
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_instrucciones IInstruccionesContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *StartContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *StartContext) GetLista() *arrayList.List { return s.lista }

func (s *StartContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Chems) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChemsRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)

		var _x = p.Instrucciones()

		localctx.(*StartContext)._instrucciones = _x
	}
	localctx.(*StartContext).lista = localctx.(*StartContext).Get_instrucciones().GetL()

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetList returns the list rule context list.
	GetList() []IInstruccionContext

	// SetList sets the list rule context list.
	SetList([]IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruccion IInstruccionContext
	list         []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetList() []IInstruccionContext { return s.list }

func (s *InstruccionesContext) SetList(v []IInstruccionContext) { s.list = v }

func (s *InstruccionesContext) GetL() *arrayList.List { return s.l }

func (s *InstruccionesContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Chems) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChemsRULE_instrucciones)

	localctx.(*InstruccionesContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ChemsR_PRINTLN || _la == ChemsR_FUNCTION {
		{
			p.SetState(21)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).list = append(localctx.(*InstruccionesContext).list, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetList()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IEnd_instrContext is an interface to support dynamic dispatch.
type IEnd_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v attribute.
	GetV() int

	// SetV sets the v attribute.
	SetV(int)

	// IsEnd_instrContext differentiates from other interfaces.
	IsEnd_instrContext()
}

type End_instrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      int
}

func NewEmptyEnd_instrContext() *End_instrContext {
	var p = new(End_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_end_instr
	return p
}

func (*End_instrContext) IsEnd_instrContext() {}

func NewEnd_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *End_instrContext {
	var p = new(End_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_end_instr

	return p
}

func (s *End_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *End_instrContext) GetV() int { return s.v }

func (s *End_instrContext) SetV(v int) { s.v = v }

func (s *End_instrContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *End_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *End_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *End_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterEnd_instr(s)
	}
}

func (s *End_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitEnd_instr(s)
	}
}

func (p *Chems) End_instr() (localctx IEnd_instrContext) {
	localctx = NewEnd_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ChemsRULE_end_instr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(31)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_PUNTOCOMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*End_instrContext).v = 1

	case ChemsEOF, ChemsR_PRINTLN, ChemsR_FUNCTION, ChemsTK_LLAVEC:
		p.EnterOuterAlt(localctx, 2)
		localctx.(*End_instrContext).v = 0

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_println returns the _instr_println rule contexts.
	Get_instr_println() IInstr_printlnContext

	// Get_instr_main returns the _instr_main rule contexts.
	Get_instr_main() IInstr_mainContext

	// Set_instr_println sets the _instr_println rule contexts.
	Set_instr_println(IInstr_printlnContext)

	// Set_instr_main sets the _instr_main rule contexts.
	Set_instr_main(IInstr_mainContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_instr_println IInstr_printlnContext
	_instr_main    IInstr_mainContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_instr_println() IInstr_printlnContext { return s._instr_println }

func (s *InstruccionContext) Get_instr_main() IInstr_mainContext { return s._instr_main }

func (s *InstruccionContext) Set_instr_println(v IInstr_printlnContext) { s._instr_println = v }

func (s *InstruccionContext) Set_instr_main(v IInstr_mainContext) { s._instr_main = v }

func (s *InstruccionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *InstruccionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *InstruccionContext) Instr_println() IInstr_printlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_printlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_printlnContext)
}

func (s *InstruccionContext) End_instr() IEnd_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnd_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnd_instrContext)
}

func (s *InstruccionContext) Instr_main() IInstr_mainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_mainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_mainContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Chems) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChemsRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsR_PRINTLN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)

			var _x = p.Instr_println()

			localctx.(*InstruccionContext)._instr_println = _x
		}
		{
			p.SetState(34)
			p.End_instr()
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_println().GetInstr()

	case ChemsR_FUNCTION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)

			var _x = p.Instr_main()

			localctx.(*InstruccionContext)._instr_main = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instr_main().GetInstr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstr_printlnContext is an interface to support dynamic dispatch.
type IInstr_printlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_PRINTLN returns the _R_PRINTLN token.
	Get_R_PRINTLN() antlr.Token

	// Set_R_PRINTLN sets the _R_PRINTLN token.
	Set_R_PRINTLN(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_printlnContext differentiates from other interfaces.
	IsInstr_printlnContext()
}

type Instr_printlnContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruction
	_R_PRINTLN   antlr.Token
	_expressions IExpressionsContext
}

func NewEmptyInstr_printlnContext() *Instr_printlnContext {
	var p = new(Instr_printlnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_println
	return p
}

func (*Instr_printlnContext) IsInstr_printlnContext() {}

func NewInstr_printlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_printlnContext {
	var p = new(Instr_printlnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_println

	return p
}

func (s *Instr_printlnContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_printlnContext) Get_R_PRINTLN() antlr.Token { return s._R_PRINTLN }

func (s *Instr_printlnContext) Set_R_PRINTLN(v antlr.Token) { s._R_PRINTLN = v }

func (s *Instr_printlnContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_printlnContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_printlnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_printlnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_printlnContext) R_PRINTLN() antlr.TerminalNode {
	return s.GetToken(ChemsR_PRINTLN, 0)
}

func (s *Instr_printlnContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_printlnContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_printlnContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_printlnContext) TK_PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PUNTOCOMA, 0)
}

func (s *Instr_printlnContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *Instr_printlnContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_COMA, 0)
}

func (s *Instr_printlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_printlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_printlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_println(s)
	}
}

func (s *Instr_printlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_println(s)
	}
}

func (p *Chems) Instr_println() (localctx IInstr_printlnContext) {
	localctx = NewInstr_printlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChemsRULE_instr_println)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)

			var _m = p.Match(ChemsR_PRINTLN)

			localctx.(*Instr_printlnContext)._R_PRINTLN = _m
		}
		{
			p.SetState(43)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(44)

			var _x = p.Expressions()

			localctx.(*Instr_printlnContext)._expressions = _x
		}
		{
			p.SetState(45)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(46)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_printlnContext).instr = instruction.NewPrintln(localctx.(*Instr_printlnContext).Get_expressions().GetP(), (func() int {
			if localctx.(*Instr_printlnContext).Get_R_PRINTLN() == nil {
				return 0
			} else {
				return localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetLine()
			}
		}()), localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)

			var _m = p.Match(ChemsR_PRINTLN)

			localctx.(*Instr_printlnContext)._R_PRINTLN = _m
		}
		{
			p.SetState(50)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(51)
			p.Match(ChemsSTRING)
		}
		{
			p.SetState(52)
			p.Match(ChemsTK_COMA)
		}
		{
			p.SetState(53)

			var _x = p.Expressions()

			localctx.(*Instr_printlnContext)._expressions = _x
		}
		{
			p.SetState(54)
			p.Match(ChemsTK_PARC)
		}
		{
			p.SetState(55)
			p.Match(ChemsTK_PUNTOCOMA)
		}
		localctx.(*Instr_printlnContext).instr = instruction.NewPrintln(localctx.(*Instr_printlnContext).Get_expressions().GetP(), (func() int {
			if localctx.(*Instr_printlnContext).Get_R_PRINTLN() == nil {
				return 0
			} else {
				return localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetLine()
			}
		}()), localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn())

	}

	return localctx
}

// IInstr_mainContext is an interface to support dynamic dispatch.
type IInstr_mainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_R_MAIN returns the _R_MAIN token.
	Get_R_MAIN() antlr.Token

	// Set_R_MAIN sets the _R_MAIN token.
	Set_R_MAIN(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstr_mainContext differentiates from other interfaces.
	IsInstr_mainContext()
}

type Instr_mainContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_R_MAIN        antlr.Token
	_instrucciones IInstruccionesContext
}

func NewEmptyInstr_mainContext() *Instr_mainContext {
	var p = new(Instr_mainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_main
	return p
}

func (*Instr_mainContext) IsInstr_mainContext() {}

func NewInstr_mainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_mainContext {
	var p = new(Instr_mainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_main

	return p
}

func (s *Instr_mainContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_mainContext) Get_R_MAIN() antlr.Token { return s._R_MAIN }

func (s *Instr_mainContext) Set_R_MAIN(v antlr.Token) { s._R_MAIN = v }

func (s *Instr_mainContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instr_mainContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instr_mainContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instr_mainContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instr_mainContext) R_FUNCTION() antlr.TerminalNode {
	return s.GetToken(ChemsR_FUNCTION, 0)
}

func (s *Instr_mainContext) R_MAIN() antlr.TerminalNode {
	return s.GetToken(ChemsR_MAIN, 0)
}

func (s *Instr_mainContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_mainContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_mainContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEA, 0)
}

func (s *Instr_mainContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_mainContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_LLAVEC, 0)
}

func (s *Instr_mainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_mainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_mainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_main(s)
	}
}

func (s *Instr_mainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_main(s)
	}
}

func (p *Chems) Instr_main() (localctx IInstr_mainContext) {
	localctx = NewInstr_mainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ChemsRULE_instr_main)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(ChemsR_FUNCTION)
	}
	{
		p.SetState(61)

		var _m = p.Match(ChemsR_MAIN)

		localctx.(*Instr_mainContext)._R_MAIN = _m
	}
	{
		p.SetState(62)
		p.Match(ChemsTK_PARA)
	}
	{
		p.SetState(63)
		p.Match(ChemsTK_PARC)
	}
	{
		p.SetState(64)
		p.Match(ChemsTK_LLAVEA)
	}
	{
		p.SetState(65)

		var _x = p.Instrucciones()

		localctx.(*Instr_mainContext)._instrucciones = _x
	}
	{
		p.SetState(66)
		p.Match(ChemsTK_LLAVEC)
	}
	localctx.(*Instr_mainContext).instr = instruction.NewMain(localctx.(*Instr_mainContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instr_mainContext).Get_R_MAIN() == nil {
			return 0
		} else {
			return localctx.(*Instr_mainContext).Get_R_MAIN().GetLine()
		}
	}()), localctx.(*Instr_mainContext).Get_R_MAIN().GetColumn())

	return localctx
}

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_expre returns the _instr_expre rule contexts.
	Get_instr_expre() IInstr_expreContext

	// Set_instr_expre sets the _instr_expre rule contexts.
	Set_instr_expre(IInstr_expreContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	_instr_expre IInstr_expreContext
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) Get_instr_expre() IInstr_expreContext { return s._instr_expre }

func (s *ExpressionsContext) Set_instr_expre(v IInstr_expreContext) { s._instr_expre = v }

func (s *ExpressionsContext) GetP() interfaces.Expression { return s.p }

func (s *ExpressionsContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ExpressionsContext) Instr_expre() IInstr_expreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_expreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_expreContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (p *Chems) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ChemsRULE_expressions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)

		var _x = p.instr_expre(0)

		localctx.(*ExpressionsContext)._instr_expre = _x
	}
	localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_instr_expre().GetP()

	return localctx
}

// IInstr_expreContext is an interface to support dynamic dispatch.
type IInstr_expreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IInstr_expreContext

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// GetRight returns the right rule contexts.
	GetRight() IInstr_expreContext

	// SetLeft sets the left rule contexts.
	SetLeft(IInstr_expreContext)

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// SetRight sets the right rule contexts.
	SetRight(IInstr_expreContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsInstr_expreContext differentiates from other interfaces.
	IsInstr_expreContext()
}

type Instr_expreContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	left         IInstr_expreContext
	op           antlr.Token
	_primitivo   IPrimitivoContext
	_expressions IExpressionsContext
	right        IInstr_expreContext
}

func NewEmptyInstr_expreContext() *Instr_expreContext {
	var p = new(Instr_expreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instr_expre
	return p
}

func (*Instr_expreContext) IsInstr_expreContext() {}

func NewInstr_expreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_expreContext {
	var p = new(Instr_expreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instr_expre

	return p
}

func (s *Instr_expreContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_expreContext) GetOp() antlr.Token { return s.op }

func (s *Instr_expreContext) SetOp(v antlr.Token) { s.op = v }

func (s *Instr_expreContext) GetLeft() IInstr_expreContext { return s.left }

func (s *Instr_expreContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Instr_expreContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instr_expreContext) GetRight() IInstr_expreContext { return s.right }

func (s *Instr_expreContext) SetLeft(v IInstr_expreContext) { s.left = v }

func (s *Instr_expreContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Instr_expreContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instr_expreContext) SetRight(v IInstr_expreContext) { s.right = v }

func (s *Instr_expreContext) GetP() interfaces.Expression { return s.p }

func (s *Instr_expreContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instr_expreContext) AllInstr_expre() []IInstr_expreContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstr_expreContext)(nil)).Elem())
	var tst = make([]IInstr_expreContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstr_expreContext)
		}
	}

	return tst
}

func (s *Instr_expreContext) Instr_expre(i int) IInstr_expreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_expreContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstr_expreContext)
}

func (s *Instr_expreContext) TK_NOT() antlr.TerminalNode {
	return s.GetToken(ChemsTK_NOT, 0)
}

func (s *Instr_expreContext) TK_MENOS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENOS, 0)
}

func (s *Instr_expreContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Instr_expreContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARA, 0)
}

func (s *Instr_expreContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_expreContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(ChemsTK_PARC, 0)
}

func (s *Instr_expreContext) TK_MULT() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MULT, 0)
}

func (s *Instr_expreContext) TK_DIV() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DIV, 0)
}

func (s *Instr_expreContext) TK_MODULO() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MODULO, 0)
}

func (s *Instr_expreContext) TK_MAS() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAS, 0)
}

func (s *Instr_expreContext) TK_MENOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENOR, 0)
}

func (s *Instr_expreContext) TK_MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MENORIGUAL, 0)
}

func (s *Instr_expreContext) TK_MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAYORIGUAL, 0)
}

func (s *Instr_expreContext) TK_MAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_MAYOR, 0)
}

func (s *Instr_expreContext) TK_DIFIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_DIFIGUAL, 0)
}

func (s *Instr_expreContext) TK_IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsTK_IGUALIGUAL, 0)
}

func (s *Instr_expreContext) TK_AND() antlr.TerminalNode {
	return s.GetToken(ChemsTK_AND, 0)
}

func (s *Instr_expreContext) TK_OR() antlr.TerminalNode {
	return s.GetToken(ChemsTK_OR, 0)
}

func (s *Instr_expreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_expreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_expreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstr_expre(s)
	}
}

func (s *Instr_expreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstr_expre(s)
	}
}

func (p *Chems) Instr_expre() (localctx IInstr_expreContext) {
	return p.instr_expre(0)
}

func (p *Chems) instr_expre(_p int) (localctx IInstr_expreContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewInstr_expreContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IInstr_expreContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, ChemsRULE_instr_expre, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsTK_MENOS, ChemsTK_NOT:
		{
			p.SetState(73)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Instr_expreContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ChemsTK_MENOS || _la == ChemsTK_NOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Instr_expreContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(74)

			var _x = p.instr_expre(3)

			localctx.(*Instr_expreContext).left = _x
		}
		localctx.(*Instr_expreContext).p = expression.NewOperacion(localctx.(*Instr_expreContext).GetLeft().GetP(), (func() string {
			if localctx.(*Instr_expreContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Instr_expreContext).GetOp().GetText()
			}
		}()), nil, true, (func() int {
			if localctx.(*Instr_expreContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Instr_expreContext).GetOp().GetLine()
			}
		}()), localctx.(*Instr_expreContext).GetOp().GetColumn())

	case ChemsNUMBER, ChemsDOUBLE, ChemsCHAR, ChemsSTRING, ChemsBOOLEAN:
		{
			p.SetState(77)

			var _x = p.Primitivo()

			localctx.(*Instr_expreContext)._primitivo = _x
		}
		localctx.(*Instr_expreContext).p = localctx.(*Instr_expreContext).Get_primitivo().GetP()

	case ChemsTK_PARA:
		{
			p.SetState(80)
			p.Match(ChemsTK_PARA)
		}
		{
			p.SetState(81)

			var _x = p.Expressions()

			localctx.(*Instr_expreContext)._expressions = _x
		}
		{
			p.SetState(82)
			p.Match(ChemsTK_PARC)
		}
		localctx.(*Instr_expreContext).p = localctx.(*Instr_expreContext).Get_expressions().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewInstr_expreContext(p, _parentctx, _parentState)
				localctx.(*Instr_expreContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_instr_expre)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(88)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Instr_expreContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(ChemsTK_MULT-44))|(1<<(ChemsTK_DIV-44))|(1<<(ChemsTK_MODULO-44)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Instr_expreContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(89)

					var _x = p.instr_expre(8)

					localctx.(*Instr_expreContext).right = _x
				}
				localctx.(*Instr_expreContext).p = expression.NewOperacion(localctx.(*Instr_expreContext).GetLeft().GetP(), (func() string {
					if localctx.(*Instr_expreContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Instr_expreContext).GetOp().GetText()
					}
				}()), localctx.(*Instr_expreContext).GetRight().GetP(), false, (func() int {
					if localctx.(*Instr_expreContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Instr_expreContext).GetOp().GetLine()
					}
				}()), localctx.(*Instr_expreContext).GetOp().GetColumn())

			case 2:
				localctx = NewInstr_expreContext(p, _parentctx, _parentState)
				localctx.(*Instr_expreContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_instr_expre)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(93)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Instr_expreContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsTK_MAS || _la == ChemsTK_MENOS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Instr_expreContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(94)

					var _x = p.instr_expre(7)

					localctx.(*Instr_expreContext).right = _x
				}
				localctx.(*Instr_expreContext).p = expression.NewOperacion(localctx.(*Instr_expreContext).GetLeft().GetP(), (func() string {
					if localctx.(*Instr_expreContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Instr_expreContext).GetOp().GetText()
					}
				}()), localctx.(*Instr_expreContext).GetRight().GetP(), false, (func() int {
					if localctx.(*Instr_expreContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Instr_expreContext).GetOp().GetLine()
					}
				}()), localctx.(*Instr_expreContext).GetOp().GetColumn())

			case 3:
				localctx = NewInstr_expreContext(p, _parentctx, _parentState)
				localctx.(*Instr_expreContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_instr_expre)
				p.SetState(97)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(98)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Instr_expreContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ChemsTK_IGUALIGUAL-36))|(1<<(ChemsTK_MAYORIGUAL-36))|(1<<(ChemsTK_MENORIGUAL-36))|(1<<(ChemsTK_DIFIGUAL-36))|(1<<(ChemsTK_MAYOR-36))|(1<<(ChemsTK_MENOR-36)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Instr_expreContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(99)

					var _x = p.instr_expre(6)

					localctx.(*Instr_expreContext).right = _x
				}
				localctx.(*Instr_expreContext).p = expression.NewOperacion(localctx.(*Instr_expreContext).GetLeft().GetP(), (func() string {
					if localctx.(*Instr_expreContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Instr_expreContext).GetOp().GetText()
					}
				}()), localctx.(*Instr_expreContext).GetRight().GetP(), false, (func() int {
					if localctx.(*Instr_expreContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Instr_expreContext).GetOp().GetLine()
					}
				}()), localctx.(*Instr_expreContext).GetOp().GetColumn())

			case 4:
				localctx = NewInstr_expreContext(p, _parentctx, _parentState)
				localctx.(*Instr_expreContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_instr_expre)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(103)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Instr_expreContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsTK_AND || _la == ChemsTK_OR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Instr_expreContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(104)

					var _x = p.instr_expre(5)

					localctx.(*Instr_expreContext).right = _x
				}
				localctx.(*Instr_expreContext).p = expression.NewOperacion(localctx.(*Instr_expreContext).GetLeft().GetP(), (func() string {
					if localctx.(*Instr_expreContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Instr_expreContext).GetOp().GetText()
					}
				}()), localctx.(*Instr_expreContext).GetRight().GetP(), false, (func() int {
					if localctx.(*Instr_expreContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Instr_expreContext).GetOp().GetLine()
					}
				}()), localctx.(*Instr_expreContext).GetOp().GetColumn())

			}

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_DOUBLE returns the _DOUBLE token.
	Get_DOUBLE() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_BOOLEAN returns the _BOOLEAN token.
	Get_BOOLEAN() antlr.Token

	// Get_CHAR returns the _CHAR token.
	Get_CHAR() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_DOUBLE sets the _DOUBLE token.
	Set_DOUBLE(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_BOOLEAN sets the _BOOLEAN token.
	Set_BOOLEAN(antlr.Token)

	// Set_CHAR sets the _CHAR token.
	Set_CHAR(antlr.Token)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	p        interfaces.Expression
	_NUMBER  antlr.Token
	_DOUBLE  antlr.Token
	_STRING  antlr.Token
	_BOOLEAN antlr.Token
	_CHAR    antlr.Token
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_primitivo
	return p
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivoContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *PrimitivoContext) Get_DOUBLE() antlr.Token { return s._DOUBLE }

func (s *PrimitivoContext) Get_STRING() antlr.Token { return s._STRING }

func (s *PrimitivoContext) Get_BOOLEAN() antlr.Token { return s._BOOLEAN }

func (s *PrimitivoContext) Get_CHAR() antlr.Token { return s._CHAR }

func (s *PrimitivoContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitivoContext) Set_DOUBLE(v antlr.Token) { s._DOUBLE = v }

func (s *PrimitivoContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitivoContext) Set_BOOLEAN(v antlr.Token) { s._BOOLEAN = v }

func (s *PrimitivoContext) Set_CHAR(v antlr.Token) { s._CHAR = v }

func (s *PrimitivoContext) GetP() interfaces.Expression { return s.p }

func (s *PrimitivoContext) SetP(v interfaces.Expression) { s.p = v }

func (s *PrimitivoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChemsNUMBER, 0)
}

func (s *PrimitivoContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ChemsDOUBLE, 0)
}

func (s *PrimitivoContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *PrimitivoContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ChemsBOOLEAN, 0)
}

func (s *PrimitivoContext) CHAR() antlr.TerminalNode {
	return s.GetToken(ChemsCHAR, 0)
}

func (s *PrimitivoContext) R_AS() antlr.TerminalNode {
	return s.GetToken(ChemsR_AS, 0)
}

func (s *PrimitivoContext) R_INT() antlr.TerminalNode {
	return s.GetToken(ChemsR_INT, 0)
}

func (s *PrimitivoContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(ChemsR_FLOAT, 0)
}

func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (p *Chems) Primitivo() (localctx IPrimitivoContext) {
	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ChemsRULE_primitivo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*PrimitivoContext)._NUMBER = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}

		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*PrimitivoContext)._DOUBLE = _m
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(116)

			var _m = p.Match(ChemsSTRING)

			localctx.(*PrimitivoContext)._STRING = _m
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(str, interfaces.STRING, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_STRING().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(118)

			var _m = p.Match(ChemsBOOLEAN)

			localctx.(*PrimitivoContext)._BOOLEAN = _m
		}

		// str:= (func() string { if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText() }}())[1:len((func() string { if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil { return "" } else { return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText() }}()))-1]
		exp, _ := strconv.ParseBool((func() string {
			if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText()
			}
		}()))
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEAN().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(120)

			var _m = p.Match(ChemsCHAR)

			localctx.(*PrimitivoContext)._CHAR = _m
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_CHAR() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CHAR().GetText()
			}
		}())[1]
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(string(str), interfaces.CHAR, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_CHAR() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_CHAR().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_CHAR().GetColumn())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(122)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*PrimitivoContext)._NUMBER = _m
		}
		{
			p.SetState(123)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(124)
			p.Match(ChemsR_INT)
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}

		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.INTEGER, (func() int {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(126)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*PrimitivoContext)._NUMBER = _m
		}
		{
			p.SetState(127)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(128)
			p.Match(ChemsR_FLOAT)
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}

		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.FLOAT, (func() int {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(130)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*PrimitivoContext)._DOUBLE = _m
		}
		{
			p.SetState(131)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(132)
			p.Match(ChemsR_INT)
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.INTEGER, (func() int {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(134)

			var _m = p.Match(ChemsDOUBLE)

			localctx.(*PrimitivoContext)._DOUBLE = _m
		}
		{
			p.SetState(135)
			p.Match(ChemsR_AS)
		}
		{
			p.SetState(136)
			p.Match(ChemsR_FLOAT)
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.FLOAT, (func() int {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())

	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *Instr_expreContext = nil
		if localctx != nil {
			t = localctx.(*Instr_expreContext)
		}
		return p.Instr_expre_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Chems) Instr_expre_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
