package main

import "fmt"

func main() {
	var x int
	var maior int

	fmt.Print("Diga vários números (0 para parar). ")

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		if x > maior {
			maior = x
		}
	}

	if maior == 0 {
		fmt.Println("Nenhum número foi inserido.")
	} else {
		fmt.Printf("O maior número inserido foi %d\n", maior)
	}
}
