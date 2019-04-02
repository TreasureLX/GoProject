package main

import (
	"fmt"
	"time"
	"runtime"
)

func putNum(intChan chan int, count int) {
	for i := 1; i <= count; i++ {
		intChan <- i
	}
	close(intChan)
}

func getPrime(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool = true
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	exitChan <- true

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()-1)
	count := 200000
	intChan := make(chan int, count)
	primeChan := make(chan int, 100000)
	resultChan := make(chan bool, 4)

	go putNum(intChan, count)
	start := time.Now().Unix()
	for i := 1; i <= 4; i++ {
		go getPrime(intChan, primeChan, resultChan)
	}

	go func() {
		for i := 1; i <= 4; i++ {
			<-resultChan
		}
		end := time.Now().Unix()
		fmt.Printf("耗时%d", end-start)
		close(primeChan)
	}()

	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(v)
	}

}

