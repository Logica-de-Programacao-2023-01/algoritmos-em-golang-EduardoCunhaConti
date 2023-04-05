package main

import "fmt"

func main() {
	var (
		x int
	)
	fmt.Print("Diga um número. ")
	fmt.Scan(&x)

	if x%3 == 0 && x%5 == 0 {
		fmt.Println("O número é múltiplo de 3 e 5.")
	} else {
		fmt.Println("O número não é múltiplo de 3 e 5 ao mesmo tempo.")
	}
}
