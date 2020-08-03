package main

import (
	"fmt"
)

func oneForX(x string) {
	person := []string{x}

	for _, name := range person {
		if name == "" {
			fmt.Println("One for you, one for me")
		} else {
			fmt.Printf("One for %s, one for me\n", name)
		}
	}
}

func main() {
	oneForX("Devin")
	oneForX("")
}
