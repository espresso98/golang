// Command-line arguments are a common way to parameterize execution of programs.
// For example, go run hello.go uses run and hello.go arguments to the go program
// refefence: https://gobyexample.com/command-line-arguments

package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args        // the first value in this slice is the path to the program
	argsWithoutProg := os.Args[1:] // os.Args[1:] holds the arguments to the program
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

// >go build command-line-args.go
// >./command-line-args a b c d
// [./command-line-args a b c d]
// [a b c d]
// c
