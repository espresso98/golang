package main

import "fmt"

func computeMarsYears(earthYears int) int {
	// Remove earthYears definition within marsYear
	earthDays := earthYears * 365
	marsYears := earthDays / 687
	return marsYears
}

func main() {
	myAge := 25

	myMartianAge := computeMarsYears(myAge)
	fmt.Println(myMartianAge) // Prints 13
}
