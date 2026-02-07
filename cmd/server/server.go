package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	team := []Pokemon{
		{Name: "Torchic", Type: "fire", Level: 5, HP: 45},
		{Name: "Mudkip", Type: "water", Level: 5, HP: 50},
		{Name: "Treecko", Type: "grass", Level: 5, HP: 40},
	}

	err := saveTeam("team.txt", team)
	if err != nil {
		fmt.Println("Save error:", err)
		return
	}

	err = loadTeam("team.txt")
	if err != nil {
		fmt.Println("Load error:", err)
		return
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
