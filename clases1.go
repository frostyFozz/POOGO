package main

import (
	"fmt"
)

type Cursos struct {
	Name     string
	Precio   float64
	Gratis   bool
	usuarios []uint
	clases   map[uint]string
}

func main() {
	Go := Cursos{
		Name:     "Curso POO en Go",
		Precio:   20.90,
		Gratis:   false,
		usuarios: []uint{7, 19, 20, 18},
		clases: map[uint]string{
			1: "Introducción",
			2: "Presentación",
			3: "Clases de Objetos",
			4: "Utilidad de los Objetos",
		},
	}
	Python := Cursos{}
	Python.Name = "Programación Orientada a Objetos en python"
	Python.usuarios = []uint{7, 9, 20, 10}
	html := Cursos{Name: "Html desde Cero", Gratis: true}
	fmt.Println(Go.Precio)
	fmt.Printf("%+v\n", html)
	fmt.Printf("%+v\n", Python)
}
