package main

import (
	"fmt"
	"math"
)

// Uma interface é um tipo abstrato que define um conjunto de métodos.
//Uma interface precisa ter métodos definidos, exceto no caso das interfaces do tipo genérico.
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}
func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func main() {
	forma := retangulo{10, 15}
	forma2 := circulo{3}
	escreverArea(forma)
	escreverArea(forma2)
}
