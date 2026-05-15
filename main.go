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

		cmd, exists := commands[inputCleaned[0]]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}
