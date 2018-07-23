
package main


import (
	"hello/packs"
	 //"fmt"
)
func main() {
	//var (
	//	name = "刘浩宇"
	//	age = 123
	//	height int
	//)
	packs.GetInstallPacks()
	//area,per := calcu(1,2)

	//空指示符忽略其中一个变量
	//area, _ := calcu(1,2)
	//
	//fmt.Print(area)
}

func calcu(height, weight float64)(float64, float64)  {
	var area = height * weight
	var per = (height + weight) * 2

	return area, per
}