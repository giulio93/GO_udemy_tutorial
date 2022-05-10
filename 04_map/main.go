package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#4f4444",
		"green": "#fd45444",
	}

	fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for color %v is %v\n", color, hex)
	}
}
