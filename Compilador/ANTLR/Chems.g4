parser grammar Chems;



options {
  tokenVocab = ChemsLexer;
}


@header {
    import "OLC2/Compilador/interfaces"
    import "OLC2/Compilador/expression"
    import "OLC2/Compilador/instruction"
    import "OLC2/Compilador/instruction/loops"
    import "OLC2/Compilador/instruction/casteo"
    import "OLC2/Compilador/instruction/nativa"
    import "OLC2/Compilador/instruction/control"
    import "OLC2/Compilador/instruction/structs"
    import "OLC2/Compilador/instruction/arrays"
    import "OLC2/Compilador/instruction/variable"
    import "OLC2/Compilador/instruction/ternario"
    import "OLC2/Compilador/instruction/function"
    import "OLC2/Compilador/instruction/transferencia"
    import arrayList "github.com/colegno/arraylist"
}


start returns [*arrayList.List lista]
  : instrucciones {$lista = $instrucciones.l}
;

instrucciones returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : list +=instruccion+   {
        listInt := localctx.(*InstruccionesContext).GetList()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

end_instr returns [int v]
  : TK_PUNTOCOMA            { $v = 1}
  |                         { $v = 0}
;


instruccion returns [interfaces.Instruction instr]
  : instr_println end_instr                 { $instr = $instr_println.instr               }
  | instr_main                              { $instr = $instr_main.instr                  }
  | instr_structs_declaration end_instr     { $instr = $instr_structs_declaration.instr   }
  | instr_declaracion                       { $instr = $instr_declaracion.instr           }
  | instr_asignacion                        { $instr = $instr_asignacion.instr            }
  | instr_structs_assignment end_instr      { $instr = $instr_structs_assignment.instr    }
  | instr_if                                { $instr = $instr_if.instr                    }
  | instr_match                             { $instr = $instr_match.instr                 }
  | instr_while                             { $instr = $instr_while.instr                 }
  | instr_for_in                            { $instr = $instr_for_in.instr                }
  | instr_loop                              { $instr = $instr_loop.instr                  }
  | instr_break end_instr                   { $instr = $instr_break.instr                 }
  | instr_continue end_instr                { $instr = $instr_continue.instr              }
  | instr_func                              { $instr = $instr_func.instr                  }
  | instr_llamada end_instr                 { $instr = $instr_llamada.instr               }
  | instr_return end_instr                  { $instr = $instr_return.instr                }
  | instr_structs_decla                     { $instr = $instr_structs_decla.instr         }
  | instr_arrays end_instr                  { $instr = $instr_arrays.instr                }
;


/******************************** [PRINTLN!] ********************************/
instr_println returns [interfaces.Instruction instr]
  : R_PRINTLN TK_PARA primitivo TK_PARC TK_PUNTOCOMA                                 { $instr = instruction.NewPrintln(nil, $primitivo.p,       "-1",         $R_PRINTLN.line, localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn()) }
  | R_PRINTLN TK_PARA STRING TK_COMA list_expression TK_PARC TK_PUNTOCOMA            { $instr = instruction.NewPrintln($list_expression.l, nil, $STRING.text[1:len($STRING.text)-1], $R_PRINTLN.line, localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn()) }
;  

/******************************** [MAIN] ********************************/
instr_main returns [interfaces.Instruction instr]
  : R_FUNCTION R_MAIN TK_PARA TK_PARC TK_LLAVEA instrucciones TK_LLAVEC          { $instr = instruction.NewMain($instrucciones.l, $R_MAIN.line, localctx.(*Instr_mainContext).Get_R_MAIN().GetColumn()) }
;


/******************************** [DECLARACION][VARIABLE] ********************************/
instr_declaracion returns [interfaces.Instruction instr]
  : R_LET R_MUT ID TK_IGUAL expressions TK_PUNTOCOMA                           { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $expressions.p, true,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET R_MUT ID TK_DOSPUNTOS instr_tipo TK_PUNTOCOMA                        { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp,  nil,           true,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET R_MUT ID TK_DOSPUNTOS instr_tipo TK_IGUAL expressions TK_PUNTOCOMA   { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $expressions.p, true,  $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET ID TK_IGUAL expressions TK_PUNTOCOMA                                 { $instr = variable.NewDeclaration($ID.text, interfaces.NULL,      $expressions.p, false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET ID TK_DOSPUNTOS instr_tipo TK_PUNTOCOMA                              { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, nil,            false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  | R_LET ID TK_DOSPUNTOS instr_tipo TK_IGUAL expressions TK_PUNTOCOMA         { $instr = variable.NewDeclaration($ID.text, $instr_tipo.tipo_exp, $expressions.p, false, $R_LET.line, localctx.(*Instr_declaracionContext).Get_R_LET().GetColumn()) }
  
;

/******************************** [ASIGNACION][VARIABLE] ********************************/
instr_asignacion returns [interfaces.Instruction instr]
  : ID TK_IGUAL expressions TK_PUNTOCOMA                                       { $instr = variable.NewAssignment($ID.text, $expressions.p, $ID.line, localctx.(*Instr_asignacionContext).Get_ID().GetColumn()) }
;

/******************************** [CONTROL][IF] ********************************/
instr_if returns [interfaces.Instruction instr]
  : R_IF expressions TK_LLAVEA left_instr=instrucciones TK_LLAVEC                                                       { $instr = control.NewIf($expressions.p, $left_instr.l, nil, nil,               $R_IF.line, localctx.(*Instr_ifContext).Get_R_IF().GetColumn()) }
  | R_IF expressions TK_LLAVEA left_instr=instrucciones TK_LLAVEC R_ELSE TK_LLAVEA right_intr=instrucciones TK_LLAVEC   { $instr = control.NewIf($expressions.p, $left_instr.l, $right_intr.l, nil,     $R_IF.line, localctx.(*Instr_ifContext).Get_R_IF().GetColumn()) }
  | R_IF expressions TK_LLAVEA left_instr=instrucciones TK_LLAVEC R_ELSE instr_else_if                                  { $instr = control.NewIf($expressions.p, $left_instr.l, nil, $instr_else_if.l,  $R_IF.line, localctx.(*Instr_ifContext).Get_R_IF().GetColumn()) }
;

instr_else_if returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e +=instr_if*  {
        listInt := localctx.(*Instr_else_ifContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

/******************************** [TERNARIO][IF] ********************************/
instr_ternario returns [interfaces.Expression p]
  : R_IF cond=expressions TK_LLAVEA left_instr=expressions TK_LLAVEC                                                      { $p = ternario.NewIf($cond.p, $left_instr.p, nil, nil,                       $R_IF.line, localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn()) }
  | R_IF cond=expressions TK_LLAVEA left_instr=expressions TK_LLAVEC R_ELSE TK_LLAVEA right_intr=expressions TK_LLAVEC    { $p = ternario.NewIf($cond.p, $left_instr.p, $right_intr.p, nil,             $R_IF.line, localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn()) }
  | R_IF cond=expressions TK_LLAVEA left_instr=expressions TK_LLAVEC R_ELSE instr_else_if_ternario                        { $p = ternario.NewIf($cond.p, $left_instr.p, nil, $instr_else_if_ternario.l, $R_IF.line, localctx.(*Instr_ternarioContext).Get_R_IF().GetColumn()) }
;

instr_else_if_ternario returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e +=instr_ternario*  {
        listInt := localctx.(*Instr_else_if_ternarioContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetP())
        }
    }
;

/******************************** [CONTROL][MATCH] ********************************/
instr_match returns [interfaces.Instruction instr]
  : R_MATCH expressions TK_LLAVEA list_case block_default TK_LLAVEC        { $instr = control.NewMatch($expressions.p, $list_case.l, $block_default.l, $R_MATCH.line, localctx.(*Instr_matchContext).Get_R_MATCH().GetColumn()) }
  | R_MATCH expressions TK_LLAVEA block_default TK_LLAVEC                  { $instr = control.NewMatch($expressions.p, nil, $block_default.l,          $R_MATCH.line, localctx.(*Instr_matchContext).Get_R_MATCH().GetColumn()) }
;

list_case returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += instr_case+  {
        listInt := localctx.(*List_caseContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

/*  CASE  */
instr_case returns [interfaces.Expression instr]
  : list_expre_case TK_IGUALMAYOR TK_LLAVEA instrucciones TK_LLAVEC     { $instr = control.NewCase(nil, $list_expre_case.l, $instrucciones.l) }
  | list_expre_case TK_IGUALMAYOR block_instr_match TK_COMA             { $instr = control.NewCase(nil, $list_expre_case.l, $block_instr_match.l) }
;

/* List Expression Case */
list_expre_case returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_case+  {
        listInt := localctx.(*List_expre_caseContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

block_case returns [interfaces.Expression instr]
  : expressions TK_BARRA                                                         { $instr =  control.NewCase($expressions.p, nil, nil)}
  | expressions                                                                  { $instr =  control.NewCase($expressions.p, nil, nil)}
;

block_instr_match returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : list += instruccion   {
        listInt := localctx.(*Block_instr_matchContext).GetList()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

/*  DEFAULT  */
instr_default returns [interfaces.Instruction instr]
  : ID TK_IGUALMAYOR TK_LLAVEA instrucciones TK_LLAVEC           { $instr = control.NewDefault($instrucciones.l) }
  | ID TK_IGUALMAYOR block_instr_match TK_COMA                   { $instr = control.NewDefault($block_instr_match.l) }
;

block_default returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += instr_default+  {
        listInt := localctx.(*Block_defaultContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

/******************************** [CONTROL][MATCH][TERNARIO] ********************************/
instr_match_ter returns [interfaces.Expression p]
  : R_MATCH expressions TK_LLAVEA list_case_ternario instr_default_ter TK_LLAVEC    { $p = ternario.NewMatch($expressions.p, $list_case_ternario.l, $instr_default_ter.p, $R_MATCH.line, localctx.(*Instr_match_terContext).Get_R_MATCH().GetColumn()) }
  | R_MATCH expressions TK_LLAVEA instr_default_ter TK_LLAVEC                       { $p = ternario.NewMatch($expressions.p, nil, $instr_default_ter.p,                   $R_MATCH.line, localctx.(*Instr_match_terContext).Get_R_MATCH().GetColumn()) }
;

list_case_ternario returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += instr_case_ter+  {
        listInt := localctx.(*List_case_ternarioContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetP())
        }
    }
;

/*  CASE  */
instr_case_ter returns [interfaces.Expression p]
  : list_expre_case_ter TK_IGUALMAYOR expressions TK_COMA                     { $p = ternario.NewCase(nil, $list_expre_case_ter.l, $expressions.p) }
  | list_expre_case_ter TK_IGUALMAYOR TK_LLAVEA expressions TK_LLAVEC         { $p = ternario.NewCase(nil, $list_expre_case_ter.l, $expressions.p) }
;


/* List Expression Case */
list_expre_case_ter returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_case_ter+  {
        listInt := localctx.(*List_expre_case_terContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetP())
        }
    }
;

block_case_ter returns [interfaces.Expression p]
  : expressions TK_BARRA                                                         { $p =  ternario.NewCase($expressions.p, nil, nil)}
  | expressions                                                                  { $p =  ternario.NewCase($expressions.p, nil, nil)}
;


/*  DEFAULT  */
instr_default_ter returns [interfaces.Expression p]
  : ID TK_IGUALMAYOR expressions TK_COMA                   { $p = ternario.NewDefault($expressions.p) }
  | ID TK_IGUALMAYOR TK_LLAVEA expressions TK_LLAVEC       { $p = ternario.NewDefault($expressions.p) }
;


/******************************** [LOOP][WHILE] ********************************/
instr_while returns [interfaces.Instruction instr]
  : R_WHILE expressions TK_LLAVEA instrucciones TK_LLAVEC                           { $instr = loops.NewWhile($expressions.p, $instrucciones.l, $R_WHILE.line, localctx.(*Instr_whileContext).Get_R_WHILE().GetColumn()) }
;

/******************************** [LOOP][FOR] ********************************/
instr_for_in returns [interfaces.Instruction instr]
  : R_FOR ID R_IN left=expressions TK_DOBLEPUNTO right=expressions TK_LLAVEA instrucciones TK_LLAVEC     { $instr = loops.NewFor($ID.text, interfaces.INTEGER, $left.p, $right.p, $instrucciones.l, $R_FOR.line, localctx.(*Instr_for_inContext).Get_R_FOR().GetColumn()) }
  | R_FOR ID R_IN left=expressions TK_LLAVEA instrucciones TK_LLAVEC                                     { $instr = loops.NewFor($ID.text, interfaces.STRING,  $left.p, nil, $instrucciones.l,      $R_FOR.line, localctx.(*Instr_for_inContext).Get_R_FOR().GetColumn()) }

;

/******************************** [LOOP][LOOP] ********************************/
instr_loop returns [interfaces.Instruction instr]
  : R_LOOP TK_LLAVEA instrucciones TK_LLAVEC                           { $instr = loops.NewLoop($instrucciones.l, $R_LOOP.line, localctx.(*Instr_loopContext).Get_R_LOOP().GetColumn()) }
;

/******************************** [LOOP][LOOP][TERNARIO] ********************************/
instr_loop_ternario returns [interfaces.Expression p]
  : R_LOOP TK_LLAVEA instrucciones TK_LLAVEC                           { $p = ternario.NewLoop($instrucciones.l, $R_LOOP.line, localctx.(*Instr_loop_ternarioContext).Get_R_LOOP().GetColumn()) }
;


/******************************** [TRANSFERENCIA][BREAK]    ********************************/
instr_break returns [interfaces.Instruction instr]
  : R_BREAK                                { $instr = transferencia.NewBreak(nil,           $R_BREAK.line, localctx.(*Instr_breakContext).Get_R_BREAK().GetColumn()) }
  | R_BREAK expressions                    { $instr = transferencia.NewBreak($expressions.p, $R_BREAK.line, localctx.(*Instr_breakContext).Get_R_BREAK().GetColumn()) }
;

/******************************** [TRANSFERENCIA][CONTINUE]  ********************************/
instr_continue returns [interfaces.Instruction instr]
  : R_CONTINUE                             { $instr = transferencia.NewContinue($R_CONTINUE.line, localctx.(*Instr_continueContext).Get_R_CONTINUE().GetColumn()) }
;

/******************************** [TRANSFERENCIA][RETURN]  ********************************/
instr_return returns [interfaces.Instruction instr]
  : R_RETURN expressions                                { $instr = transferencia.NewReturn($expressions.p, $R_RETURN.line, localctx.(*Instr_returnContext).Get_R_RETURN().GetColumn()) }
; 

/******************************** [FUNCTION]  ********************************/
instr_func returns [interfaces.Instruction instr]
  : R_FUNCTION ID TK_PARA TK_PARC TK_LLAVEA instrucciones TK_LLAVEC                                                       { $instr = function.NewFunction($ID.text, nil, $instrucciones.l, interfaces.NULL,      $R_FUNCTION.line, localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn()) }
  | R_FUNCTION ID TK_PARA TK_PARC TK_MENOSMAYOR instr_tipo TK_LLAVEA instrucciones TK_LLAVEC                              { $instr = function.NewFunction($ID.text, nil, $instrucciones.l, $instr_tipo.tipo_exp, $R_FUNCTION.line, localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn()) }
  | R_FUNCTION ID TK_PARA list_function_parameters TK_PARC TK_LLAVEA instrucciones TK_LLAVEC                              { $instr = function.NewFunction($ID.text, $list_function_parameters.l, $instrucciones.l, interfaces.NULL,      $R_FUNCTION.line, localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn()) }
  | R_FUNCTION ID TK_PARA list_function_parameters TK_PARC TK_MENOSMAYOR instr_tipo TK_LLAVEA instrucciones TK_LLAVEC     { $instr = function.NewFunction($ID.text, $list_function_parameters.l, $instrucciones.l, $instr_tipo.tipo_exp, $R_FUNCTION.line, localctx.(*Instr_funcContext).Get_R_FUNCTION().GetColumn()) }
;

list_function_parameters returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_parameters_fn+  {
        listInt := localctx.(*List_function_parametersContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;


block_parameters_fn returns [interfaces.Instruction instr]
  : ID TK_DOSPUNTOS instr_tipo TK_COMA                    { $instr = function.NewListExpre($ID.text, $instr_tipo.tipo_exp, $ID.line, localctx.(*Block_parameters_fnContext).Get_ID().GetColumn()) }
  | ID TK_DOSPUNTOS instr_tipo                            { $instr = function.NewListExpre($ID.text, $instr_tipo.tipo_exp, $ID.line, localctx.(*Block_parameters_fnContext).Get_ID().GetColumn()) }
;


/******************************** [FUNCTION]  ********************************/
instr_llamada returns [interfaces.Instruction instr]
  : ID TK_PARA TK_PARC                            { $instr = function.NewLlamada($ID.text, nil,                $ID.line, localctx.(*Instr_llamadaContext).Get_ID().GetColumn()) }
  | ID TK_PARA list_expression TK_PARC            { $instr = function.NewLlamada($ID.text, $list_expression.l, $ID.line, localctx.(*Instr_llamadaContext).Get_ID().GetColumn()) }

;

/******************************** [FUNCTION]  ********************************/
instr_llamada_expre returns [interfaces.Expression p]
  : ID TK_PARA TK_PARC                            { $p = function.NewLlamadaExpresion($ID.text, nil,                $ID.line, localctx.(*Instr_llamada_expreContext).Get_ID().GetColumn()) }
  | ID TK_PARA list_expression TK_PARC            { $p = function.NewLlamadaExpresion($ID.text, $list_expression.l, $ID.line, localctx.(*Instr_llamada_expreContext).Get_ID().GetColumn()) }

;

/******************************** [STRUCT][DECLARATION]  ********************************/
instr_structs_decla returns [interfaces.Instruction instr]
  : R_STRUCT ID TK_LLAVEA list_struct_parameters TK_LLAVEC    { $instr = structs.NewDefinition($ID.text, interfaces.STRUCT, $list_struct_parameters.l, $R_STRUCT.line, localctx.(*Instr_structs_declaContext).Get_R_STRUCT().GetColumn()) }
;

list_struct_parameters returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_structs_parameters+  {
        listInt := localctx.(*List_struct_parametersContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

block_structs_parameters returns [interfaces.Instruction instr]
  : ID TK_DOSPUNTOS instr_tipo TK_COMA                    { $instr = structs.NewListExpre($ID.text, $instr_tipo.tipo_exp, $ID.line, localctx.(*Block_structs_parametersContext).Get_ID().GetColumn()) }
  | ID TK_DOSPUNTOS instr_tipo                            { $instr = structs.NewListExpre($ID.text, $instr_tipo.tipo_exp, $ID.line, localctx.(*Block_structs_parametersContext).Get_ID().GetColumn()) }
;

/******************************** [STRUCT][DECLARATION]  ********************************/
instr_arrays returns [interfaces.Instruction instr]
  : R_LET R_MUT ID TK_DOSPUNTOS list_arrays_definition TK_IGUAL list_arrays_datos   { $instr = arrays.NewDefinition($ID.text, true,  $list_arrays_definition.l, $list_arrays_datos.l, $R_LET.line, localctx.(*Instr_arraysContext).Get_R_LET().GetColumn()) }
  | R_LET       ID TK_DOSPUNTOS list_arrays_definition TK_IGUAL list_arrays_datos   { $instr = arrays.NewDefinition($ID.text, false, $list_arrays_definition.l, $list_arrays_datos.l, $R_LET.line, localctx.(*Instr_arraysContext).Get_R_LET().GetColumn()) }

;


list_arrays_datos returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_dimensiones_datos  {
        listInt := localctx.(*List_arrays_datosContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

block_dimensiones_datos returns [interfaces.Instruction instr]
  : block_array_dimensionUno_datos        { $instr = $block_array_dimensionUno_datos.instr }     
;


block_array_dimensionUno_datos returns [interfaces.Instruction instr]
  : TK_CORA list_array_dimUno TK_CORC   { $instr = arrays.NewDimUnoDatos($list_array_dimUno.l) }
;

list_array_dimUno returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_array_dimUno_datos+  {
        listInt := localctx.(*List_array_dimUnoContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

block_array_dimUno_datos returns [interfaces.Instruction instr]
  : expressions TK_COMA   { $instr = arrays.NewListDimUno($expressions.p)}
  | expressions           { $instr = arrays.NewListDimUno($expressions.p)}
;



list_arrays_definition returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_dimensiones+  {
        listInt := localctx.(*List_arrays_definitionContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;


block_dimensiones returns [interfaces.Instruction instr]
  : block_array_dimensionUno              { $instr = $block_array_dimensionUno.instr }
;


block_array_dimensionUno returns [interfaces.Instruction instr]
  : TK_CORA instr_tipo TK_PUNTOCOMA expressions TK_CORC   { $instr = arrays.NewDimUno($instr_tipo.tipo_exp, $expressions.p)}
;

/******************************** [STRUCT][DECLARACION] ********************************/
instr_arrays_identifier returns [interfaces.Expression p]
  : ID list_arrays_parameters_id                            { $p = arrays.NewIdentifier($ID.text, $list_arrays_parameters_id.l, $ID.line, localctx.(*Instr_arrays_identifierContext).Get_ID().GetColumn()) }
;

list_arrays_parameters_id returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_arrays_identifier+  {
        listInt := localctx.(*List_arrays_parameters_idContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

block_arrays_identifier returns [interfaces.Instruction instr]
  : TK_CORA expressions TK_CORC                            { $instr = arrays.NewListIdentifier($expressions.p)}
;

/******************************** [STRUCT][DECLARACION] ********************************/
instr_structs_declaration returns [interfaces.Instruction instr]
  : R_LET R_MUT left=ID TK_IGUAL right=ID TK_LLAVEA list_struct_parameters_decla TK_LLAVEC     { $instr = structs.NewDeclaration($left.text, $list_struct_parameters_decla.l, true,  $right.text, $R_LET.line, localctx.(*Instr_structs_declarationContext).Get_R_LET().GetColumn()) }
  | R_LET       left=ID TK_IGUAL right=ID TK_LLAVEA list_struct_parameters_decla TK_LLAVEC     { $instr = structs.NewDeclaration($left.text, $list_struct_parameters_decla.l, false, $right.text, $R_LET.line, localctx.(*Instr_structs_declarationContext).Get_R_LET().GetColumn()) } 
;

list_struct_parameters_decla returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_structs_parameters_decla+  {
        listInt := localctx.(*List_struct_parameters_declaContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

block_structs_parameters_decla returns [interfaces.Instruction instr]
  : ID TK_DOSPUNTOS expressions TK_COMA                    { $instr = structs.NewListExpreDecla($ID.text, $expressions.p, $ID.line, localctx.(*Block_structs_parameters_declaContext).Get_ID().GetColumn()) }
  | ID TK_DOSPUNTOS expressions                            { $instr = structs.NewListExpreDecla($ID.text, $expressions.p, $ID.line, localctx.(*Block_structs_parameters_declaContext).Get_ID().GetColumn()) }
;



/******************************** [STRUCT][DECLARACION] ********************************/
instr_structs_identifier returns [interfaces.Expression p]
  : ID list_struct_parameters_id               { $p = structs.NewIdentifier($ID.text, $list_struct_parameters_id.l, $ID.line, localctx.(*Instr_structs_identifierContext).Get_ID().GetColumn()) }
;

list_struct_parameters_id returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_structs_identifier+  {
        listInt := localctx.(*List_struct_parameters_idContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

block_structs_identifier returns [interfaces.Instruction instr]
  : TK_PUNTO ID                   { $instr = structs.NewListIdentifier($ID.text)}
;

/******************************** [STRUCT][ASIGNACION] ********************************/
instr_structs_assignment returns [interfaces.Instruction instr]
  : ID list_struct_parameters_id TK_IGUAL expressions        { $instr = structs.NewAssignment($ID.text, $list_struct_parameters_id.l, $expressions.p ,$ID.line, localctx.(*Instr_structs_assignmentContext).Get_ID().GetColumn()) }            
;

/******************************** [TIPO] ********************************/
instr_tipo returns [interfaces.TypeExpression tipo_exp]
  : R_INT               {$tipo_exp = interfaces.INTEGER}
  | R_FLOAT             {$tipo_exp = interfaces.FLOAT}
  | R_STRING            {$tipo_exp = interfaces.STRING}
  | TK_AMPERSAND R_STR  {$tipo_exp = interfaces.STR}
  | R_BOOL              {$tipo_exp = interfaces.BOOLEAN}
  | R_CHAR              {$tipo_exp = interfaces.CHAR}
;

/* List Expression Case */
list_expression returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_list_expression+  {
        listInt := localctx.(*List_expressionContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetP())
        }
    }
;


block_list_expression returns [interfaces.Expression p]
  : expressions TK_COMA                    { $p =  instruction.NewListExpre($expressions.p) }
  | expressions                            { $p =  instruction.NewListExpre($expressions.p) }
;


/******************************** [EXPRESIONES] ********************************/
expressions returns [interfaces.Expression p]
  : expre_log                                   { $p = $expre_log.p }
  | expre_rel                                   { $p = $expre_rel.p } 
  | expre_arit                                  { $p = $expre_arit.p }
  | expre_casteo                                { $p = $expre_casteo.p }

;

expre_log returns [interfaces.Expression p]
  : op='!' left=expre_log                                               { $p = expression.NewOperacion($left.p, $op.text, nil,      true,  $op.line, localctx.(*Expre_logContext).GetOp().GetColumn()) }
  | left=expre_log op=('&&'|'||') right=expre_log                       { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_logContext).GetOp().GetColumn()) }
  | expre_rel                                                           { $p = $expre_rel.p } 
;

expre_rel returns [interfaces.Expression p]
  : left=expre_rel op=('<'|'<='|'>='|'>'|'!='|'==') right=expre_rel     { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_relContext).GetOp().GetColumn()) }
  | expre_arit                                                          { $p = $expre_arit.p }
;

expre_arit returns [interfaces.Expression p]
  : op='-' left=expre_arit                                              { $p = expression.NewOperacion($left.p, $op.text, nil,      true,  $op.line, localctx.(*Expre_aritContext).GetOp().GetColumn()) }
  | left=expre_arit op=('*'|'/'|'%') right=expre_arit                   { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_aritContext).GetOp().GetColumn()) }
  | left=expre_arit op=('+'|'-') right=expre_arit                       { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_aritContext).GetOp().GetColumn()) }
  | expre_valor                                                         { $p = $expre_valor.p }
  | TK_PARA expressions TK_PARC                                         { $p = $expressions.p }
;


expre_valor returns [interfaces.Expression p]
  : primitivo                                                           { $p = $primitivo.p }
  
;

primitivo returns[interfaces.Expression p]
    :NUMBER {
              num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }

            $p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.NULL, $NUMBER.line, localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())
       } 
    |DOUBLE  {  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.NULL, $DOUBLE.line, localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())
              }
    |STRING { 
              str:= $STRING.text[1:len($STRING.text)-1]
              $p = expression.NewPrimitivo(str, interfaces.STR, interfaces.NULL, $STRING.line, localctx.(*PrimitivoContext).Get_STRING().GetColumn())
            
            }
            
    |BOOLEAN { 
              // str:= $BOOLEAN.text[1:len($BOOLEAN.text)-1]
              exp,_ := strconv.ParseBool($BOOLEAN.text)
              $p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.NULL, $BOOLEAN.line, localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())
            }

    |CHAR {

            str:= $CHAR.text[1]
            $p = expression.NewPrimitivo(string(str), interfaces.CHAR, interfaces.NULL, $CHAR.line, localctx.(*PrimitivoContext).Get_CHAR().GetColumn())
          
          }
    | instr_llamada_expre       { $p = $instr_llamada_expre.p }
    | instr_structs_identifier  { $p = $instr_structs_identifier.p }
    | instr_arrays_identifier   { $p = $instr_arrays_identifier.p }
    | ID                        { $p = variable.NewIdentifier($ID.text, $ID.line, localctx.(*PrimitivoContext).Get_ID().GetColumn()) }

    | nativa_expre              { $p = $nativa_expre.p }
    | primitivo_casteo          { $p = $primitivo_casteo.p }
    | instr_ternario            { $p = $instr_ternario.p }
    | instr_match_ter           { $p = $instr_match_ter.p }
    | instr_loop_ternario       { $p = $instr_loop_ternario.p }
;


primitivo_casteo returns[interfaces.Expression p]
  : NUMBER R_AS R_INT{
              num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }

            $p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.INTEGER, $NUMBER.line, localctx.(*Primitivo_casteoContext).Get_NUMBER().GetColumn())
       }
    |NUMBER R_AS R_FLOAT{
              num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }

            $p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.FLOAT, $NUMBER.line, localctx.(*Primitivo_casteoContext).Get_NUMBER().GetColumn())
       }  
    |DOUBLE  R_AS R_INT{  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.INTEGER, $DOUBLE.line, localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetColumn())
              }
    |DOUBLE  R_AS R_FLOAT{  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.FLOAT, $DOUBLE.line, localctx.(*Primitivo_casteoContext).Get_DOUBLE().GetColumn())
              }

    |BOOLEAN  R_AS R_INT { 
              // str:= $BOOLEAN.text[1:len($BOOLEAN.text)-1]
              exp,_ := strconv.ParseBool($BOOLEAN.text)
              $p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.INTEGER, $BOOLEAN.line, localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetColumn())
            }

    |BOOLEAN  R_AS R_FLOAT { 
              // str:= $BOOLEAN.text[1:len($BOOLEAN.text)-1]
              exp,_ := strconv.ParseBool($BOOLEAN.text)
              $p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.FLOAT, $BOOLEAN.line, localctx.(*Primitivo_casteoContext).Get_BOOLEAN().GetColumn())
            }
;

expre_casteo returns[interfaces.Expression p]
  : expre_arit R_AS type_casteo                              { $p = casteo.NewCasteo($expre_arit.p, $type_casteo.tipo_exp, $R_AS.line, localctx.(*Expre_casteoContext).Get_R_AS().GetColumn()) }
;

type_casteo returns[interfaces.TypeExpression tipo_exp]
  :  R_INT                                              { $tipo_exp = interfaces.INTEGER }
  |  R_FLOAT                                            { $tipo_exp = interfaces.FLOAT }
;


nativa_expre returns[interfaces.Expression p]
  : R_INT TK_DOSPUNTOS TK_DOSPUNTOS R_POW TK_PARA b=expressions TK_COMA e=expressions TK_PARC    { $p = nativa.NewPotencia($b.p, $e.p, interfaces.INTEGER, $R_POW.line,      localctx.(*Nativa_expreContext).Get_R_POW().GetColumn()) }
  | R_FLOAT TK_DOSPUNTOS TK_DOSPUNTOS R_POWF TK_PARA b=expressions TK_COMA e=expressions TK_PARC { $p = nativa.NewPotencia($b.p, $e.p, interfaces.FLOAT,   $R_POWF.line,     localctx.(*Nativa_expreContext).Get_R_POWF().GetColumn()) }
  | native_block_abs TK_PUNTO R_ABS TK_PARA TK_PARC                                              { $p = nativa.NewAbs($native_block_abs.p,                 $R_ABS.line,      localctx.(*Nativa_expreContext).Get_R_ABS().GetColumn()) }
  | native_block_abs TK_PUNTO R_TOSTRING TK_PARA TK_PARC                                         { $p = nativa.NewToString($native_block_abs.p,            $R_TOSTRING.line, localctx.(*Nativa_expreContext).Get_R_TOSTRING().GetColumn()) }
  | TK_AMPERSAND native_block_abs TK_PUNTO R_TOSTRING TK_PARA TK_PARC                            { $p = nativa.NewToString($native_block_abs.p,            $R_TOSTRING.line, localctx.(*Nativa_expreContext).Get_R_TOSTRING().GetColumn()) }
;


native_block_abs returns[interfaces.Expression p]
  : instr_llamada_expre       { $p = $instr_llamada_expre.p }
  | ID                        { $p = variable.NewIdentifier($ID.text, $ID.line, localctx.(*Native_block_absContext).Get_ID().GetColumn()) }
  |STRING { 
              str:= $STRING.text[1:len($STRING.text)-1]
              $p = expression.NewPrimitivo(str, interfaces.STRING, interfaces.NULL, $STRING.line, localctx.(*Native_block_absContext).Get_STRING().GetColumn())
            
            }
;