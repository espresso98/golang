package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // recieve-only type chan int
	cs := make(chan<- int) // send-only type chann int

	fmt.Println("------------------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// general to specific converts
	fmt.Println("------------------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))
}

// ------------------
// c       chan int
// cr      <-chan int
// cs      chan<- int
// ------------------
// c       <-chan int
// c       chan<- int
