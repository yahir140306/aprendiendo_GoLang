package main

import "fmt"

func suma(a int, b int) int {
	return a + b
}

func resultado(a int, b int) {
	fmt.Printf("La suma de %d y %d es: %d \n", a, b, suma(a, b))
}

// !

func presentaResultado(operacion string, a int, b int) {
	suma := func(a int, b int) int {
		return a + b
	}
	resta := func(a int, b int) int {
		return a - b
	}

	resultado := 0

	if operacion == "suma" {
		resultado = suma(a, b)
	} else if operacion == "resta" {
		resultado = resta(a, b)
	}

	fmt.Printf("Para a = %d y b = %d la %s, es resultado es: %d  \n", a, b, operacion, resultado)

}

// !

var funciones = map[string]func(int, int) int{
	"suma":           func(a int, b int) int { return a + b },
	"resta":          func(a int, b int) int { return a - b },
	"multiplicacion": func(a int, b int) int { return a * b },
	"division":       func(a int, b int) int { return a / b },
}

func presentaResultado2(operacion string, a int, b int) {

	f, exists := funciones[operacion]

	if !exists {
		fmt.Println("Operacion no valida.")
		return
	}

	fmt.Printf("Para a = %d y b = %d la %s, es resultado es: %d  \n", a, b, operacion, f(a, b))

}

func funcion() {
	resultado(5, 8)
	resultado(4, 1)
	resultado(7, 5)
	resultado(91, 47)

	fmt.Println("\n")

	// !

	presentaResultado("suma", 5, 8)
	presentaResultado("resta", 4, 1)
	presentaResultado("suma", 7, 5)
	presentaResultado("resta", 91, 47)

	//!
	fmt.Println("\n")

	presentaResultado2("suma", 5, 8)
	presentaResultado2("resta", 4, 1)
	presentaResultado2("multiplicacion", 7, 5)
	presentaResultado2("division", 91, 47)
}
