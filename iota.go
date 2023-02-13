// Iota in Go is used to represent constant increasing sequences.
package main

import (
	"fmt"
)

const (
	_  = 1 << (iota * 10) // ignore the first value
	KB                    // decimal:       1024 -> binary 00000000000000000000010000000000
	MB                    // decimal:    1048576 -> binary 00000000000100000000000000000000
	GB                    // decimal: 1073741824 -> binary 01000000000000000000000000000000
)

func main() {
	fmt.Println("KB =", KB)
	fmt.Println("MB =", MB)
	fmt.Println("GB =", GB)
}

// KB = 1024
// MB = 1048576
// GB = 1073741824

// const (
// 	January Month = 1 + iota
// 	February
// 	March
// 	April
// 	May
// 	June
// 	July
// 	August
// 	September
// 	October
// 	November
// 	December
// )

// const (
// 	Adventure Genre = iota + 1
// 	Comic
// 	Crime
// 	Fiction
// 	Fantasy
// 	Historical
// 	Horror
// 	Magic
// 	Mystery
// 	Philosophical
// 	Political
// 	Romance
// 	Science
// 	Superhero
// 	Thriller
// 	Western
// )
