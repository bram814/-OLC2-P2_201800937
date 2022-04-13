#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5;

	
void main() {
	P = 0; H = 0;

	/************ Declaracion ************/
	stack[(int)0] = 1;
	/************ Relacional == ************/
	if(1 == 1) goto L0;
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
	t1 = P + 1;
	stack[(int)t1] = 2;
	/************ Declaracion ************/
	t2 = P + 2;
	stack[(int)t2] = 3;
	/************ Relacional == ************/
	if(2 == 2) goto L6;
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
	stack[(int)t4] = 4;
	/************ Identificador ************/
	t5 = stack[(int)1];
	/************ Printf Integer ************/
	printf("%d",(int)t5);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L11;
	L10:
	goto L11;
	L11:
	goto L5;
	L4:
	goto L5;
	L5:
	
	return;
}
