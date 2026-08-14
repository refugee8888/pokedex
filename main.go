package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex>")
		scanner.Scan()
		input := scanner.Text()

		clean := cleanInput(input)
		command, exists := supportedCommands()[clean[0]]
		if !exists {
			fmt.Println("Unknown command")
		}
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}

	}
}
