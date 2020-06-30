/* EXERCISE:
Write a program which prompts the user to enter a floating point number and prints the integer
which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.*/

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
