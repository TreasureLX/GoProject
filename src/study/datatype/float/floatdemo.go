package main

import "fmt"

func main() {
	//1.float32(单精度)和float64(双精度)，浮点数存在进度丢失的问题
	var price float32 = 23.000546464646
	var price2 float64 = -84646462324.65
	fmt.Println(price)
	//结果为-8.464646232465e+10通过科学计数法来表示
	fmt.Println(price2)
	//2.浮点数的存储结构==>符号位+指数位+尾数位
	//3. .999表示0.999,默认值为float64
	var n = .999
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	//4.科学计数法
	test := 5.003e-2
	fmt.Println(test)
}
