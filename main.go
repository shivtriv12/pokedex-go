package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			cleantext := cleanInput(scanner.Text())
			fmt.Println("Your command was: " + cleantext[0])
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
