package main

import "fmt"

func main() {
	var (
		primeiro, segundo, terceiro float64
	)

	fmt.Print("Qual é o primeiro número?")
	fmt.Scan(&primeiro)
	fmt.Print("Qual é o segundo número?")
	fmt.Scan(&segundo)
	fmt.Print("Qual é o terceiro número?")
	fmt.Scan(&terceiro)

	media := (primeiro*2 + segundo*3 + terceiro*5) / 10

	fmt.Println("A média ponderada é", media)

}
