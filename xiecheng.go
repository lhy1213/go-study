package main

import (
	"fmt"
	"time"
)

func hello()  {
	fmt.Println("这是我输出的协程的内容")
}

func numbers()  {
	for i := 1; i<5; i++  {
		time.Sleep(200 * time.Millisecond)
	}
}

func alpha()  {
	for i := 1; i<10; i++  {
		time.Sleep(400 * time.Millisecond)
	}
}

func main()  {
	go hello()
	time.Sleep(1*time.Second)
	fmt.Println("这是主协程的内容")
}