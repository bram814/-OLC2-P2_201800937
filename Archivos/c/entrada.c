#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5;

	
int main() {
	P = 0; H = 0;

	/************ Declaracion ************/
	stack[(int)0] = 0;
	/************ Loop ************/
	L0:
	/************ Identificador ************/
	t1 = stack[(int)0];
	/************ Aritmetica + ************/
	t2 = t1 + 1;
	/************ Asignacion ************/
	stack[(int)0] = t2;
	/************ Identificador ************/
	t3 = stack[(int)0];
	/************ Relacional == ************/
	if(t3 == 10) goto L2;
	goto L3;
	L2:
	t4 = 1 + 0;
	goto L4;
	L3:
	t4 = 0 + 0;
	goto L4;
	L4:
	/************ Control - If ************/
	if(t4 == 1) goto L5;
	goto L6;
	L5:
	/************ Break ************/
	goto L1;
	goto L7;
	L6:
	goto L7;
	L7:
	goto L0;
	L1:
	/************ Break temp ************/
	/************ Declaracion ************/
	stack[(int)1] = t0;
	/************ Identificador ************/
	t5 = stack[(int)1];
	/************ Printf Integer ************/
	printf("%d",(int)t5);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return 0;
}
