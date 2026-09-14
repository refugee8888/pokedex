package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/refugee8888/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	registry func() map[string]cliCommand
	next     *string
	previous *string
	cache    *pokecache.Cache
}

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func main() {
	cfg := &config{}
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
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}
