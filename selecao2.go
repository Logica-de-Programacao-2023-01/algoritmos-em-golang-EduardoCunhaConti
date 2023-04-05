package main

import "fmt"

func main() {
	var (
		x, y, z float64
	)
	fmt.Print("Diga o primeiro número. ")
	fmt.Scan(&x)
	fmt.Print("Diga o segundo número. ")
	fmt.Scan(&y)
	fmt.Print("Diga o terceiro número. ")
	fmt.Scan(&z)

	if x > z && y > z {
		fmt.Println("O menor número é ", z)
	} else if x > y && z > y {
		fmt.Println("O menor número é ", y)
	} else if y > x && z > x {
		fmt.Println("O menor número é ", x)
	}
}
