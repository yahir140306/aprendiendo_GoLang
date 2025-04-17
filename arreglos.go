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
}
