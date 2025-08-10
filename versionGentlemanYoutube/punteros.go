package main

import "fmt"

// * -> puntero
// & -> dirección de memoria
// copia los argumentos

// var a = 1
// referencia de nombre "a" que apunta a un espacio de memoria que contiene
// el valor 1

func incrementar(numero *int) {
	*numero++
}

func punteros() {
	valor := 10
	fmt.Println("Valor antes de incrementar:", valor)
	incrementar(&valor)
	fmt.Println("Valor despues de incrementar:", valor)

	// new()
	puntero := new(int) // puntero int inicializado en 0
	fmt.Println("Valor inicial con new: ", *puntero)
	*puntero = 20
	fmt.Println("Valor inicial con new: ", *puntero)

}
