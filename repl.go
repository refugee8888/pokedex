package main

import (
	"fmt"
	"os"
	"strings"
)

func supportedCommands() map[string]cliCommand {
	inputCommand := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Instructions on how to use Pokedex commands",
			callback:    commandHelp,
		},
	}
	return inputCommand
}

func cleanInput(text string) []string {
	slice := make([]string, 0)
	lowerStr := strings.ToLower(text)
	whiteSpace := " "
	trimmedStr := strings.Trim(lowerStr, whiteSpace)

	splitStr := strings.Split(trimmedStr, " ")
	for _, word := range splitStr {
		slice = append(slice, word)
	}

	return slice
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return nil
}

func commandHelp() error {
	commands := make([]string, 0)
	fmt.Println("Welcome to the Pokedex!")

	for _, v := range supportedCommands() {
		commands = append(commands, v.name)
	}

	if commands != nil {
		fmt.Println("Usage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	}

	return nil
}
