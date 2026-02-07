package main

import "fmt"

func healPokemon(name string) {
	fmt.Printf("Nurse Joy: Welcome! Placing %s in the healing machine.\n", name)
	defer fmt.Printf("Nurse Joy: The healing machine is now free. \n")

	fmt.Printf("Healing %s...\n", name)
	fmt.Printf("%s is fully healed!\n", name)
}

func healDangerous(name string) {
	fmt.Printf("Nurse Joy: Let me look at %s...\n", name)
	defer fmt.Println("Nurse Joy: Cleaning up the healing machine.")

	if name == "MissingNo" {
		panic("GLITCH POKEMON DETECTED - SYSTEM CRASH")
	}

	fmt.Printf("%s is fully healed!\n", name)
}

func safeHeal(name string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Nurse Joy: Something went wrong with %s! Error: %v\n", name, r)
			fmt.Printf("Nurse Joy: System rebooted. We're back online.\n")
		}
	}()

	healDangerous(name)
}
