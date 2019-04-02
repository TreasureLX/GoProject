package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var cat1 Cat
	fmt.Printf("%p", &cat1 )
	cat1.Name = "tom"
	cat1.Age = 1
	//fmt.Println(cat1)
	//fmt.Println(&cat1)
	//创建结构体的四种方式

}
