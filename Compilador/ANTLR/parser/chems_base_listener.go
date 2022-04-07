// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChemsListener is a complete listener for a parse tree produced by Chems.
type BaseChemsListener struct{}

var _ ChemsListener = &BaseChemsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChemsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChemsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChemsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChemsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseChemsListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseChemsListener) ExitStart(ctx *StartContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseChemsListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseChemsListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterEnd_instr is called when production end_instr is entered.
func (s *BaseChemsListener) EnterEnd_instr(ctx *End_instrContext) {}

// ExitEnd_instr is called when production end_instr is exited.
func (s *BaseChemsListener) ExitEnd_instr(ctx *End_instrContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseChemsListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseChemsListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterInstr_println is called when production instr_println is entered.
func (s *BaseChemsListener) EnterInstr_println(ctx *Instr_printlnContext) {}

// ExitInstr_println is called when production instr_println is exited.
func (s *BaseChemsListener) ExitInstr_println(ctx *Instr_printlnContext) {}

// EnterInstr_main is called when production instr_main is entered.
func (s *BaseChemsListener) EnterInstr_main(ctx *Instr_mainContext) {}

// ExitInstr_main is called when production instr_main is exited.
func (s *BaseChemsListener) ExitInstr_main(ctx *Instr_mainContext) {}

// EnterInstr_declaracion is called when production instr_declaracion is entered.
func (s *BaseChemsListener) EnterInstr_declaracion(ctx *Instr_declaracionContext) {}

// ExitInstr_declaracion is called when production instr_declaracion is exited.
func (s *BaseChemsListener) ExitInstr_declaracion(ctx *Instr_declaracionContext) {}

// EnterInstr_asignacion is called when production instr_asignacion is entered.
func (s *BaseChemsListener) EnterInstr_asignacion(ctx *Instr_asignacionContext) {}

// ExitInstr_asignacion is called when production instr_asignacion is exited.
func (s *BaseChemsListener) ExitInstr_asignacion(ctx *Instr_asignacionContext) {}

// EnterInstr_tipo is called when production instr_tipo is entered.
func (s *BaseChemsListener) EnterInstr_tipo(ctx *Instr_tipoContext) {}

// ExitInstr_tipo is called when production instr_tipo is exited.
func (s *BaseChemsListener) ExitInstr_tipo(ctx *Instr_tipoContext) {}

// EnterList_expression is called when production list_expression is entered.
func (s *BaseChemsListener) EnterList_expression(ctx *List_expressionContext) {}

// ExitList_expression is called when production list_expression is exited.
func (s *BaseChemsListener) ExitList_expression(ctx *List_expressionContext) {}

// EnterBlock_list_expression is called when production block_list_expression is entered.
func (s *BaseChemsListener) EnterBlock_list_expression(ctx *Block_list_expressionContext) {}

// ExitBlock_list_expression is called when production block_list_expression is exited.
func (s *BaseChemsListener) ExitBlock_list_expression(ctx *Block_list_expressionContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseChemsListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseChemsListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterExpre_log is called when production expre_log is entered.
func (s *BaseChemsListener) EnterExpre_log(ctx *Expre_logContext) {}

// ExitExpre_log is called when production expre_log is exited.
func (s *BaseChemsListener) ExitExpre_log(ctx *Expre_logContext) {}

// EnterExpre_rel is called when production expre_rel is entered.
func (s *BaseChemsListener) EnterExpre_rel(ctx *Expre_relContext) {}

// ExitExpre_rel is called when production expre_rel is exited.
func (s *BaseChemsListener) ExitExpre_rel(ctx *Expre_relContext) {}

// EnterExpre_arit is called when production expre_arit is entered.
func (s *BaseChemsListener) EnterExpre_arit(ctx *Expre_aritContext) {}

// ExitExpre_arit is called when production expre_arit is exited.
func (s *BaseChemsListener) ExitExpre_arit(ctx *Expre_aritContext) {}

// EnterExpre_valor is called when production expre_valor is entered.
func (s *BaseChemsListener) EnterExpre_valor(ctx *Expre_valorContext) {}

// ExitExpre_valor is called when production expre_valor is exited.
func (s *BaseChemsListener) ExitExpre_valor(ctx *Expre_valorContext) {}

// EnterPrimitivo is called when production primitivo is entered.
func (s *BaseChemsListener) EnterPrimitivo(ctx *PrimitivoContext) {}

// ExitPrimitivo is called when production primitivo is exited.
func (s *BaseChemsListener) ExitPrimitivo(ctx *PrimitivoContext) {}
