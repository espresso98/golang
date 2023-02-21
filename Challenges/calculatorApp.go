package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	val1 := getInput("value 1")
	val2 := getInput("value 2")

	var result float64

	switch operation := getOperation(); operation {
	case "+":
		result = addValues(val1, val2)
	case "-":
		result = subtractValues(val1, val2)
	case "*":
		result = multiplyValues(val1, val2)
	case "/":
		result = divideValues(val1, val2)
	default:
		panic("Invalid operation")
	}
	result = math.Round(result*100) / 100
	fmt.Printf("The result is %v\n", result)
}

func getInput(prompt string) float64 {
	fmt.Printf("%v: ", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}
	return value
}

func getOperation() string {
	fmt.Print("Select an operation (+ - * /): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func addValues(val1, val2 float64) float64 {
	return val1 + val2
}

func subtractValues(val1, val2 float64) float64 {
	return val1 - val2
}

func multiplyValues(val1, val2 float64) float64 {
	return val1 * val2
}

func divideValues(val1, val2 float64) float64 {
	return val1 / val2
}

// go run Challenges/calculatorApp.go
// value 1: 34
// value 2: 56.1
// Select an operation (+ - * /): +
// The result is 90.1
