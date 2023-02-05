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

	fmt.Println("Você está um uma floresta, proseguir?")

	w := ""

	fmt.Scanf("%s\n", &w)

	s := "sim"

	n := "não"

	rand.Seed(time.Now().UnixNano())
	v := (rand.Intn(10))

	if w == s {
		fmt.Printf("Você encontrou um monstro que possui %v de vida, deseja enfrenta-lo?\n", v)
	} else if w == n {
		fmt.Println("Jogo finalizado")
	}

	y := ""

	fmt.Scanf("%s\n", &y)

	if y == s {
		fmt.Println("")

		rand.Seed(time.Now().UnixNano())
		x := (rand.Intn(6))
		d := x

		i := []string{
			"Você tirou 6 no dado\nSeu dano foi: 10\n",
			"Você tirou 5 no dado\nSeu dano foi: 8\n",
			"Você tirou 4 no dado\nSeu dano foi: 6\n",
			"Você tirou 3 no dado\nSeu dano foi: 4\n",
			"Você tirou 2 no dado\nSeu dano foi: 2\n",
			"Você tirou 1 no dado\nVocê errou o ataque\n"}

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

		d = x * 2

		if d >= v {
			fmt.Println("Você o derrotou\n ")
		} else {
			fmt.Println("Você foi derrotado\n ")
		}
	} else if y == n {
		x := (rand.Intn(2))
		if x == 1 {
			fmt.Println("Você morreu ao tentar fugir")
		} else if x == 0 {
			fmt.Println("Você conseguiu fugir")
		}
	}
}
