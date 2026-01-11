package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	// typeChart := map[string]map[string]float64{
	// 	"fire": {
	// 		"grass": 2.0,
	// 		"water": 0.5,
	// 		"fire":  0.5,
	// 	},
	// 	"water": {
	// 		"fire":  2.0,
	// 		"grass": 0.5,
	// 		"water": 0.5,
	// 	},
	// 	"grass": {
	// 		"water": 2.0,
	// 		"fire":  0.5,
	// 		"grass": 0.5,
	// 	},
	// }

	// attackType := "grass"
	// defenseType := "electric"

	// if multipliers, ok := typeChart[attackType]; ok {
	// 	if mult, ok := multipliers[defenseType]; ok {
	// 		fmt.Printf("%s vs %s: %.1fx damage\n", attackType, defenseType, mult)
	// 	} else {
	// 		fmt.Printf("%s vs %s: 1.0x damage (neutral)\n", attackType, defenseType)
	// 	}
	// }

	// team := []Pokemon{
	// 	{
	// 		Name:    "Torchic",
	// 		Type:    "fire",
	// 		Level:   5,
	// 		HP:      45,
	// 		Attack:  15.5,
	// 		Defense: 8.0,
	// 	},
	// 	{
	// 		Name:    "Mudkip",
	// 		Type:    "water",
	// 		Level:   5,
	// 		HP:      50,
	// 		Attack:  12,
	// 		Defense: 10.0,
	// 	},
	// 	{
	// 		Name:    "Treecko",
	// 		Type:    "grass",
	// 		Level:   5,
	// 		HP:      40,
	// 		Attack:  14.0,
	// 		Defense: 7.0,
	// 	},
	// }

	// team = append(team, Pokemon{
	// 	Name:    "Pikachu",
	// 	Type:    "electric",
	// 	Level:   5,
	// 	HP:      45,
	// 	Attack:  12.0,
	// 	Defense: 9.0,
	// })

	// fmt.Println("My team:")
	// for _, pokemon := range team {
	// 	pokemon.PrintInfo()
	// }

	// team[0].TakeDamage(10)

	// team[0].PrintInfo()

	pokemon := ""
	catchRate := 140

	caught, err := catchPokemon(pokemon, catchRate)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if caught {
		fmt.Printf("Gotcha! %s was caught!\n", pokemon)
	} else {
		fmt.Printf("Oh no! %s broke free!\n", pokemon)
	}
}

type Pokemon struct {
	Name    string
	Type    string
	Level   int
	HP      int
	Attack  float64
	Defense float64
}

func (p Pokemon) PrintInfo() {
	fmt.Printf("%s (Lv.%d) - %s type, HP: %d\n", p.Name, p.Level, p.Type, p.HP)
}

func (p *Pokemon) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Printf("%s took %d damage! HP: %d\n", p.Name, amount, p.HP)
}

func catchPokemon(pokemonName string, catchRate int) (bool, error) {
	if strings.TrimSpace(pokemonName) == "" {
		return false, errors.New("pokemon name cannot be empty")
	}

	if catchRate < 0 || catchRate > 100 {
		return false, errors.New("catch rate must be between 0 and 100")
	}

	roll := rand.Intn(100)
	caught := roll < catchRate
	return caught, nil
}

func getTypeMultiplier(attackType string, defenseType string) float64 {
	if attackType == "fire" && defenseType == "grass" {
		return 2.0 // Super effective!
	} else if attackType == "fire" && defenseType == "water" {
		return 0.5 // Not very effective...
	} else if attackType == "fire" && defenseType == "fire" {
		return 0.5 // Not very effective...
	} else if attackType == "electric" && defenseType == "water" {
		return 2.0
	} else {
		return 1.0 // Normal damage
	}
}

func calculateDamage(attack float64, defense float64, typeMultiplier float64) float64 {
	return (attack - (defense / 2)) * typeMultiplier
}
