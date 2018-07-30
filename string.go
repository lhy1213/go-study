package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	stringTest := "sdfsdfsdfsdf"
	for k,v := range stringTest{
		fmt.Printf("%c starts at byte %d\n", v, k)
	}

	stringUFT8 := "SDFSDF张三"
	fmt.Printf("length of %s is %d\n", stringUFT8, utf8.RuneCountInString(stringUFT8))
}