/------HEADER------/
#include <stdio.h>
#include <math.h>
float heap[30101999];
float stack[30101999];
float P;
float H;
float t0, t1, t2, t3, t4;

/------MAIN------/
 void main() { 
	P = 0; H = 0;

	/* ---- Inicio IF ---- */
	if (10 > 0) goto L0;
	goto L1;
	L0:
	t1 = P + 0;
	stack[(int)t1] = 3;
	/* ---- Inicio IF ---- */
	goto L2;
	L2:
	t3 = P + 1;
	stack[(int)t3] = 5;
	L3:
	goto L4;
	L1:
	t4 = P + 0;
	stack[(int)t4] = 4;
	L4:

	 return; 
 }