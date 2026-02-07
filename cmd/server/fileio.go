package main

import (
	"fmt"
	"os"
)

func saveTeam(filename string, team []Pokemon) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, p := range team {
		line := fmt.Sprintf("%s, %s, %d, %d\n", p.Name, p.Type, p.Level, p.HP)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Team saved to %s\n", filename)
	return nil
}

func loadTeam(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	fmt.Println("Loaded team data:")
	fmt.Println(string(data))
	return nil
}
