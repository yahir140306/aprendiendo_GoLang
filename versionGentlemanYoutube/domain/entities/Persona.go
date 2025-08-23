package entities

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func (p Persona) Saludar() string {
	return "Hola " + p.Nombre
}

func (p *Persona) CumplirAnios() {
	p.Edad++
}
