package main

import "fmt"

func main() {
	var (
		salario float64
	)

	fmt.Print("Qual é o seu salário?")
	fmt.Scan(&salario)

	aumento := (salario * 115) / 100

	fmt.Println("O seu novo salário será", aumento)

}
