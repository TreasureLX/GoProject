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


}
