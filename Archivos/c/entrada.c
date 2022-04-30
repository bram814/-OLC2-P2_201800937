#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18, t19, t20, t21, t22, t23, t24, t25, t26, t27, t28, t29, t30, t31, t32, t33, t34, t35, t36, t37, t38, t39, t40, t41;

	/************ Fucniones ************/
	
void f() {
	/************ Identificador ************/
	t0 = P + 1;
	t1 = stack[(int)t0];
	/************ Identificador ************/
	t2 = P + 1;
	t3 = stack[(int)t2];
	/************ Relacional < ************/
	if(t3 < 2) goto L1;
	goto L2;
	L1:
	t4 = 1 + 0;
	goto L3;
	L2:
	t4 = 0 + 0;
	goto L3;
	L3:
	/************ Control - If ************/
	if(t4 == 1) goto L4;
	goto L5;
	L4:
	/************ Return ************/
	stack[(int)P] = 1;
	goto L0;
	goto L6;
	L5:
	/************ Control - If (else) ************/
	/************ Return ************/
	/************ Identificador ************/
	t5 = P + 1;
	t6 = stack[(int)t5];
	/************ Guardando Temporal ************/
	t7 = P + 2;
	stack[(int)t7] = t6;
	/************ Llamada de Funcion ************/
	/************ Identificador ************/
	t10 = P + 1;
	t11 = stack[(int)t10];
	/************ Aritmetica - ************/
	t9 = t11 - 1;
	t8 = P + 4;
	stack[(int)t8] = t9;
	P = P + 3;
	f();
	t8 = stack[(int)P];
	P = P - 3;
	t12 = P + 2;
	t6 = stack[(int)t12];
	/************ Aritmetica * ************/
	t13 = t6 * t8;
	stack[(int)P] = t13;
	goto L0;
	goto L6;
	L6:
	L0:
	
	return;
}
	
void ack() {
	/************ Identificador ************/
	t14 = P + 1;
	t15 = stack[(int)t14];
	/************ Identificador ************/
	t16 = P + 2;
	t17 = stack[(int)t16];
	/************ Identificador ************/
	t18 = P + 1;
	t19 = stack[(int)t18];
	/************ Relacional == ************/
	if(t19 == 0) goto L8;
	goto L9;
	L8:
	t20 = 1 + 0;
	goto L10;
	L9:
	t20 = 0 + 0;
	goto L10;
	L10:
	/************ Control - If ************/
	if(t20 == 1) goto L11;
	goto L12;
	L11:
	/************ Return ************/
	/************ Identificador ************/
	t21 = P + 2;
	t22 = stack[(int)t21];
	/************ Aritmetica + ************/
	t23 = t22 + 1;
	stack[(int)P] = t23;
	goto L7;
	goto L13;
	L12:
	/************ Control - If (else if) ************/
	/************ Identificador ************/
	t24 = P + 2;
	t25 = stack[(int)t24];
	/************ Relacional == ************/
	if(t25 == 0) goto L14;
	goto L15;
	L14:
	t26 = 1 + 0;
	goto L16;
	L15:
	t26 = 0 + 0;
	goto L16;
	L16:
	/************ Control - If ************/
	if(t26 == 1) goto L17;
	goto L18;
	L17:
	/************ Return ************/
	/************ Llamada de Funcion ************/
	/************ Identificador ************/
	t29 = P + 1;
	t30 = stack[(int)t29];
	/************ Aritmetica - ************/
	t28 = t30 - 1;
	t27 = P + 4;
	stack[(int)t27] = t28;
	t27 = t27 + 1;
	stack[(int)t27] = 1;
	P = P + 3;
	ack();
	t27 = stack[(int)P];
	P = P - 3;
	stack[(int)P] = t27;
	goto L7;
	goto L19;
	L18:
	/************ Control - If (else) ************/
	/************ Return ************/
	/************ Llamada de Funcion ************/
	/************ Identificador ************/
	t33 = P + 1;
	t34 = stack[(int)t33];
	/************ Aritmetica - ************/
	t32 = t34 - 1;
	t31 = P + 4;
	stack[(int)t31] = t32;
	/************ Llamada de Funcion ************/
	/************ Identificador ************/
	t36 = P + 1;
	t37 = stack[(int)t36];
	t35 = P + 4;
	stack[(int)t35] = t37;
	/************ Identificador ************/
	t39 = P + 2;
	t40 = stack[(int)t39];
	/************ Aritmetica - ************/
	t38 = t40 - 1;
	t35 = t35 + 1;
	stack[(int)t35] = t38;
	P = P + 3;
	ack();
	t35 = stack[(int)P];
	P = P - 3;
	t31 = t31 + 1;
	stack[(int)t31] = t35;
	P = P + 3;
	ack();
	t31 = stack[(int)P];
	P = P - 3;
	stack[(int)P] = t31;
	goto L7;
	goto L19;
	L19:
	goto L13;
	L13:
	L7:
	
	return;
}
	/************ Main ************/
	
int main() {
	P = 0; H = 0;

	/************ Llamada de Funcion ************/
	t41 = P + 1;
	stack[(int)t41] = 3;
	t41 = t41 + 1;
	stack[(int)t41] = 0;
	P = P + 0;
	ack();
	t41 = stack[(int)P];
	P = P - 0;
	/************ Printf Integer ************/
	printf("%d",(int)t41);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return 0;
}
