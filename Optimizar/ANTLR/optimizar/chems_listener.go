// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChemsListener is a complete listener for a parse tree produced by Chems.
type ChemsListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterHeaders is called when entering the headers production.
	EnterHeaders(c *HeadersContext)

	// EnterList_temps is called when entering the list_temps production.
	EnterList_temps(c *List_tempsContext)

	// EnterTemps is called when entering the temps production.
	EnterTemps(c *TempsContext)

	// EnterListFuncs is called when entering the listFuncs production.
	EnterListFuncs(c *ListFuncsContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterInstructions is called when entering the instructions production.
	EnterInstructions(c *InstructionsContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterAssign_stack is called when entering the assign_stack production.
	EnterAssign_stack(c *Assign_stackContext)

	// EnterAssign_heap is called when entering the assign_heap production.
	EnterAssign_heap(c *Assign_heapContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterIf_instr is called when entering the if_instr production.
	EnterIf_instr(c *If_instrContext)

	// EnterGoto_instr is called when entering the goto_instr production.
	EnterGoto_instr(c *Goto_instrContext)

	// EnterLabel_instr is called when entering the label_instr production.
	EnterLabel_instr(c *Label_instrContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpr_rel is called when entering the expr_rel production.
	EnterExpr_rel(c *Expr_relContext)

	// EnterExpr_arit is called when entering the expr_arit production.
	EnterExpr_arit(c *Expr_aritContext)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterExpr_valor is called when entering the expr_valor production.
	EnterExpr_valor(c *Expr_valorContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitHeaders is called when exiting the headers production.
	ExitHeaders(c *HeadersContext)

	// ExitList_temps is called when exiting the list_temps production.
	ExitList_temps(c *List_tempsContext)

	// ExitTemps is called when exiting the temps production.
	ExitTemps(c *TempsContext)

	// ExitListFuncs is called when exiting the listFuncs production.
	ExitListFuncs(c *ListFuncsContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitInstructions is called when exiting the instructions production.
	ExitInstructions(c *InstructionsContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitAssign_stack is called when exiting the assign_stack production.
	ExitAssign_stack(c *Assign_stackContext)

	// ExitAssign_heap is called when exiting the assign_heap production.
	ExitAssign_heap(c *Assign_heapContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitIf_instr is called when exiting the if_instr production.
	ExitIf_instr(c *If_instrContext)

	// ExitGoto_instr is called when exiting the goto_instr production.
	ExitGoto_instr(c *Goto_instrContext)

	// ExitLabel_instr is called when exiting the label_instr production.
	ExitLabel_instr(c *Label_instrContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpr_rel is called when exiting the expr_rel production.
	ExitExpr_rel(c *Expr_relContext)

	// ExitExpr_arit is called when exiting the expr_arit production.
	ExitExpr_arit(c *Expr_aritContext)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitExpr_valor is called when exiting the expr_valor production.
	ExitExpr_valor(c *Expr_valorContext)
}
