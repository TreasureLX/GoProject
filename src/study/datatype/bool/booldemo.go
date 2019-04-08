package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1.bool占用一个的空间为一个字节
	var flag bool=true
	if flag {
		fmt.Println("Hello")
		fmt.Printf("%d",unsafe.Sizeof(flag))
	}
	//2.bool类型的取值只能为true或者false
}
