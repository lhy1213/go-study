
package main

import "fmt"

func main() {
	//var (
	//	name = "刘浩宇"
	//	age = 123
	//	height int
	//)

	area,per := calcu(1,2)
	//fmt.Println("my name is", name, ", age is", age, "and height is", height)
	fmt.Print(area,per)
}

func calcu(height, weight float64)(float64, float64)  {
	var area = height * weight
	var per = (height + weight) * 2

	return area, per
}