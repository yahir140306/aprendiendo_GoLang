package main

import (
	"fmt"
	"math"
)

// las interfaces no se implementan al menos manualmente
// las interfaces se cumplen

type forma interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func imprimirArea(f forma) {
	fmt.Printf("El area es: %.2f \n", f.Area())
}

func interfaces() {
	c := Circulo{Radio: 5}

	imprimirArea(c)

}

// concurrencia

// go => GoRoutine
// GoRoutine => es un hilo de ejecucion ligero virtual

func concurrencia() {

}
