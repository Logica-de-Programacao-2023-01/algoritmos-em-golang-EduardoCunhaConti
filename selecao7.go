package main

import "fmt"

func main() {
	var (
		salario float64
	)
	fmt.Print("Qual o seu salário? ")
	fmt.Scan(&salario)

	if salario < 1000 {
		novo := salario * 1.1
		fmt.Println("O novo salário será", novo)
	} else {
		novo := salario * 1.05
		fmt.Println("O novo salário será", novo)
	}

}
