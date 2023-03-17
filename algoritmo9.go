package main

import "fmt"

func main() {
	var (
		valor float64
	)

	fmt.Print("Qual o valor do produto?")
	fmt.Scan(&valor)

	desconto := (valor * 90) / 100

	fmt.Println("O seu produto custará", desconto)

}
