package main

import (
	"fmt"
	"math/rand"

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
	exercicio3()
	exercicio4()
	comparacao()
	condicional()
	arthures(10)
	ParOuImpar(10)
	slice(7)
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

	fmt.Printf("\n\na vale %v", a)
	fmt.Printf("\nb vale %v", b)
	fmt.Printf("\nv vale %v", c)
}

func exercicio3() {
	x := float64(42)
	y := "James Bound"
	z := true

	s := fmt.Sprintf("\n\n%.2f %v %v\n", x, y, z)
	fmt.Println(s)
}

func exercicio4() {
	type bolo int
	var d bolo
	fmt.Printf("O valor de d é %v e seu tipo: %T\n", d, d)
	d = 42
	fmt.Printf("Agora o valor de d é %v\n", d)
}

// Teste do if com meu pai

func comparacao() {
	x := rand.Intn(100)

	if x < 50 {
		fmt.Printf("\n%v: menor do que 50\n\n", x)
	} else {
		fmt.Printf("%v: maior ou igual a 50\n\n", x)
	}
}

func condicional() {
	i := 5
	if i < 100 {
		fmt.Println("If")
	} else {
		fmt.Println("Else")
	}
}

func arthures(n int) {

	for i := 0; i < n; i++ {
		fmt.Println("Arthur")

	}
}

func ParOuImpar(n int) {
	metade := n / 2
	dobroDaMetade := metade * 2

	par := dobroDaMetade == n

	if par {
		fmt.Println("par")
	} else {
		fmt.Println("impar")
	}
}

func slice(n int) {

	s := make([]int, n, n)

	for i := 0; i < n; i++ {

		s[i] = 14

	}

	fmt.Println(s)
}
