package main

import (
	"fmt"
)

func main() {
	//1.go语言中需要显式转换不能自动转换，低精度->高精度
	//eg:将int装换为float
	var num int = 10
	//转换完后num还是为int类型
	var num2 float32 = float32(num)
	fmt.Println(num, num2)
	//64位转8位结果会溢出
	var num3 int64 = 99999999
	var num4 int8 = int8(num3)
	fmt.Println(num3, num4)

}
