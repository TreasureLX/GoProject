package main

import (
	"fmt"
	"strconv"
)

func main() {
	//方式一:通过Sprintf方法来转换类型
	var num int = 10
	var num2 float32=343.554
	dataToString("%d",num)
	dataToString("%f",num2)
	//方式二：通过strconv包来转换类型
	var str1=strconv.Itoa(num)
	printValueAndType(str1)
}

func dataToString(format string,a interface{}) string  {
	var value=fmt.Sprintf(format,a)
	printValueAndType(value)
	return value
}

func printValueAndType(a interface{})  {
	fmt.Printf("value->%q,type->%T\n",a,a)
}

