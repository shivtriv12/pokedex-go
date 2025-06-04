package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	pokecache "github.com/shivtriv12/pokedex-go/internal"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(c *config, ch *pokecache.Cache) error
}

type config struct {
	Next     string
	Previous string
}

type locationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandExit(c *config, ch *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, ch *pokecache.Cache) error {
	fmt.Println(
		`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func commandMap(c *config, ch *pokecache.Cache) error {
	cache, ok := ch.Get(c.Next)
	if !ok {
		res, err := http.Get(c.Next)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		cache = body
		ch.Add(c.Next, cache)
	}
	locations := locationArea{}
	err1 := json.Unmarshal(cache, &locations)
	if err1 != nil {
		log.Fatal(err1)
	}
	for _, name := range locations.Results {
		fmt.Println(name.Name)
	}
	if locations.Next != nil && *locations.Next != "" {
		c.Next = *locations.Next
	}
	if locations.Previous != nil && *locations.Previous != "" {
		c.Previous = *locations.Previous
	}
	return nil
}

func commandMapb(c *config, ch *pokecache.Cache) error {
	if c.Previous == "" {
		fmt.Println("No Backwards")
		return nil
	}
	cache, ok := ch.Get(c.Previous)
	if !ok {
		res, err := http.Get(c.Previous)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		cache = body
		ch.Add(c.Previous, cache)
	}
	locations := locationArea{}
	err1 := json.Unmarshal(cache, &locations)
	if err1 != nil {
		log.Fatal(err1)
	}
	for _, name := range locations.Results {
		fmt.Println(name.Name)
	}
	if locations.Next != nil && *locations.Next != "" {
		c.Next = *locations.Next
	}
	if locations.Previous != nil && *locations.Previous != "" {
		c.Previous = *locations.Previous
	}
	return nil
}
