/*------HEADER------*/
#include <stdio.h>
#include <math.h>
float heap[30101999];
float stack[30101999];
float P;
float H;
float t0, t1, t2, t3, t4, t5, t6, t7;

/*------NATIVES------*/
void printString () {
	t2 = P + 1;
	t3 = stack[(int)t2];
	L1:
	t4 = heap[(int)t3];
	if (t4 == -1) goto L0; 
	printf("%c", (char)t4);
	t3 = t3 + 1;
	goto L1;
	L0:
	return; 
}


/*------MAIN------*/
 void main() { 
	P = 1; H = 0;
 	/* ----COMPILE---- */
	t0 = H  ;
	heap[(int)H] = 97;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	stack[(int)1] = t0;
	t1 = stack[(int)1];
	t6 = P + 1;
	t6 = t6 + 1;
	stack[(int)t6] = t1;
	P = P + 1;
	printString(); 
	t7 = stack[(int)P];
	P = P - 1;
	printf("%c", (char)10);

	 return; 
 }