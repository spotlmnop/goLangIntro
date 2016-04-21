package main

import "fmt"

//Structs
type Employee struct {
	name string
	id   int
	salary int
}

func NewEmployee(name string, id int, salary int) *Employee{
	x := new(Employee);
	x.name = name
	x.id = id
	x.salary = salary
	return x
}

func (e *Employee) incSalary (raise int){
	e.salary += raise;
}
func main() {
	emp := new(Employee)
	fmt.Println(emp);

	emp1 := NewEmployee("Pradeep", 1, 100);
	emp1.incSalary(1000);

	fmt.Print(emp1)
}
