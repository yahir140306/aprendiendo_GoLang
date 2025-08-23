package domain

import (
	"fmt"
	entities "versionGentlemanYoutube/domain/entities"
)

func App() {
	persona := entities.Persona{
		Nombre:   "Juanchis",
		Apellido: "Isidro",
		Edad:     20,
	}

	elemento := entities.TipoElemento{Nombre: "Elementos"}

	persona.Saludar()
	persona.CumplirAnios()
	fmt.Println(persona)
	fmt.Println(elemento)
}
