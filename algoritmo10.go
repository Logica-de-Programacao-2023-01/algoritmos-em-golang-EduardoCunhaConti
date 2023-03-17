package main

import "fmt"

func main() {
	var (
		quilos float64
	)

	fmt.Print("Qual é o seu peso em quilos?")
	fmt.Scan(&quilos)

	libras := (quilos * 22046) / 10000

	fmt.Println("O seu peso em libras será", libras)
}
