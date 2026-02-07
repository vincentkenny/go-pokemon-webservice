package main

import "fmt"

type Battler interface {
	Attack() int
	GetName() string
}

type Pikachu struct {
	Name  string
	Power int
}

type Onix struct {
	Name    string
	Defense int
}

type Gengar struct {
	Name  string
	SpAtk int
	Speed int
}

func (g Gengar) GetName() string {
	return g.Name
}

func (g Gengar) Attack() int {
	return g.SpAtk + g.Speed
}

func (p Pikachu) Attack() int {
	return p.Power * 2
}

func (p Pikachu) GetName() string {
	return p.Name
}

func (o Onix) Attack() int {
	return o.Defense / 2
}

func (o Onix) GetName() string {
	return o.Name
}

func battleCry(b Battler) {
	fmt.Printf("%s enters the arena! Attack power: %d\n", b.GetName(), b.Attack())
}

func announce(b Battler) {
	switch fighter := b.(type) {
	case Pikachu:
		fmt.Printf("%s uses Thunderbolt! Special power: %d\n", fighter.Name, fighter.Power)
	case Onix:
		fmt.Printf("%s uses Rock Trhow! Wall defense: %d\n", fighter.Name, fighter.Defense)
	case Gengar:
		fmt.Printf("%s uses Shadow Ball! SpAtk: %d, Speed: %d\n", fighter.Name, fighter.SpAtk, fighter.Speed)
	default:
		fmt.Printf("Unknown fighter: %s\n", b.GetName())
	}
}
