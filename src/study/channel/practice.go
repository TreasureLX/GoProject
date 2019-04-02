package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		time.Sleep(time.Second*10)
		intChan <- i
		fmt.Printf("写入数据%d\n", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		//fmt.Println(ok)
		if !ok {
			break
		}
		fmt.Printf("读数据%d\n", v)
	}
	exitChan <- true
	close(exitChan)
}
func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		//如果管道中没有数据可读会，这里会阻塞，直到有数据为止
		//后面的ok表示管道是否关闭
		_, ok := <-exitChan
		fmt.Println(ok)
		if !ok {
			break
		}
	}
}
