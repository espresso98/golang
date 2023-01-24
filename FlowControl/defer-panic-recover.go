package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls
	defer cleanup("first")
	defer cleanup("second")

	// defer recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic:", err)
		}
	}()

	fmt.Println("running in main...")

	// panic
	panic("something bad happened")
}

// $ go run main.go
// running in main...
// recovered from panic: something bad happened
// second
// first
