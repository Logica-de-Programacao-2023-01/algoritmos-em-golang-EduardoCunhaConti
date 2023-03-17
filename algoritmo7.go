package main

import "fmt"

func main() {
	var (
		numero float64
	)

	fmt.Print("Diga um número.")
	fmt.Scan(&numero)

	sucessor := numero + 1
	antecessor := numero - 1

	fmt.Println("O sucessor e antecessor do número são", sucessor, antecessor)

}
