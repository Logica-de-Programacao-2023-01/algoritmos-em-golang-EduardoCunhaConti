package main

import "fmt"

func main() {
	var (
		dia int
	)
	fmt.Print("Diga um número. ")
	fmt.Scan(&dia)

	if dia == 1 {
		fmt.Println("O dia é domingo.")
	} else if dia == 2 {
		fmt.Println("O dia é segunda-feira.")
	} else if dia == 3 {
		fmt.Println("O dia é terça-feira.")
	} else if dia == 4 {
		fmt.Println("O dia é quarta-feira.")
	} else if dia == 5 {
		fmt.Println("O dia é quinta-feira.")
	} else if dia == 6 {
		fmt.Println("O dia é sexta-feira.")
	} else if dia == 7 {
		fmt.Println("O dia é sábado.")
	} else {
		fmt.Println("O número não corresponde à um dia da semana.")
	}
}
