package main

import "fmt"

func main() {
	//1.go语言没有字符类型，能用byte来表示字符
	//byte只能表示ASSCLL码中的字符，如果要存汉字则要用int来存
	var char1 byte='a'
	var char2 byte='0'
	var char3 int='兰'

	fmt.Printf("%c\n",char1)
	fmt.Printf("%c\n",char2)
	fmt.Printf("%c\n",char3)
}
