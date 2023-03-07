package main

import "fmt"

func main() {
	// A sender can close a channel to indicate that no more values will be sent.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// elem receives values from the channel repeatedly until it is closed.
	for elem := range queue {
		fmt.Println(elem)
	}
}

// one
// two
