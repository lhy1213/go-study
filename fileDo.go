package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main()  {
	filepath.Walk("G:\\ebook", walkfunc)
}

func walkfunc(path string, info os.FileInfo, err error)  error{
	fmt.Println(path)
	return nil
}