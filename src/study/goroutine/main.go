package main

//go语言协程用法
import (
	"fmt"
	"strconv"
	"runtime"
	"time"
	"sync"
)

const (
	Name     = ""
	Password = ""
)

var(
	lock sync.Mutex
)
func main() {
	//go test()
	//go test()
	//for i := 1; i < 10; i++ {
	//	fmt.Println("main hello golang:" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}
	runtime.GOMAXPROCS(1)
	//fmt.Print(runtime.NumCPU())
	//for{
	//	go test1()
	//	test0()
	//}
	//go test0()
	//go test1()
	go test(1)
	go test3(2)
	time.Sleep(time.Hour)

}

func test1()  {
	fmt.Print("a")
	fmt.Print("b")
	fmt.Print("c")

}
func test0()  {
	time.Sleep(time.Second)
	fmt.Print(1)
	fmt.Print(2)
	fmt.Print(3)
	time.Sleep(time.Hour)
}
func test(num int) {
	for i := 1; i < 10; i++ {
		fmt.Println(strconv.Itoa(num)+"test hello golang:" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test3(num int) {
	for i := 1; i < 10; i++ {
		fmt.Println(strconv.Itoa(num)+"test hello golang:" + strconv.Itoa(i))
		time.Sleep(time.Second)
		time.Sleep(time.Second)
	}
}