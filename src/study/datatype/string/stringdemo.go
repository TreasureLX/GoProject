package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1.默认大小为16个字节
	var address = "Hello world"
	fmt.Printf("%d\n", unsafe.Sizeof(address))
	//2.字符串一旦被赋值是不能修改的，即GO中的字符串是不可变的
	//address[0]='3'
	//3.有转移字符会自动转移
	var test = "abc\neeee"
	fmt.Println(test)
	//4.``号不识别特殊字符
	str := `package main

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
}`
	fmt.Println(str)
	//5.字符串拼接,换行时加号留在最后，防止编译报错
	var str2 = "Hello" + " world" +
		"Hello" + " world" +
		"Haddfd"
	str2 += " test"
	fmt.Println(str2)
}
