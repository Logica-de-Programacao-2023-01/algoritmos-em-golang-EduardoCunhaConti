package main

import "fmt"

func main() {
	var (
		dias, diaria float64
	)

	fmt.Print("Quantos dias você trabalha?")
	fmt.Scan(&dias)
	fmt.Print("E qual o valor da sua diária?")
	fmt.Scan(&diaria)

	salario := dias * diaria

	fmt.Println("O seu salário é", salario)
}
