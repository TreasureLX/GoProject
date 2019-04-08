package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	//类型推导
	//1.var test=10
	//2.test:=10
	//3.var n int=10
	//0-255  256个数字
	// 有符号int8 int16 int32 int64
	// 无符号uint8 uint16 uint32 uint64
	var n1 int8=127
	fmt.Println(n1)
	var n2 =100
	//查看某个值得类型
	fmt.Printf("%T\n",n2)
	var n3 rune=111
	var b byte=12
	fmt.Println(b)
	fmt.Println(n3)
	//查看某个int变量的占用的字节大小
	var n4 int64=646464646
	fmt.Printf("数值：%d，占用：%d字节空间\n",n4,unsafe.Sizeof(n4))
}





