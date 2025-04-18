package main

import "fmt"

func arreglos() {
	fmt.Println("Arreglo - Slices")

	var listasFrutas = [4]string{"Pera", "Manzana", "Plantano", "Piña"}
	// declararVariables = [cantidadElementos] tipoDato {elemenyos}

	fmt.Println(listasFrutas)
	fmt.Println(listasFrutas[0])

	// listaPaises := [3]string{"EUA", "Canada", "España"}
	listaPaises := []string{"EUA", "Canada", "España"}

	fmt.Println(listaPaises)

	// listaPaises[0] = "Mexico"
	listaPaises = append(listaPaises, "Mexico")

	listaPaises2 := listaPaises[1:3]
	fmt.Println(listaPaises2)

	listaPaises3 := listaPaises[2:]
	fmt.Println(listaPaises3)

	// fmt.Println(listaPaises)

	///

	/// position		{0, 1, 2, 3, 4, 5 }
	var primos = [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primos[1:4]
	fmt.Println(s) // [3 5 7]

	s = append(s, 17) // [3 5 7] -> [3 5 7 17]

	fmt.Println(s) // [3 5 7 17]

	fmt.Println(primos) // [2 3 5 7 17 11 13]

	//

	// var a [10]int : es lo mismo ⬇️
	// a[0:10]	a[0:]
	// a[:10]	a[:]

	// make

	a := make([]int, 0, 5)

	fmt.Println(a)
}
