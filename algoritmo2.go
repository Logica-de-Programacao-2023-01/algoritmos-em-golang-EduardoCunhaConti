package main

import "fmt"

func main() {
	var (
		numero float64
	)

	fmt.Print("Qual é o número?")
	fmt.Scan(&numero)

	dobro := numero * 2
	triplo := numero * 3
	quadruplo := numero * 4

	fmt.Println("O dobro é", dobro)
	fmt.Println("O triplo é", triplo)
	fmt.Println("O quadruplo é", quadruplo)

}
