package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	pokecache "github.com/shivtriv12/pokedex-go/internal"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(c *config, ch *pokecache.Cache, loc string, pokedex map[string]Pokemon) error
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

type pokemonEncounter struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExit(c *config, ch *pokecache.Cache, loc string, pokedex map[string]Pokemon) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, ch *pokecache.Cache, loc string, pokedex map[string]Pokemon) error {
	fmt.Println(
		`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func commandMap(c *config, ch *pokecache.Cache, loc string, pokedex map[string]Pokemon) error {
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

func commandMapb(c *config, ch *pokecache.Cache, loc string, pokedex map[string]Pokemon) error {
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

func commandExplore(c *config, ch *pokecache.Cache, loc string, pokedex map[string]Pokemon) error {
	cache, ok := ch.Get(loc)
	if !ok {
		res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + loc)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		if res.StatusCode == 404 {
			fmt.Println("No such location found")
			return nil
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		cache = body
		ch.Add(loc, cache)
	}
	pokemons := pokemonEncounter{}
	err1 := json.Unmarshal(cache, &pokemons)
	if err1 != nil {
		log.Fatal(err1)
	}
	if len(pokemons.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon found in this area.")
		return nil
	}
	fmt.Printf("Exploring %s...\n", loc)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Println(" -" + pokemon.Pokemon.Name)
	}
	return nil
}

type Pokemon struct {
	Name     string `json:"name"`
	Height   int    `json:"height"`
	Weight   int    `json:"weight"`
	Base_Exp int    `json:"base_experience"`
	Stats    []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func commandCatch(c *config, ch *pokecache.Cache, pokemon string, pokedex map[string]Pokemon) error {
	fmt.Println("Throwing a Pokeball at " + pokemon + "...")
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokemon)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if res.StatusCode == 404 {
		fmt.Println("Not Valid Pokemon")
		return nil
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var poke Pokemon
	err1 := json.Unmarshal(body, &poke)
	if err1 != nil {
		log.Fatal(err1)
	}
	if poke.Base_Exp <= 0 {
		poke.Base_Exp = 1
	}
	catchChance := rand.Intn(poke.Base_Exp)
	if catchChance > 40 {
		fmt.Printf("%s was caught!\n", pokemon)
		pokedex[pokemon] = poke
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}
	return nil
}

func commandInspect(c *config, ch *pokecache.Cache, pokemon string, pokedex map[string]Pokemon) error {
	val, ok := pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", val.Name)
	fmt.Printf("Height: %d\n", val.Height)
	fmt.Printf("Weight: %d\n", val.Weight)

	fmt.Println("Stats:")
	for _, stat := range val.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemonType := range val.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}

	return nil
}
