
package main

import (
	"fmt"
)

type SalaryCalculator interface {  //定义了一个接口，里面有一个方法
	CalculateSalary() int
}

type Permanent struct {  //定义了一个结构体
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId  int
	basicpay int
}

//定义了接口里面的函数，传入了Permanent结构体，相当于Permanent继承了SalaryCalculator接口
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

//空接口，可以传入任何类型
func describe(i interface{})  {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(i interface{})  {
	v, ok := i.(int)
	fmt.Println(v,ok)
}


func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}


func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	fmt.Print(employees)
	totalExpense(employees)


	assert("哈哈哈")


	findType("Naveen")
	findType(77)
	findType(89.98)
}