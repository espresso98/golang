package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var slice []int = make([]int, 3) // make([]int, 0, 3)
	var in string
	fmt.Print("Please enter an integer (X to quit): ")
	for {
		fmt.Scan(&in)
		if in == "X" {
			break
		}
		el, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println("Wrong input. Please enter an integer!")
			continue
		}
		slice = append(slice, el)
		sort.Ints(slice[:])
		fmt.Println(slice)
		fmt.Print("Please again enter an integer (X to quit): ")
	}
}
