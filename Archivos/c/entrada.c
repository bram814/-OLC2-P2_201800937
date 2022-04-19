#include <stdio.h>
#include <math.h>
double heap[23111998];
double stack[23111998];
double P;
double H;
double t0, t1, t2, t3, t4, t5, t6, t7;

	
/************ NATIVE PRINTF STRING ************/
	
void printfString(){
	t3 = P + 1;
	t4 = stack[(int) t3];
	L1:
	t5 = heap[(int)t4];
	if(t5 == -1) goto L0;
	printf("%c", (int)t5);
	t4 = t4 + 1;
	goto L1;
	L0:
	return;
}
	/************ Fucniones ************/
	/************ Main ************/
	
int main() {
	P = 0; H = 0;

	/************ Declaracion ************/
	t0 = P + 0;
	stack[(int)t0] = 0;
	/************ Identificador ************/
	t1 = stack[(int)0];
	/************ Printf Integer ************/
	printf("%d",(int)t1);
	/************ Salto de Linea \n ************/
	printf("%c",10);
	/************ Asignacion ************/
	stack[(int)0] = 500;
	/************ PRIMITIVO STRING ************/
	t2 = H + 0;
	heap[(int)H] = 78;
	H = H + 1;
	heap[(int)H] = 85;
	H = H + 1;
	heap[(int)H] = 76;
	H = H + 1;
	heap[(int)H] = 79;
	H = H + 1;
	heap[(int)H] = -1;
	H = H + 1;
	/************ Printf String ************/
	t6 = P + 1;
	t6 = t6 + 1;
	stack[(int)t6] = t2;
	P = P + 1;
	printfString();
	t7 = stack[(int)P];
	P = P - 1;
	/************ Salto de Linea \n ************/
	printf("%c",10);
	
	return 0;
}
