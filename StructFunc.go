package main

import "fmt"

type Employee struct {
	name string
	salary int
	currency string
}

func (a Employee) display()  {
	fmt.Printf("Salary of %s is %s%d", a.name, a.currency, a.salary)
}


func main()  {
	emp1 := Employee{
		name:"sdf",
		salary:123,
		currency:"123",
	}

	emp1.display()
}