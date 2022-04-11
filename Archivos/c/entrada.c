#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6;

	
/************ NATIVE PRINTF STRING ************/
	
void printfString(){
	t2 = P + 1;
	t3 = stack[(int) t2];
	L7:
	t4 = heap[(int)t3];
	if(t4 == -1) goto L6;
	printf("%c", (int)t4);
	t3 = t3 + 1;
	goto L7;
	L6:
	return;
}
	
void main() {
	P = 0; H = 0;

	/************ Declaracion ************/
	stack[(int)0] = 2;
	/************ Identificador ************/
	t0 = stack[(int)0];
	/************ Match ************/
	goto L0;
	/************ Match [CASE] ************/
	L3:
	/************ Printf Integer ************/
	printf("%d",(int)1);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L1;
	L4:
	/************ Printf Integer ************/
	printf("%d",(int)2);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L1;
	L5:
	/************ Printf Integer ************/
	printf("%d",(int)2);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L1;
	/************ Match [DEFAULT] ************/
	L2:
	/************ PRIMITIVO STRING ************/
	t1 = H + 0;
	heap[(int)H] = 82;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 116;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 101;
	H = H + 1;
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 99;
	H = H + 1;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = 115;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t5 = P + 1;
	t5 = t5 + 1;
	stack[(int)t5] = t1;
	printfString();
	t6 = stack[(int)P];
	P = P - 1;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L1;
	L0:
	
	if(t0 == 1) goto L3;
	if(t0 == 6) goto L4;
	if(t0 == 7) goto L5;
	if(t0 == t0) goto L2;

	L1:
	
	return;
}
