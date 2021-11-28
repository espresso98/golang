package main

import "fmt"

func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
	averageGrade := (midtermGrade + finalGrade) / 2
	var gradeLetter string

	if averageGrade > 90 {
		gradeLetter = "A"
	} else if averageGrade > 80 {
		gradeLetter = "B"
	} else if averageGrade > 70 {
		gradeLetter = "C"
	} else if averageGrade > 60 {
		gradeLetter = "D"
	} else {
		gradeLetter = "F"
	}

	return gradeLetter, averageGrade
}

func main() {
	var myMidterm, myFinal float32
	myMidterm = 89.4
	myFinal = 74.9
	var myAverage float32
	var myGrade string

	myGrade, myAverage = GPA(myMidterm, myFinal)
	fmt.Println(myAverage, myGrade) // Prints 82.12 B
}
