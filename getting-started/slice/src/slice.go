/* EXERCISE:
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop.
Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import "fmt"

func main() {
	var sli = make([]int, 0, 3)
	var input int
	for ok := true; ok; ok = input != 'X' {
		fmt.Println("Enter integer, press X for exit")
		n, err := fmt.Scanln(&input)
		if n < 1 || err != nil {
			fmt.Println("expecting integer, error.")
			break
		}
		sli = append(sli, input)
		bubble(sli)
		fmt.Println("Sorted Slice: ", sli)
	}
	fmt.Println("Sorted Slice: ", sli)
}

func bubble(sli []int) {
	size := len(sli)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if sli[j] < sli[j-1] {
				sli[j], sli[j-1] = sli[j-1], sli[j]
			}
		}
	}
}
