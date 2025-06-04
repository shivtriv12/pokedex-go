package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokecache "github.com/shivtriv12/pokedex-go/internal"
)

func startRepl() {
	baseConfig := config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
	}

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Lists all the Command's Usage",
		Callback:    commandHelp,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the name of next 20 location areas in pokemon world",
		Callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the name of prev 20 location areas in pokemon world",
		Callback:    commandMapb,
	}
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Minute)
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			cleantext := cleanInput(scanner.Text())
			if len(cleantext) == 0 {
				continue
			}
			cf, ok := commands[cleantext[0]]
			if ok {
				cf.Callback(&baseConfig, cache)
			} else {
				fmt.Println("Unknown command")
			}
			break
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	splitText := strings.Fields(text)
	return splitText
}
