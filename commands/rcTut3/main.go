package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Using for loop as while.
	j := 10

	for j >= 0 {
		if j%2 == 0 {
			fmt.Println("j --> ", j)
		}
		j--
	}

}
