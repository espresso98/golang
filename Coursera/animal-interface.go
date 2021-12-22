/*
go run animal-interface.go

Enter your request with newanimal or query or quit.
> newanimal momoi bird
Created it!
Enter your request with newanimal or query or quit.
> query momoi eat
momoi eats worms
Enter your request with newanimal or query or quit.
> query momoi speak
momoi peep
Enter your request with newanimal or query or quit.
> query momoi move
momoi fly
Enter your request with newanimal or query or quit.
> quit
*/

package main

import (
	"fmt"
	"os"
)

type animal interface {
	Name() string
	Eat()
	Move()
	Speak()
}

type cow struct {
	name, food, locomotion, noise string
}

type bird struct {
	name, food, locomotion, noise string
}

type snake struct {
	name, food, locomotion, noise string
}

func main() {
	animals := make(map[string]animal, 0)

	for {
		var cmd, name, req string
		fmt.Println("Enter your request with newanimal or query or quit.")
		fmt.Print("> ")
		fmt.Scan(&cmd)

		switch cmd {

		case "newanimal":
			fmt.Scan(&name)
			fmt.Scan(&req)
			switch req {
			case "cow":
				animals[name] = cow{name, "grass", "walk", "moo"}
			case "bird":
				animals[name] = bird{name, "worms", "fly", "peep"}
			case "snake":
				animals[name] = snake{name, "mice", "slither", "hsss"}
			default:
				fmt.Println("Wrong animal type. Please try again with cow/bird/snake")
				continue
			}
			fmt.Println("Created it!")
		case "query":
			fmt.Scan(&name)
			fmt.Scan(&req)
			a := animals[name]
			if a == nil {
				fmt.Println("Animal with given name does not exist. Please try again.")
				continue
			}
			switch req {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				fmt.Println("Wrong animal info queried, please try again with eat/move/speak.")
				continue
			}
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Wrong command, Please try again with newanimal/query/quit.")
		}

	}
}

func (a cow) Name() string {
	return a.name
}

func (a bird) Name() string {
	return a.name
}

func (a snake) Name() string {
	return a.name
}

func (a cow) Eat() {
	fmt.Println(a.name, "eats", a.food)
}

func (a bird) Eat() {
	fmt.Println(a.name, "eats", a.food)
}

func (a snake) Eat() {
	fmt.Println(a.name, "eats", a.food)
}

func (a cow) Move() {
	fmt.Println(a.name, a.locomotion)
}

func (a bird) Move() {
	fmt.Println(a.name, a.locomotion)
}

func (a snake) Move() {
	fmt.Println(a.name, a.locomotion)
}

func (a cow) Speak() {
	fmt.Println(a.name, a.noise)
}

func (a bird) Speak() {
	fmt.Println(a.name, a.noise)
}

func (a snake) Speak() {
	fmt.Println(a.name, a.noise)
}
