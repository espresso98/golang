/* 
go run animal.go
> Please enter the animal name and info(type "exit" to exit): snake eat
I eat mice.
> Please enter the animal name and info(type "exit" to exit): snake speak
I hiss.
> Please enter the animal name and info(type "exit" to exit): snake move
I slither.
> Please enter the animal name and info(type "exit" to exit): exit
*/

package main

import (
	"fmt"
	"os"
)

//Animal definition
type Animal struct {
	food, locomotion, noise string
}

func (animal Animal) Eat() {
	fmt.Println("I eat", animal.food+".")
}

func (animal Animal) Move() {
	fmt.Println("I", animal.locomotion+".")
}

func (animal Animal) Speak() {
	fmt.Println("I", animal.noise+".")
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	snake := Animal{"mice", "slither", "hiss"}
	bird := Animal{"worms", "fly", "peep"}

	var name, info string
	var animal Animal 

	fmt.Printf("\nThere are three animals: cow, snake, and bird. They can do eat, move, or speak.\n")

	for {
		fmt.Printf("> Please enter the animal name and info(type \"exit\" to exit): ")
		fmt.Scanln(&name, &info)
		if name == "exit" {
			os.Exit(0)
		}

		switch name {
			case "cow":
				animal = cow
			case "snake":
				animal = snake
			case "bird":
				animal = bird
			default:
				fmt.Println("Please try again. Exit to type \"exit\"")
				continue
		}

		switch info {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Please try again. Exit to type \"exit\"")
				continue
		}
	}
}
