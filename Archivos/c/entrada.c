#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6;

	
void main() {
	P = 0; H = 0;

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
	heap[(int)H] = 32;
	H = H + 1;
	heap[(int)H] = 109;
	H = H + 1;
	heap[(int)H] = 117;
	H = H + 1;
	heap[(int)H] = 110;
	H = H + 1;
	heap[(int)H] = 100;
	H = H + 1;
	heap[(int)H] = 111;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Declaracion ************/
	stack[(int)0] = t0;
	/************ For - String ************/
	/************ Identificador ************/
	t1 = stack[(int)0];
	/************ Identificador ************/
	t2 = P + 1;
	stack[(int)t2] = t1;
	t3 = stack[(int)t2];
	L0:
	t4 = heap[(int)t3];
	stack[(int)t2] = t3;
	t3 = t3 + 1;
	/************ Add If ************/
	if(t4 == -1) goto L1;
	/************ Identificador ************/
	t5 = stack[(int)1];
	/************ Printf Char ************/
	t6 = heap[(int)t5];
	printf("%c",(char)t6);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	goto L0;
	L1:
	
	return;
}
