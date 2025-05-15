package main

import (
	"fmt"
)

func main() {
	// Enterosa
	entero := 10
	decimal := 3.14
	suma := entero + int(decimal)
	fmt.Println(suma)

	// Texto

	mensaje := "Hola, "
	concatenado := mensaje + "Juanchis"
	fmt.Println(concatenado)

	// bool
	esVerdadero := true
	fmt.Println(esVerdadero)

	// Arrays y Slices
	arrayFijo := [3]int{1, 2, 3}
	capacidad := cap(arrayFijo)
	fmt.Println(capacidad)

	sliceVariable := []int{4, 5, 6}
	sliceVariable = append(sliceVariable, 7)

	// Mapas
	diccionario := map[string]int{
		"clave1": 1,
		"clave2": 2,
	}

	fmt.Println(diccionario)

	// Structs
	type Persona struct {
	}
}
