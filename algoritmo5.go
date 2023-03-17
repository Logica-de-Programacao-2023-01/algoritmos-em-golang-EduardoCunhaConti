package main

import "fmt"

func main() {
	var (
		idade float64
	)

	fmt.Print("Qual é a sua idade em anos?")
	fmt.Scan(&idade)

	resultado := idade * 365

	fmt.Println("A sua idade em dias é", resultado)

}
