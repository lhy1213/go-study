package main

import "fmt"

func main()  {
	var mapOne map[string]int

	if mapOne == nil {
		fmt.Println("是空的")
		mapOne = make(map[string]int)
	}

	mapOne["name"] = 3
	mapOne["height"] = 123;

	fmt.Print(mapOne["name"])

	value, ok := mapOne["height"]

	if ok == true {
		fmt.Print(value)
	}else {
		fmt.Print("这是不存在的")
	}


}