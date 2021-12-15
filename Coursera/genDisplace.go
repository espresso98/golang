// s = ½ a t2 + vo t + so
// three float64
// prompts the user to enter values for acceleration, initial velocity, and initial displacement.
// generate a function fn which will compute displacement as a function of time.
// fn := GenDisplaceFn(10, 2, 1)
// fmt.Println(fn(5))

package main

import "fmt"

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	return func(t float64) float64 {
		displacement := (0.5 * a * (t * t) +(vo * t) + so)
		return displacement
	}
}

func main() {
	var acc, vo, so, t float64

	fmt.Print("Enter acceleration: ")
	fmt.Scan(&acc)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&vo)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&so)

	fn := GenDisplaceFn(acc, vo, so)

	fmt.Print("Enter time in seconds: ")
	fmt.Scan(&t)

	fmt.Printf("\nDisplacement at %f seconds is %f.\n", t, fn(t))
}