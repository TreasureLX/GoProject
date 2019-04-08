package main

import "fmt"

func getValue(n, m int) (int, int) {
	sum := n + m
	sub := n - m
	return sum, sub
}

func main() {
	sum, sub := getValue(10, 5)
	fmt.Println(sum)
	fmt.Print(sub)

	var test int
	fmt.Println(test)
	fmt.Print(test)
}
