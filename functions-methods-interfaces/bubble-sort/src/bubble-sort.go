/* EXERCISE:
Write a Bubble Sort program in Go.
The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an
argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.
*/

package main

import "fmt"

func main() {
	var sli = make([]int, 0, 3)
	var input int
	for i := 0; i < 10; i++  {
		fmt.Printf("Integer number %d\n", i+1)
		n, err := fmt.Scanln(&input)
		if n < 1 || err != nil {
			fmt.Println("expecting integer, error.")
			break
		}
		sli = append(sli, input)

	}

	bubbleSort(sli)
	fmt.Println("Sorted Slice: ", sli)

}

func bubbleSort(sli []int) {
	size := len(sli)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if sli[j] < sli[j-1] {
				swap(sli, j)
			}
		}
	}
}

func swap(sli []int, index int){
	sli[index], sli[index-1] = sli[index-1], sli[index]
}
