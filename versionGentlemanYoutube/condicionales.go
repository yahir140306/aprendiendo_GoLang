package main

import "fmt"

func condicionales() {

	// condicionales
	edad := 20
	if edad < 18 {
		fmt.Printf("Eres menor de edad")
		return
	}

	fmt.Println("Eres mayor de edad")

	// ciclos

	for i := 0; i < 5; i++ {
		fmt.Println("Iteración:", i)
		fmt.Printf("Iteración: %d\n", i)
	}

	// while
	n := 0
	for n < 3 {
		fmt.Printf("Iteración: %d\n", n)
		n++
	}

	// infinito
	n = 0
	for {
		n++
		if n == 5 {
			continue
		}

		fmt.Printf("n en bucle infinito: %d\n", n)

		if n >= 7 {
			break
		}
	}

	// range
	slice := []string{"uno", "dos", "tres"}
	for index, value := range slice {
		fmt.Printf("Indice: %d, Valor: %s\n", index, value)
	}

	// switch
	valor := 2
	switch valor {
	case 1:
		fmt.Println("Es 1")
	case 2:
		fmt.Printf("Es 2")
	default:
		fmt.Printf("No es 1 ni 2")
	}
}
