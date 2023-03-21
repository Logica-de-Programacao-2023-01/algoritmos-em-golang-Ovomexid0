package main

import "fmt"

func main() {

	var idade float64

	fmt.Print("qual é a sua idade?")
	fmt.Scan(&idade)

	dias := 365 * idade
	fmt.Println("sua idade em dias é", dias)

}
