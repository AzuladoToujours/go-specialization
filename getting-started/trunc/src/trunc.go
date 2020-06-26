package main

import "fmt"

func main() {
	var floatNumber float32
	fmt.Printf("Enter the float number:")
	fmt.Scan(&floatNumber)
	var intNumber = int32(floatNumber)
	fmt.Printf("\n Float number: %f", floatNumber)
	fmt.Printf("\n Integer number: %d", intNumber)
	fmt.Printf("\n Bye bye \n")
}
