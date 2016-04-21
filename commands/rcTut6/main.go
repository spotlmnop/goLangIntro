package main

import "fmt"

//Arrays
func main() {
	var x map[string]int
	x["key"] = 10
	fmt.Println(x)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	fmt.Println(elements["Li"])

	name, ok := elements["Un"]
	fmt.Println(name, ok)
	x = 11

	//Really cool way to do things in Go.
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}
