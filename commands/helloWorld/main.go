package main

import "fmt"

func main() {
	fmt.Println("Hello World !");

	//Types
	fmt.Println("1 + 1 =", 1 + 1);

	fmt.Println(len("Hello World"));

	fmt.Println(!true);

	variables();
	constants();


}

func variables() {
	// Variables
	var x string
	x = "Yo"
	//The only reason var exists is because const exist.
	//y := "Yo"
	fmt.Println(x)
	// fmt.Println("Are yo's equals: ", x == y )
	var (
		a = 5
		b = 10
		c = 15
	)

	fmt.Println(a, b, c);

}

func constants() {
	const y string = "spot";
	fmt.Println(y);

}
