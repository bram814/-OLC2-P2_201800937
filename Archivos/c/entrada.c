#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4;

	
void main() {
	P = 0; H = 0;

	/************ Declaracion ************/
	stack[(int)0] = 10;
	/************ While ************/
	L0:
	/************ Identificador ************/
	t0 = stack[(int)0];
	/************ Relacional < ************/
	if(0 < t0) goto L3;
	goto L4;
	L3:
	t1 = 1 + 0;
	goto L5;
	L4:
	t1 = 0 + 0;
	goto L5;
	L5:
	if(t1 == 1) goto L1;
	goto L2;
	L1:
	/************ Identificador ************/
	t2 = stack[(int)0];
	/************ Printf Integer ************/
	printf("%d",(int)t2);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ Identificador ************/
	t4 = stack[(int)0];
	/************ Aritmetica - ************/
	t3 = t4 - 1;
	/************ Asignacion ************/
	stack[(int)0] = t3;
	goto L0;
	L2:
	
	return;
}
