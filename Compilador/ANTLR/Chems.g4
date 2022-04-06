parser grammar Chems;



options {
  tokenVocab = ChemsLexer;
}


@header {
    import "OLC2/Compilador/interfaces"
    import "OLC2/Compilador/expression"
    import "OLC2/Compilador/instruction"
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
  : instr_println end_instr       { $instr = $instr_println.instr   }
  | instr_main                    { $instr = $instr_main.instr      }

;


/******************************** [PRINTLN!] ********************************/
instr_println returns [interfaces.Instruction instr]
  : R_PRINTLN TK_PARA expressions TK_PARC TK_PUNTOCOMA                           { $instr = instruction.NewPrintln($expressions.p, $R_PRINTLN.line, localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn()) }
  | R_PRINTLN TK_PARA STRING TK_COMA expressions TK_PARC TK_PUNTOCOMA            { $instr = instruction.NewPrintln($expressions.p, $R_PRINTLN.line, localctx.(*Instr_printlnContext).Get_R_PRINTLN().GetColumn()) }
;  

instr_main returns [interfaces.Instruction instr]
  : R_FUNCTION R_MAIN TK_PARA TK_PARC TK_LLAVEA instrucciones TK_LLAVEC          { $instr = instruction.NewMain($instrucciones.l, $R_MAIN.line, localctx.(*Instr_mainContext).Get_R_MAIN().GetColumn()) }
;

expressions returns [interfaces.Expression p]
  : primitivo                                   {$p = $primitivo.p}
;



primitivo returns[interfaces.Expression p]
    :NUMBER {
              num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }

            $p = expression.NewPrimitivo(num, interfaces.INTEGER, $NUMBER.line, localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())
       } 
    |DOUBLE  {  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expression.NewPrimitivo(num, interfaces.FLOAT, $DOUBLE.line, localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())
              }
    |STRING { 
              str:= $STRING.text[1:len($STRING.text)-1]
              $p = expression.NewPrimitivo(str, interfaces.STRING, $STRING.line, localctx.(*PrimitivoContext).Get_STRING().GetColumn())
            
            }
            
    |BOOLEAN { 
              // str:= $BOOLEAN.text[1:len($BOOLEAN.text)-1]
              exp,_ := strconv.ParseBool($BOOLEAN.text)
              $p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, $BOOLEAN.line, localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())
            }

    |CHAR {

            str:= $CHAR.text[1]
            $p = expression.NewPrimitivo(string(str), interfaces.CHAR, $CHAR.line, localctx.(*PrimitivoContext).Get_CHAR().GetColumn())
          
          }
;