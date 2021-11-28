package main

import "fmt"

func trunc() {
	var x float32
	fmt.Print("Enter a float number: ")
	fmt.Scan(&x)
	fmt.Println("The truncated number you just entered is ", int(x))
}
func main() {
	trunc()
}
