#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17;

	
/************ NATIVE PRINTF STRING ************/
	
void printfString(){
	t3 = P + 1;
	t4 = stack[(int) t3];
	L5:
	t5 = heap[(int)t4];
	if(t5 == -1) goto L4;
	printf("%c", (int)t5);
	t4 = t4 + 1;
	goto L5;
	L4:
	return;
}
	
void main() {
	P = 0; H = 0;

	/************ PRIMITIVO BOOLEAN ************/
	goto L1;
	goto L0;
	L0:
	t1 = 1 + 0;
	goto L3;
	L1:
	t1 = 0 + 0;
	goto L3;
	L3:
	/************ Declaracion ************/
	stack[(int)0] = t0;
	/************ PRINTF ************/
	/************ concat format {} ************/
	t2 = H + 0;
	heap[(int)H] = 82;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t6 = P + 1;
	t6 = t6 + 1;
	stack[(int)t6] = t2;
	P = P + 1;
	printfString();
	t7 = stack[(int)P];
	P = P - 1;
	/************ concat format {} ************/
	t8 = H + 0;
	heap[(int)H] = 58;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t9 = P + 1;
	t9 = t9 + 1;
	stack[(int)t9] = t8;
	P = P + 1;
	printfString();
	t10 = stack[(int)P];
	P = P - 1;
	/************ concat format {} ************/
	t11 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t12 = P + 1;
	t12 = t12 + 1;
	stack[(int)t12] = t11;
	P = P + 1;
	printfString();
	t13 = stack[(int)P];
	P = P - 1;
	/************ PRIMITIVO BOOLEAN ************/
	/************ Aritmetica + ************/
	t14 = 0 + 1;
	/************ Printf Integer ************/
	printf("%d",(int)t14);
	/************ concat format {} ************/
	t15 = H + 0;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf {} ************/
	t16 = P + 1;
	t16 = t16 + 1;
	stack[(int)t16] = t15;
	P = P + 1;
	printfString();
	t17 = stack[(int)P];
	P = P - 1;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return;
}
