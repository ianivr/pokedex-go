package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		inputCleaned := cleanInput(input)
		firstWord := inputCleaned[0]

		fmt.Println("Your command was:", firstWord)
	}
}
