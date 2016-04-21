package main

import "fmt"

//Functions
func main() {
	x := add(1, 2)
	fmt.Println(x)
	a, b := addSubtract(5, 4)
	fmt.Println(a, b)

	sum := addArbitrary(1 ,2, 3, 4);
	fmt.Println(sum)

}

func add(x int, y int) int {
	return x + y
}

//multiple return types.
func addSubtract(x int, y int) (int, int) {
	return x + y, x - y
}

func addArbitrary(x... int) int {
	sum := 0
	for _, value := range x {
		sum += value;
	}
	return sum
}
