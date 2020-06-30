/* EXERCISE
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”,
respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should
print the JSON object.
*/

package main

import (
	"bytes"
	"fmt"
	"encoding/json"

)

func main() {
	var name string
	var address string
	fmt.Println("Person's name: ")
	fmt.Scanln(&name)
	fmt.Println("Person's address: ")
	fmt.Scanln(&address)

	personMap := map[string] string {
		"name": name,
		"address": address,
	}

	jsonPerson, err := json.Marshal(personMap)
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, jsonPerson, "", "\t")

	if err == nil && error == nil{
		fmt.Println( string(prettyJSON.Bytes()))
	} else {
		fmt.Println(err)
	}
}
