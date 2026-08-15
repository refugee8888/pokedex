package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

		"map": {
			name:        "map",
			description: "Map of Pokemon",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Map of Pokemon",
			callback:    commandMapb,
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

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
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

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		url = *cfg.next
	}
	//prev_url := "https://pokeapi.co/api/v2/location-area/previous"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var dump LocationArea
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&dump); err != nil {
		return err
	}
	cfg.previous = dump.Previous
	cfg.next = dump.Next
	for _, v := range dump.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.previous != nil {
		url = *cfg.previous
	}
	//prev_url := "https://pokeapi.co/api/v2/location-area/previous"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var dump LocationArea
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&dump); err != nil {
		return err
	}
	cfg.previous = dump.Previous
	for _, v := range dump.Results {
		fmt.Println(v.Name)
	}

	return nil
}
