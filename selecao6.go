package main

import "fmt"

func main() {
	var (
		x, y int
	)
	fmt.Print("Diga um número. ")
	fmt.Scan(&x)
	fmt.Print("Diga outro número. ")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		resultado := x * y
		fmt.Println("O produto dos dois números é", resultado)
	} else {
		resultado := x + y
		fmt.Println("A soma dos dois números é", resultado)
	}
}
