#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9;

	
void main() {
	P = 0; H = 0;

	/************ Relacional > ************/
	if(10 > 0) goto L0;
	goto L1;
	L0:
	t0 = 1 + 0;
	goto L2;
	L1:
	t0 = 0 + 0;
	goto L2;
	L2:
	/************ Control - If ************/
	if(t0 == 1) goto L3;
	goto L4;
	L3:
	/************ Declaracion ************/
	t1 = P + 0;
	stack[(int)t1] = 3;
	/************ Identificador ************/
	t2 = stack[(int)0];
	/************ Printf Integer ************/
	printf("%d",(int)t2);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ PRIMITIVO BOOLEAN ************/
	goto L6;
	goto L7;
	L6:
	t3 = 1 + 0;
	goto L8;
	L7:
	t3 = 0 + 0;
	goto L8;
	L8:
	/************ Control - If ************/
	if(t3 == 1) goto L9;
	goto L10;
	L9:
	/************ Declaracion ************/
	t4 = P + 1;
	stack[(int)t4] = 5;
	/************ Identificador ************/
	t5 = stack[(int)1];
	/************ Printf Integer ************/
	printf("%d",(int)t5);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ Asignacion ************/
	stack[(int)1] = 6;
	/************ Identificador ************/
	t6 = stack[(int)1];
	/************ Printf Integer ************/
	printf("%d",(int)t6);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L11;
	L10:
	goto L11;
	L11:
	/************ Asignacion ************/
	stack[(int)0] = 7;
	/************ Identificador ************/
	t7 = stack[(int)0];
	/************ Printf Integer ************/
	printf("%d",(int)t7);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L5;
	L4:
	/************ Control - If (else) ************/
	/************ Declaracion ************/
	t8 = P + 0;
	stack[(int)t8] = 4;
	/************ Identificador ************/
	t9 = stack[(int)0];
	/************ Printf Integer ************/
	printf("%d",(int)t9);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L5;
	L5:
	
	return;
}
