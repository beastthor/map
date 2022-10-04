package main

import "fmt"

func main() {
	//maps representation 2
	//var colors map[string]string

	//maps represntation 3
	colors := make(map[int]string)

	//maps represntation 1
	//colors := map[string]string{

	//	"red":   "#ff0000",
	//	"green": "#4bf745",
	//}

	//Asinging a value to map after creating it
	colors[10] = "#ffffff"

	fmt.Println(colors)

}
