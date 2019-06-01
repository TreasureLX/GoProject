package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 1; i < 10; i++ {
		fmt.Println("Hello world")
	}
}

func test() {
	//错误处理机制defer+recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生错误",err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang"
}

func main2() {
	go test()
	go sayHello()
	time.Sleep(time.Minute)
}
