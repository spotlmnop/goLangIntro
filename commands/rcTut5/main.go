package main

import "fmt"

//Arrays
func main() {
	//Arrays
	var x [5]int; //Defining an array of 5 integers.
	x[2] = 10
	fmt.Println(x)

	b := [2]string{"Penn", "Teller"} // or b := [...]string{"Penn", "Teller"}
	fmt.Println(b);


	//Slices
	arr := []float64{1, 2, 3, 4, 5};   // Look above to see the difference
	slice1 := arr[0:5];
	fmt.Println(slice1);

	slice2 := make([]float64, 5)
	fmt.Println(slice2)

	//Iterate?

	//Maps


}
