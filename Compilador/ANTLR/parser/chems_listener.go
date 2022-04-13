// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChemsListener is a complete listener for a parse tree produced by Chems.
type ChemsListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterEnd_instr is called when entering the end_instr production.
	EnterEnd_instr(c *End_instrContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterInstr_println is called when entering the instr_println production.
	EnterInstr_println(c *Instr_printlnContext)

	// EnterInstr_main is called when entering the instr_main production.
	EnterInstr_main(c *Instr_mainContext)

	// EnterInstr_declaracion is called when entering the instr_declaracion production.
	EnterInstr_declaracion(c *Instr_declaracionContext)

	// EnterInstr_asignacion is called when entering the instr_asignacion production.
	EnterInstr_asignacion(c *Instr_asignacionContext)

	// EnterInstr_if is called when entering the instr_if production.
	EnterInstr_if(c *Instr_ifContext)

	// EnterInstr_else_if is called when entering the instr_else_if production.
	EnterInstr_else_if(c *Instr_else_ifContext)

	// EnterInstr_ternario is called when entering the instr_ternario production.
	EnterInstr_ternario(c *Instr_ternarioContext)

	// EnterInstr_else_if_ternario is called when entering the instr_else_if_ternario production.
	EnterInstr_else_if_ternario(c *Instr_else_if_ternarioContext)

	// EnterInstr_match is called when entering the instr_match production.
	EnterInstr_match(c *Instr_matchContext)

	// EnterList_case is called when entering the list_case production.
	EnterList_case(c *List_caseContext)

	// EnterInstr_case is called when entering the instr_case production.
	EnterInstr_case(c *Instr_caseContext)

	// EnterList_expre_case is called when entering the list_expre_case production.
	EnterList_expre_case(c *List_expre_caseContext)

	// EnterBlock_case is called when entering the block_case production.
	EnterBlock_case(c *Block_caseContext)

	// EnterBlock_instr_match is called when entering the block_instr_match production.
	EnterBlock_instr_match(c *Block_instr_matchContext)

	// EnterInstr_default is called when entering the instr_default production.
	EnterInstr_default(c *Instr_defaultContext)

	// EnterBlock_default is called when entering the block_default production.
	EnterBlock_default(c *Block_defaultContext)

	// EnterInstr_while is called when entering the instr_while production.
	EnterInstr_while(c *Instr_whileContext)

	// EnterInstr_for_in is called when entering the instr_for_in production.
	EnterInstr_for_in(c *Instr_for_inContext)

	// EnterInstr_tipo is called when entering the instr_tipo production.
	EnterInstr_tipo(c *Instr_tipoContext)

	// EnterList_expression is called when entering the list_expression production.
	EnterList_expression(c *List_expressionContext)

	// EnterBlock_list_expression is called when entering the block_list_expression production.
	EnterBlock_list_expression(c *Block_list_expressionContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterExpre_log is called when entering the expre_log production.
	EnterExpre_log(c *Expre_logContext)

	// EnterExpre_rel is called when entering the expre_rel production.
	EnterExpre_rel(c *Expre_relContext)

	// EnterExpre_arit is called when entering the expre_arit production.
	EnterExpre_arit(c *Expre_aritContext)

	// EnterExpre_valor is called when entering the expre_valor production.
	EnterExpre_valor(c *Expre_valorContext)

	// EnterPrimitivo is called when entering the primitivo production.
	EnterPrimitivo(c *PrimitivoContext)

	// EnterPrimitivo_casteo is called when entering the primitivo_casteo production.
	EnterPrimitivo_casteo(c *Primitivo_casteoContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitEnd_instr is called when exiting the end_instr production.
	ExitEnd_instr(c *End_instrContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitInstr_println is called when exiting the instr_println production.
	ExitInstr_println(c *Instr_printlnContext)

	// ExitInstr_main is called when exiting the instr_main production.
	ExitInstr_main(c *Instr_mainContext)

	// ExitInstr_declaracion is called when exiting the instr_declaracion production.
	ExitInstr_declaracion(c *Instr_declaracionContext)

	// ExitInstr_asignacion is called when exiting the instr_asignacion production.
	ExitInstr_asignacion(c *Instr_asignacionContext)

	// ExitInstr_if is called when exiting the instr_if production.
	ExitInstr_if(c *Instr_ifContext)

	// ExitInstr_else_if is called when exiting the instr_else_if production.
	ExitInstr_else_if(c *Instr_else_ifContext)

	// ExitInstr_ternario is called when exiting the instr_ternario production.
	ExitInstr_ternario(c *Instr_ternarioContext)

	// ExitInstr_else_if_ternario is called when exiting the instr_else_if_ternario production.
	ExitInstr_else_if_ternario(c *Instr_else_if_ternarioContext)

	// ExitInstr_match is called when exiting the instr_match production.
	ExitInstr_match(c *Instr_matchContext)

	// ExitList_case is called when exiting the list_case production.
	ExitList_case(c *List_caseContext)

	// ExitInstr_case is called when exiting the instr_case production.
	ExitInstr_case(c *Instr_caseContext)

	// ExitList_expre_case is called when exiting the list_expre_case production.
	ExitList_expre_case(c *List_expre_caseContext)

	// ExitBlock_case is called when exiting the block_case production.
	ExitBlock_case(c *Block_caseContext)

	// ExitBlock_instr_match is called when exiting the block_instr_match production.
	ExitBlock_instr_match(c *Block_instr_matchContext)

	// ExitInstr_default is called when exiting the instr_default production.
	ExitInstr_default(c *Instr_defaultContext)

	// ExitBlock_default is called when exiting the block_default production.
	ExitBlock_default(c *Block_defaultContext)

	// ExitInstr_while is called when exiting the instr_while production.
	ExitInstr_while(c *Instr_whileContext)

	// ExitInstr_for_in is called when exiting the instr_for_in production.
	ExitInstr_for_in(c *Instr_for_inContext)

	// ExitInstr_tipo is called when exiting the instr_tipo production.
	ExitInstr_tipo(c *Instr_tipoContext)

	// ExitList_expression is called when exiting the list_expression production.
	ExitList_expression(c *List_expressionContext)

	// ExitBlock_list_expression is called when exiting the block_list_expression production.
	ExitBlock_list_expression(c *Block_list_expressionContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitExpre_log is called when exiting the expre_log production.
	ExitExpre_log(c *Expre_logContext)

	// ExitExpre_rel is called when exiting the expre_rel production.
	ExitExpre_rel(c *Expre_relContext)

	// ExitExpre_arit is called when exiting the expre_arit production.
	ExitExpre_arit(c *Expre_aritContext)

	// ExitExpre_valor is called when exiting the expre_valor production.
	ExitExpre_valor(c *Expre_valorContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)

	// ExitPrimitivo_casteo is called when exiting the primitivo_casteo production.
	ExitPrimitivo_casteo(c *Primitivo_casteoContext)
}
