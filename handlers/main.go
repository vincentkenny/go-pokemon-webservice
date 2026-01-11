package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Score struct {
	Points int
}

func (p Person) Greet() {
	fmt.Println("Hi, I'm", p.Name, "and I'm", p.Age, "years old.")
}

func (p *Person) updateAge(newAge int) {
	p.Age = newAge
}

func PrintMessage(message string) {
	fmt.Println("Printing:", message)
}

func PrintAnything(msg string) {
	fmt.Println(">>", msg)
}

func (s Score) AddOneCopy() {
	s.Points++
}

func (s *Score) AddOne() {
	s.Points++
}

// func main() {
// 	msg := "hello go"
// 	fmt.Println(ToUpperCase(msg))
// }
