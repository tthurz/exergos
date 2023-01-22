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

	fmt.Printf("A metade de %v é: %v\n", valor, matematica.Metade(v))
}
