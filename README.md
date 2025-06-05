# 🎮 Pokédex CLI

A command-line Pokédex application built in Go that allows you to explore the Pokémon world, catch Pokémon, and manage your collection using the [PokéAPI](https://pokeapi.co/).

## 🚀 Getting Started

### Prerequisites

- Go 1.19 or higher installed on your system
- Internet connection (to fetch data from PokéAPI)

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/pokedex-go.git
   cd pokedex-go
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run .
   ```

## 🎯 How to Play

Once you start the application, you'll see the Pokédex prompt:

```
Pokedex >
```

### Available Commands

| Command   | Usage                    | Description                             |
| --------- | ------------------------ | --------------------------------------- |
| `help`    | `help`                   | Shows all available commands            |
| `exit`    | `exit`                   | Exit the Pokédex                        |
| `map`     | `map`                    | Browse to next 20 location areas        |
| `mapb`    | `mapb`                   | Browse to previous 20 location areas    |
| `explore` | `explore <area-name>`    | Explore a specific area to find Pokémon |
| `catch`   | `catch <pokemon-name>`   | Attempt to catch a Pokémon              |
| `inspect` | `inspect <pokemon-name>` | View details of a caught Pokémon        |
| `pokedex` | `pokedex`                | List all your caught Pokémon            |

### 🎮 Example Gameplay

```bash
Pokedex > map
canalave-city-area
pastoria-city-area
sunyshore-city-area
...

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - magikarp
 - gyarados

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!
You may now inspect it with the inspect command.

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  -hp: 20
  -attack: 10
  -defense: 55
  -special-attack: 15
  -special-defense: 20
  -speed: 80
Types:
  - water

Pokedex > pokedex
Your Pokedex:
 - magikarp
```

## 📂 Project Structure

```
pokedex-go/
├── main.go                 # Application entry point
├── repl.go                 # Read-Eval-Print Loop and CLI logic
├── commands_callback.go    # Command implementations
├── internal/
│   └── caching.go         # Cache implementation
├── repl_test.go           # Unit tests
├── go.mod                 # Go module dependencies
└── README.md              # This file
```

## 🧪 Running Tests

Execute the test suite to ensure everything works correctly:

```bash
go test ./...
```

The tests cover:

- ✅ Input cleaning and validation
- ✅ Cache add/get operations
- ✅ Cache cleanup (reap loop) functionality
