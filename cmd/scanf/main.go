package main

import "fmt"

// 1° imprimir uma msg perguntando o nome do usuario

// 2° declarar uma variavel string

// 3° usar o scanf para ler a string e colocar na variavel

// 4° imprimir a frase "bom dia %v"

func main() {

	fmt.Println("Qual se nome?")

	x := ""

	fmt.Scanf("%s", &x)

	fmt.Printf("\nOlá %v", x)

}
