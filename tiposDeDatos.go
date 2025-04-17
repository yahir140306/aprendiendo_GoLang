package main

import "fmt"

var numero int = 2

func tiposDeDatos() {
	var nombrePersona string = `Juanchis`
	apellido := "Isidro"

	fmt.Println(nombrePersona)
	fmt.Println("Hola " + nombrePersona)
	fmt.Println("Apellido " + apellido)

	// int8 = 127

	var asoActual int16 = 2024
	edad := 40
	i, j := 42, 2500

	fmt.Println("Año: ", asoActual)
	fmt.Println("Edad: ", edad)
	fmt.Println(i, j)
}

// ejemplo
// type Persona struct {
// 	Nombre   string
// 	Apellido string
// 	Edad     int
// }

// var persona = Persona{
// 	Nombre:   "Paco",
// 	Apellido: "Hernandez",
// 	Edad:     32,
// }

// func imprimir() {
// 	fmt.Println(persona.Nombre)
// }

// var persona2 = Persona{}
