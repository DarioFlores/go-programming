package main

import "fmt"

type persona struct {
	nombre string
}

func (p *persona) hablar() {
	fmt.Println("Hola! Soy " + p.nombre)
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func crearPersona(nombre string) *persona {
	return &persona{nombre: nombre}
}
func main() {
	p1 := persona{nombre: "Dario"}
	// diAlgo(p1) // ERROR
	diAlgo(&p1)

	p2 := new(persona)
	p2.nombre = "Exequiel"
	diAlgo(p2)
	// diAlgo(&p2) // ERROR

	var p3 *persona
	p3 = crearPersona("Gisella")
	diAlgo(p3)
	// diAlgo(&p3) // ERROR
}
