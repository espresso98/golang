// Implementation of Bubble Sort Algorithm

package main

import "fmt"

func Swap(v []int, i int) []int {
	v[i], v[i+1] = v[i+1], v[i]
	return v
}

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
	return arr
}

func main() {
	fmt.Printf("Enter 10 integers value: ")
	arr := make([]int, 10) // array:= []int{11, 14, 3, 8, 18, 17, 43, ...};
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i]) // 3 8 14 18 17 43 22 15 11 13
	}
	fmt.Println(BubbleSort(arr))
}
