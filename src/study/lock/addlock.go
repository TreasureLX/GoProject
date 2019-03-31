package main

import (
	"time"
	"fmt"
	"strconv"
	"sync"
)

var (
	result = make(map[int]int)
	lock sync.Mutex
)

func main() {
	for i := 1; i < 200; i++ {
		go addNum(i)
	}
	time.Sleep(time.Second * 6)
	//打印Map
	for i, v := range result {
		fmt.Printf("key==>%s,value==>%s\n",strconv.Itoa(i),strconv.Itoa(v))
	}
}

func addNum(num int) {
	var sum int = 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	lock.Lock()
	result[num] = sum
	lock.Unlock()
}
