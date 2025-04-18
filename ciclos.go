package main

import "fmt"

func ciclos() {
	var arr = []int{5, 4, 3, 2, 1}

	for i, v := range arr {
		fmt.Println("Indice: %d, valor: %d", i, v)
	}

	// for _,  v := range arr {
	// 	fmt.Println("Indice: %d, valor: %d", i, v)
	// }

	// for i, _ := range arr {
	// 	fmt.Println("Indice: %d, valor: %d", i, v)
	// }

	var suma int = 0

	for i := 0; i <= 100; i++ {
		suma = suma + i
	}

	fmt.Println("Suma de 1 al 100: ", suma)
}
