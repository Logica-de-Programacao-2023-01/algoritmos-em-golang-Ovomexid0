package main

import "fmt"

func main() {

	var altura, peso float64

	fmt.Print("qual é a sua altura?")
	fmt.Scan(&altura)
	fmt.Print("qual é o seu peso?")
	fmt.Scan(&peso)

	resultado := peso / (altura * altura)
	fmt.Print("o seu IMC é", resultado)
	fmt.Scan(&resultado)

}
