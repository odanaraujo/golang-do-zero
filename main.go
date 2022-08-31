package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %2.0f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pow(c.raio, 2) * math.Pi
}

func main() {

	/*
	 */

	ret := retangulo{10, 20}
	escreverArea(ret)

	c := circulo{10}
	escreverArea(c)
}
