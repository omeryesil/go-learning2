package main

import "fmt"

func main() {

	//defining maps ----------------------------
	//option 1
	colors := map[string]string{
		"red":   "ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
	}

	//option 2
	//var colors2 map[string]string
	//colors2["red"] = "ff0000"

	//option 3
	//var colors3 = make(map[string]string)
	//colors3["red"] = "ff0000"

	//deleting a map ---------------------------
	delete(colors, "red")
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("Color: %s - HexValue: %s\n", key, value)
	}
}
