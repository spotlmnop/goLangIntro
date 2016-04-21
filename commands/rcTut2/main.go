package main

import "fmt"

func main() {
	var x string
	x = "Yo"
	//The only reason var exists is because const exists.
	//y := "Yo"
	fmt.Println(x)
	// fmt.Println("Are yo's equals: ", x == y )
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c)

	const spot string = "spot"
	fmt.Println(spot)
}
