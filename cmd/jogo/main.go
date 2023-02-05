package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	jogo()
}

func jogo() {

	rand.Seed(time.Now().UnixNano())
	x := (rand.Intn(6))
	d := x

	i := []string{"\nSeu dano foi: 10", "Seu dano foi: 8", "Seu dano foi: 6", "Seu dano foi: 4", "Seu dano foi: 2", "Você errou"}

	if d == 5 {
		fmt.Println(i[0])
	} else if d == 4 {
		fmt.Println(i[1])
	} else if d == 3 {
		fmt.Println(i[2])
	} else if d == 2 {
		fmt.Println(i[3])
	} else if d == 1 {
		fmt.Println(i[4])
	} else if d == 0 {
		fmt.Println(i[5])
	}
}
