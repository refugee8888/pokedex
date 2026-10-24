package main

import (
	"encoding/json"
	"fmt"
	"io"
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
			name:        "mapb",
			description: "Map of Pokemon",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the map location-area",
			callback:    commandExplore,
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

func commandExit(config *config, param *Argument) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return nil
}

func commandHelp(config *config, param *Argument) error {
	commands := make([]string, 0)
	fmt.Println("Welcome to the Pokedex!")

	for _, v := range supportedCommands() {
		commands = append(commands, v.name)
	}

	if commands != nil {
		fmt.Println("Usage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\nmap: Displays a maps page\nmapb: Goes back to previous map page\nexplore: Used with location argument as 'explore <argument>'. Reveals pokemon in the specified area")
	}

	return nil
}

func commandMap(cfg *config, param *Argument) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		url = *cfg.next
	}
	if cachedEntry, ok := cfg.cache.Get(url); ok {
		var dumpEntry LocationArea
		err := json.Unmarshal(cachedEntry, &dumpEntry)
		if err != nil {
			return err
		}
		for _, v := range dumpEntry.Results {
			fmt.Println(v.Name)
		}
		cfg.previous = dumpEntry.Previous
		cfg.next = dumpEntry.Next
	} else {

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
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(url, body)
		if err := json.Unmarshal(body, &dump); err != nil {
			return err
		}

		cfg.previous = dump.Previous
		cfg.next = dump.Next

		for _, v := range dump.Results {
			fmt.Println(v.Name)
		}
	}

	return nil
}

func commandMapb(cfg *config, param *Argument) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.previous != nil {
		url = *cfg.previous
	}
	if cachedEntry, ok := cfg.cache.Get(url); ok {

		var dumpEntry LocationArea
		err := json.Unmarshal(cachedEntry, &dumpEntry)
		if err != nil {
			return err
		}
		for _, v := range dumpEntry.Results {
			fmt.Println(v.Name)
		}
		cfg.previous = dumpEntry.Previous
		cfg.next = dumpEntry.Next
	} else {

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
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(url, body)
		if err := json.Unmarshal(body, &dump); err != nil {
			return err
		}
		for _, v := range dump.Results {
			fmt.Println(v.Name)
		}
		cfg.previous = dump.Previous
		cfg.next = dump.Next

	}

	return nil
}

func commandExplore(cfg *config, name *Argument) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if name != nil {
		url = url + name.parameter
	}
	fmt.Printf("Exploring %s", name.parameter)

	if cachedEntry, ok := cfg.cache.Get(url); ok {
		var dumpEntry Pokemons
		err := json.Unmarshal(cachedEntry, &dumpEntry)
		if err != nil {
			return err
		}
		for _, v := range dumpEntry.Encounters {
			fmt.Println(v.Pokemon.Name)
		}
	} else {

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

		var dump Pokemons
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(url, body)
		if err := json.Unmarshal(body, &dump); err != nil {
			return err
		}

		for _, v := range dump.Encounters {
			fmt.Println(v.Pokemon.Name)
		}
	}

	return nil
}
