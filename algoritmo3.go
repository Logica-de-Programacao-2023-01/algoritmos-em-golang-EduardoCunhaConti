package main

import "fmt"

func main() {
	var (
		peso, altura float64
	)

	fmt.Print("Qual é o peso?")
	fmt.Scan(&peso)
	fmt.Print("E qual é a altura?")
	fmt.Scan(&altura)

	imc := peso / (altura * altura)

	fmt.Println("O IMC é", imc)

}
