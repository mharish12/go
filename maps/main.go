package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	// colors := map[string]string

	// colors := make(map[string]string)

	// delete(colors, "red")

	fmt.Println(colors)

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hexcode for", color, "is", hex)
	}
}
