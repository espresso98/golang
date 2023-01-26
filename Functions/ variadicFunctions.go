// Variadic functions can be called with any number of trailing arguments.
// Reference: [https://gobyexample.com/variadic-functions](https://gobyexample.com/variadic-functions)
package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

// $ go run variadic-functions.go
// [1 2] 3
// [1 2 3] 6
// [1 2 3 4] 10
