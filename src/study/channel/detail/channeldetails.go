package main

import (
	"strconv"
	"fmt"
)

func main() {
	//管道可以声明为只读和只写

	//1.在默认情况下，管道是双向的
	//var channel1 chan int 可读可写

	//2.声明为只读
	//var onlyReadChan chan <- int

	//3.声明为只写
	//var onlyWriteChan <- chan int

	//2.使用select可以解决从管道读取数据阻塞问题
	intChan := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		intChan <- i
	}

	strChan := make(chan string, 5)
	for i := 1; i <= 5; i++ {
		strChan <- "hello" + strconv.Itoa(i)
	}

	//用传统方法遍历管道时，如果不关闭则会阻塞而导致deadLock
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(<-intChan)
	//}

	//1.可以通过select来解阻塞问题

	for {
		select {
		case v := <-intChan:
			fmt.Println(v)
		case v := <-strChan:
			fmt.Println(v)
		default:
			return
		}
	}


}
