package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Create a map
	m := make(map[string]string)
	var name string
	var address string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your address: ")
	fmt.Scan(&address)

	// Add values to map
	m["name"] = name
	m["address"] = address

	// Create JSON object from the map
	json, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(json))
}
