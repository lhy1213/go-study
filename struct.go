package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age,salary int
}
func main()  {
	emp5 := Employee{
		firstName:"John",
		lastName:"hahaha",
	}
	fmt.Println(emp5)
}