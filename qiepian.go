package main

import "fmt"

func main()  {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	//fruitslice := fruitarray[1:3]
	//fmt.Printf("length of slice %d capacity %d\n", fruitslice, cap(fruitslice)) //length of is 2 and capacity is 6
	////fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
	//fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
	//fmt.Println("After re-slicing length is",fruitslice, "and capacity is",cap(fruitslice))

	fruitSlice1 := fruitarray[:len(fruitarray)-2]
	fruitSlice := make([]string, len(fruitSlice1))

	fmt.Print(cap(fruitSlice))

}