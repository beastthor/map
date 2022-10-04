package main

import "fmt"

func main() {
	//maps representation 2
	//var colors map[string]string

	//maps represntation 3
	//colors := make(map[int]string)

	//maps represntation 1
	colors := map[string]string{

		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//Asinging a value to map after creating it
	//colors[10] = "#ffffff"

	//delete(colors, 10)

	printMap(colors)

}
func printMap(c map[string]string) {

	for color, hex := range c {

		fmt.Println("Hex code for", color, "is", hex)

	}

}
