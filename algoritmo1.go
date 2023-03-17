package main

import "fmt"

func main() {
	var (
		primeiro, segundo, terceiro float64
	)
	fmt.Print("Qual é o primeiro número?")
	fmt.Scan(&primeiro)
	fmt.Print("Qual é o segundo valor?")
	fmt.Scan(&segundo)
	fmt.Print("Qual é o terceiro valor?")
	fmt.Scan(&terceiro)

	soma := primeiro + segundo + terceiro

	fmt.Println("A soma dos números é", soma)

}
