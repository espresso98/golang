// Command-line flags are a common way to specify options for command-line programs.
// For example, in wc -l the -l is a command-line flag.
// reference: https://gobyexample.com/command-line-flags

package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag declarations
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	// StringVar defines a string flag with specified name, default value, and usage string.
	// The argument p points to a string variable in which to store the value of the flag.
	// func StringVar(p *string, name string, value string, usage string)
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// execute the command-line parsing
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// >go build command-line-flag.go
// >./command-line-flags -word=opt -numb=7 -fork -svar=flag
// word: opt
// numb: 7
// fork: true
// svar: flag
// tail: []

// If you omit flags they automatically take their default values.
// Trailing positional arguments can be provided after any flags.
// >./command-line-flags -word=opt a1 a2 a3
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: [a1 a2 a3]

// >./command-line-flags -word=opt a1 a2 a3 -numb=7
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: [a1 a2 a3 -numb=7]

// Use -h or --help flags to get automatically generated help text for the command-line program.
// >./command-line-flags -h
// Usage of ./command-line-flags:
//   -fork
//         a bool
//   -numb int
//         an int (default 42)
//   -svar string
//         a string var (default "bar")
//   -word string
//         a string (default "foo")
