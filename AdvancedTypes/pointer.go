package main

import "fmt"

func main() {
	var a int = 500
	var p *int

	p = &a
	fmt.Printf("The value of p: %p\n", p)
	fmt.Printf("The value of *p: %d\n", *p)

	*p = 100
	fmt.Printf("The value of a: %d\n", a)
}

// The value of p: 0xc000016088
// The value of *p: 500
// The value of a: 100
