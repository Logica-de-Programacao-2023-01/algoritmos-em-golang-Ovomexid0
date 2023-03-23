package main

import "fmt"

func main() {

	var quilos float64

	fmt.Print("quanto você pesa?")
	fmt.Scan(&quilos)

	ibs := quilos * 2.2046
	fmt.Println("seu valor em libras é", ibs)

}
