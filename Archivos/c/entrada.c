#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16;

	/************ Fucniones ************/
	/************ Main ************/
	
int main() {
	P = 0; H = 0;

	/************ Declaracion ************/
	t0 = P + 0;
	stack[(int)t0] = 0;
	/************ While ************/
	L0:
	/************ Identificador ************/
	t1 = stack[(int)0];
	/************ Relacional >= ************/
	if(t1 >= 0) goto L3;
	goto L4;
	L3:
	t2 = 1 + 0;
	goto L5;
	L4:
	t2 = 0 + 0;
	goto L5;
	L5:
	if(t2 == 1) goto L1;
	goto L2;
	L1:
	/************ Identificador ************/
	t3 = stack[(int)0];
	/************ Relacional == ************/
	if(t3 == 0) goto L6;
	goto L7;
	L6:
	t4 = 1 + 0;
	goto L8;
	L7:
	t4 = 0 + 0;
	goto L8;
	L8:
	/************ Control - If ************/
	if(t4 == 1) goto L9;
	goto L10;
	L9:
	/************ Identificador ************/
	t5 = stack[(int)0];
	/************ Aritmetica + ************/
	t6 = t5 + 100;
	/************ Asignacion ************/
	stack[(int)0] = t6;
	goto L11;
	L10:
	/************ Control - If (else if) ************/
	/************ Identificador ************/
	t7 = stack[(int)0];
	/************ Relacional > ************/
	if(t7 > 50) goto L12;
	goto L13;
	L12:
	t8 = 1 + 0;
	goto L14;
	L13:
	t8 = 0 + 0;
	goto L14;
	L14:
	/************ Control - If ************/
	if(t8 == 1) goto L15;
	goto L16;
	L15:
	/************ Identificador ************/
	t10 = stack[(int)0];
	/************ Aritmetica / ************/
	t11 = t10 / 2;
	/************ Aritmetica - ************/
	t9 = t11 - 25;
	/************ Asignacion ************/
	stack[(int)0] = t9;
	goto L17;
	L16:
	/************ Control - If (else) ************/
	/************ Identificador ************/
	t13 = stack[(int)0];
	/************ Aritmetica / ************/
	t14 = t13 / 2;
	/************ Aritmetica - ************/
	t12 = t14 - 1;
	/************ Asignacion ************/
	stack[(int)0] = t12;
	goto L17;
	L17:
	goto L11;
	L11:
	/************ Identificador ************/
	t15 = stack[(int)0];
	/************ Printf Integer ************/
	printf("%f",(float)t15);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L0;
	L2:
	/************ Identificador ************/
	t16 = stack[(int)0];
	/************ Printf Integer ************/
	printf("%f",(float)t16);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return 0;
}
