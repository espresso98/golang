package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func first20Char(s string) string {
	runes := []rune(s)
	return string(runes[0:20])
}

func main() {
	var fileName string

	nameSlice := make([]Name, 0)

	fmt.Print("Enter file name: ")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer file.Close()

	var nameObj Name

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if len(s[0]) > 20 {
			s[0] = first20Char(s[0])
		}
		if len(s[1]) > 20 {
			s[1] = first20Char(s[1])
		}

		nameObj.fname, nameObj.lname = s[0], s[1]
		nameSlice = append(nameSlice, nameObj)
	}

	for _, p := range nameSlice {
		fmt.Println(p.fname, p.lname)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}
}
