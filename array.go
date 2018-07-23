package main

import "fmt"

func main()  {
	//a := [...]float64{76.3,98.8,21,78}
	a := [3][3]string{
		{"a","b"},
		{"c","d"},
		{"e","f"},
	}
	//sum := float64(0)
	for _,v := range a {
		for _,v := range v {
			fmt.Print(v)
		}
	}
}