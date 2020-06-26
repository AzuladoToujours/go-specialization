package main

import (
	"fmt"
	"strings"
)

func main() {
	var stringToTest string
	fmt.Printf("Enter the string:")
	fmt.Scan(&stringToTest)
	var test= findian(stringToTest)
	fmt.Printf("%s \n", test)
}
func findian (stringReceived string) string {
	var found bool = false;
	var stringLower = strings.ToLower(stringReceived)
	if strings.HasPrefix(stringLower, "i") && strings.HasSuffix(stringLower, "n"){
			if strings.Contains(stringLower, "a"){
				found = true;
			}
	}
	if found == true {
		return "Found!"
	} else {
		return "Not Found!"
	}
}
