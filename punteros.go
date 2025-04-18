package main

import "fmt"

type TipoElemento struct {
	Nombre string
}

var elementoEjemplo = TipoElemento{
	Nombre: "Nose",
}

func MiFuncion(*TipoElemento) {

	var p *int

	i := 42

	p = &i

	fmt.Println(*p)

	*p = 21

	fmt.Println(*p)
}
