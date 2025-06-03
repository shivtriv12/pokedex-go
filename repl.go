package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Lists all the Command's Usage",
		callback:    commandHelp,
	}
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			cleantext := cleanInput(scanner.Text())
			cf, ok := commands[cleantext[0]]
			if ok {
				cf.callback()
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println(
		`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}
