package main

import "fmt"

func main() {
	//1.往管道中增加数据是不能超过管道容量
	var intChan chan int = make(chan int, 3)
	intChan <- 10
	intChan <- 19
	intChan <- 20
	//2.从管道中取数据
	num := <-intChan
	fmt.Println(num)
	//3.打印管道容量
	fmt.Println(cap(intChan))
	//4.打印管道常量
	fmt.Println(len(intChan))
	//5.管道地址
	fmt.Println(intChan)
	//6.管道变量地址
	fmt.Println(&intChan)
	//7.如果管道中没有数据再取的话会报异常==>all goroutines are asleep - deadlock!
	num2 := <-intChan
	fmt.Println(num2)

	num3 := <-intChan
	fmt.Println(num3)

	intChan <- 20
	//num4 := <-intChan
	//fmt.Println(num4)
	//结构体（struct），接口（interface），分片(slice)，错误处理
	//channel关闭后不能再往里面写值了
	close(intChan)
	//intChan <- 100

	//遍历管道
	var intChan2 chan int = make(chan int, 100)
	for i := 1; i <= 100; i++ {
		intChan2 <- i
	}
	for i := 1; i <= 100; i++ {
		fmt.Println(<-intChan2)
	}
	//遍历管道之前先关闭管道
	close(intChan2)
	for v := range intChan2 {
		fmt.Println(v)
	}
}
