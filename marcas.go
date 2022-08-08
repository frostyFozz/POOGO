package main

import "fmt"

type Marcas struct {
	Nombre      string
	Modelo      string
	Anio        uint
	Precio      float64
	Companias   []string
	Referencias map[uint]string
}

func (c Marcas) PrintReferencias() {
	texto := "Las referencias del vehiculo son: "
	for _, Ref := range c.Referencias {
		texto += Ref + ", "
	}

	fmt.Printf("%+v", texto[:len(texto)-2])
}

func main() {
	Pedidos := Marcas{
		Nombre:    "Mazda",
		Modelo:    "Demio",
		Anio:      2009,
		Precio:    400,
		Companias: []string{"Viamar", "Grupo Peravia"},
		Referencias: map[uint]string{
			1: "Tipo de luces",
			2: "Cantidad de neumaticos",
			3: "Disenio",
		},
	}
	fmt.Println(Pedidos.Nombre)
	fmt.Printf("%+s", Pedidos.Modelo)
	Pedidos.PrintReferencias()
}
