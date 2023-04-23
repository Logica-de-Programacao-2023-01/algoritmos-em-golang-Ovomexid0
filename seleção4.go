package main

import "fmt"

func main() {

	var altura, peso, sexo float64

	fmt.Print("digite sua altura em metros:")
	fmt.Scan(&altura)
	fmt.Print("digite seu peso em kilos:")
	fmt.Scan(&peso)
	fmt.Print("qual é o seu sexo: digite 1 para homem e 2 para mulher: ")
	fmt.Scan(&sexo)

	if sexo == 1 {
		println("você é homem")
	} else if sexo == 2 {
		println("você é mulher")
	} else {
		println("essa opção não é válida")
	}

	resultado := peso / (altura * altura)
	fmt.Println("o seu IMC é", resultado)

	if resultado < 18.5 {
		fmt.Println("seu peso está abaixo do normal")
	}
	if resultado >= 18.5 && resultado < 24.9 {
		fmt.Println("seu peso é normal")
	}
	if resultado >= 25.0 && resultado < 29.9 {
		fmt.Println("excesso de peso")
	}
	if resultado >= 30.0 && resultado < 34.9 {
		fmt.Println("obesidade classe 1")
	}
	if resultado >= 35.0 && resultado < 39.9 {
		fmt.Println("obesidade classe 2")
	}
	if resultado >= 40 {
		fmt.Println("obesidade 3")
	}
}
