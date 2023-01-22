package main

import (
	"fmt"

	"github.com/tthurz/exergos/matematica"
	"github.com/tthurz/exergos/util"
)

func main() {
	fmt.Println("hello world")

	x := util.ConcatenadorDoTur("para", "béns")

	fmt.Println(x)

	valor := 33

	v := float64(valor)

	fmt.Printf("\nA metade de %v é: %v\n", valor, matematica.Metade(v))

	exercicio1()
	exercicio2()
}

func exercicio1() {
	x := 42
	y := "James Bound"
	z := true
	fmt.Printf("\n%v, %v, %v", x, y, z)
	fmt.Printf("\n\n%v", x)
	fmt.Printf("\n%v", y)
	fmt.Printf("\n%v", z)
}

// variaveis para fazer o ex2
var a int
var b string
var c bool

func exercicio2() {

	fmt.Printf("\na vale %v", a)
	fmt.Printf("\nb vale %v", b)
	fmt.Printf("\nv vale %v", c)
}
