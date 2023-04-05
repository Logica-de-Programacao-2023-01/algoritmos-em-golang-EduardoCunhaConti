package main

import "fmt"

func main() {
	var (
		sexo         int
		altura, peso float64
	)
	fmt.Println("1.Masculino")
	fmt.Println("2.Feminino")
	fmt.Print("Escolha o seu sexo. ")
	fmt.Scan(&sexo)
	fmt.Print("Qual o seu peso? ")
	fmt.Scan(&peso)
	fmt.Print("E qual a sua altura? ")
	fmt.Scan(&altura)

	IMC := peso / (altura * altura)

	fmt.Println("O seu IMC é ", IMC)

	if sexo == 1 {
		if IMC < 20.7 {
			fmt.Println("Você está abaixo do peso.")
		} else if IMC >= 20.7 && IMC <= 26.4 {
			fmt.Println("Você está no peso ideal.")
		} else if IMC > 26.4 && IMC <= 27.8 {
			fmt.Println("Você está um pouco acima do peso.")
		} else if IMC > 27.8 && IMC <= 31.1 {
			fmt.Println("Você está muito acima do peso.")
		} else {
			fmt.Println("Você está com obesidade")
		}
	}

	if sexo == 2 {
		if IMC < 19.1 {
			fmt.Println("Você está abaixo do peso.")
		} else if IMC >= 19.1 && IMC <= 25.8 {
			fmt.Println("Você está no peso ideal.")
		} else if IMC > 25.8 && IMC <= 27.3 {
			fmt.Println("Você está um pouco acima do peso.")
		} else if IMC > 27.3 && IMC <= 32.3 {
			fmt.Println("Você está muito acima do peso.")
		} else {
			fmt.Println("Você está com obesidade")
		}
	}
}
