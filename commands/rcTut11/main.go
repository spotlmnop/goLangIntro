package main

import "fmt"

//Pointers
func main() {
	var i interface{} = "hello"

	s, ok := i.(string)
	fmt.Println(s, ok)

	do(21)
	do("hello")
	do(true)

}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
