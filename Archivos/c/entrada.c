#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16;

	
/************ NATIVE PRINTF STRING ************/
	
void printfString(){
	t12 = P + 1;
	t13 = stack[(int) t12];
	L11:
	t14 = heap[(int)t13];
	if(t14 == -1) goto L10;
	printf("%c", (int)t14);
	t13 = t13 + 1;
	goto L11;
	L10:
	return;
}
	
int main() {
	P = 0; H = 0;

	/************ Declaracion ************/
	stack[(int)0] = 555;
	/************ PRIMITIVO STRING ************/
	t0 = H + 0;
	heap[(int)H] = 104;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 108;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Declaracion ************/
	stack[(int)1] = t0;
	/************ Identificador ************/
	t1 = stack[(int)0];
	/************ Ternario - Match ************/
	goto L0;
	/************ Match [CASE] ************/
	L3:
	/************ PRIMITIVO STRING ************/
	t3 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	goto L1;
	t2 = t3;
	L4:
	/************ PRIMITIVO STRING ************/
	t4 = H + 0;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	goto L1;
	t2 = t4;
	L5:
	/************ PRIMITIVO STRING ************/
	t5 = H + 0;
	heap[(int)H] = 98;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	goto L1;
	t2 = t5;
	L6:
	/************ PRIMITIVO STRING ************/
	t6 = H + 0;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	goto L1;
	t2 = t6;
	L7:
	/************ PRIMITIVO STRING ************/
	t7 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	goto L1;
	t2 = t7;
	L8:
	/************ PRIMITIVO STRING ************/
	t8 = H + 0;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	goto L1;
	t2 = t8;
	L9:
	/************ PRIMITIVO STRING ************/
	t9 = H + 0;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	goto L1;
	t2 = t9;
	/************ Match [DEFAULT] ************/
	L2:
	/************ Identificador ************/
	t10 = stack[(int)1];
	goto L1;
	t2 = t10;
	L0:
	
	if(t1 == 1) goto L3;
	if(t1 == 2) goto L4;
	if(t1 == 3) goto L5;
	if(t1 == 4) goto L6;
	if(t1 == 5) goto L7;
	if(t1 == 6) goto L8;
	if(t1 == 235) goto L9;
	if(t1 == t1) goto L2;

	L1:
	/************ Declaracion ************/
	stack[(int)2] = t2;
	/************ Identificador ************/
	t11 = stack[(int)2];
	/************ Printf String ************/
	t15 = P + 3;
	t15 = t15 + 1;
	stack[(int)t15] = t11;
	P = P + 3;
	printfString();
	t16 = stack[(int)P];
	P = P - 3;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return 0;
}
