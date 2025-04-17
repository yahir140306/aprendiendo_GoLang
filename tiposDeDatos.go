package main

import "fmt"

func tiposDeDatos() {
	var nombrePersona string = `Juanchis`
	apellido := "Isidro"

	fmt.Println(nombrePersona)
	fmt.Println("Hola " + nombrePersona)
	fmt.Println("Apellido " + apellido)

	// int8 = 127

	var asoActual int16 = 2024
	edad := 40

	fmt.Println("Año: ", asoActual)
	fmt.Println("Edad: ", edad)
}
