package main

import "fmt"

func main() {
	var (
		x, y float64
	)
	fmt.Print("Qual é o primeiro número? ")
	fmt.Scan(&x)
	fmt.Print("E qual é o segundo número? ")
	fmt.Scan(&y)

	if x > y {
		fmt.Println("O maior número é: ", x)
	} else {
		fmt.Println("O maior número é: ", y)
	}
}
