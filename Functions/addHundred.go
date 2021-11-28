package main

import "fmt"

func addHundred(num *int) {
	*num += 100
}

func main() {
	x := 1

	addHundred(&x)

	fmt.Println(x) // Print 101
}
