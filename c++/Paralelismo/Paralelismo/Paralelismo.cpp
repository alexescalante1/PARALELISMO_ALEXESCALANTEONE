#include <iostream>
#include <omp.h>    ///habilitar OMP
#include <stdio.h>
#include <Windows.h>

using namespace std;

int SUMA(int a,int b) {
	return a+b;
}

int main(){
	int i,* Matriz1,* Matriz2,* MatrizR;
	Matriz1 = new int[5];Matriz2 = new int[5];MatrizR = new int[5];
	*(Matriz1 + 0) = 12;
	*(Matriz1 + 1) = 41;
	*(Matriz1 + 2) = 31;
	*(Matriz1 + 3) = 76;
	*(Matriz1 + 4) = 23;
	*(Matriz1 + 5) = NULL;
	
	*(Matriz2 + 0) = 2;
	*(Matriz2 + 1) = 32;
	*(Matriz2 + 2) = 21;
	*(Matriz2 + 3) = 5;
	*(Matriz2 + 4) = 8;
	*(Matriz2 + 5) = NULL;

	*(MatrizR + 5) = NULL;

	cout << "SUMA SECUENCIAL\n\n";

	for (i = 0; *(Matriz1 + i);i++) {
		cout <<SUMA(*(Matriz1 + i), *(Matriz2 + i))<< ", ";
	}

	cout << "\n\nSUMA EN PARALELO\n\n";

	#pragma omp parallel for default(none) num_threads(4)
	for (i = 0; i < 5; i++) {
		*(MatrizR + i) = SUMA(*(Matriz1 + i), *(Matriz2 + i));
		printf("Numero de Thread %d, Numero de iteracion %d: Resultado = %d\n", omp_get_thread_num(), i, *(MatrizR + i));
	}
	
	return 0;
}