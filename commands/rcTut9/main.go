package main

import "fmt"


//Interfaces
type SimpleInterface interface {
	sayHi()
}

type Employee struct {
	name string
}

func ( e Employee) sayHi(){
	fmt.Print(e.name)
}


func main() {
	//Literal declaration
	var e SimpleInterface = Employee{ "Pradeep"}
	e.sayHi()



}
