package main

import "fmt"

func main() {
	soma := 0
	quantidade := 0

	for {
		var x int
		fmt.Print("Diga um número (0 para parar). ")
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		soma += x
		quantidade++
	}

	if quantidade > 0 {
		media := float64(soma) / float64(quantidade)
		fmt.Printf("A média é: %.0f\n", media)
	} else {
		fmt.Println("Nenhum número foi inserido.")
	}
}
