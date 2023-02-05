package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	histograma := make([]int, 10, 10)

	// o for vai repetir 1000 vezes
	for i := 0; i < 1000; i++ {

		x := rand.Intn(6)

		histograma[x]++
	}

	for i := 0; i < 10; i++ {

		fmt.Printf("Histograma[%v]: %v\n", i, histograma[i])
	}

}
