package main

import "fmt"

func maps() {

	mapa := map[string]string{
		"Colombia":  "Bogota",
		"Venezuela": "Caracas",
		"Brasil":    "Brasilia",
		"Chile":     "Santiago",
	}

	fmt.Println("Contenido: ", mapa)
	fmt.Println("Venezuela: ", mapa["Venezuela"])

	mapa["Argetina"] = "Buenos Aires"
	fmt.Println("Contenido: ", mapa)

	delete(mapa, "Veezuela")
	fmt.Println("Contenido: ", mapa)

}

func condiciones() {
	var edad int = 39

	if edad >= 60 {
		fmt.Println("Viejo")
	} else if edad >= 18 {
		fmt.Println("Joven")
	} else {
		fmt.Println("Pequeño")
	}
}
