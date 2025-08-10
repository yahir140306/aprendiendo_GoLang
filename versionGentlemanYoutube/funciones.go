package main

import (
	"errors"
	"fmt"
)

func suma(a, b int) int {
	return a + b
}

// funcion que devuelve mas de un valor
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		// fmt.Errorf("No se puede dividir por 0")
		return 0, errors.New("No se puede dividir por 0")
	}
	cociente := a / b
	// resto := float64(int(a) % int(b))
	return cociente, nil
}

// funcion con numero variable de argumentos
func imprimirNombres(nombres ...string) {
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
}

// ejemplo de closure
func contador() func() int {
	c := 0

	return func() int {
		c++
		return c
	}
}

type Rectangulo struct {
	Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func funciones() {
	// cociente, error := dividir(10, 0)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }
	// fmt.Println(cociente)k

	// closure
	cont := contador()
	fmt.Println("Contador: ", cont())
	fmt.Println("Contador: ", cont())
	fmt.Println("Contador: ", cont())
	fmt.Println("Contador: ", cont())
	fmt.Println("Contador: ", cont())

	rect := Rectangulo{Ancho: 5, Alto: 10}
	fmt.Println("Área del rectángulo:", rect.Area())
}
