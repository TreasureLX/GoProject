package prime

import (
	"time"
	"fmt"
)

func main() {
	var flag bool = false
	start := time.Now().Unix()
	for i := 1; i <= 200000; i++ {
		flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
		//	fmt.Println(i)
		}

	}
	end := time.Now().Unix()
	fmt.Printf("耗时%d", end-start)

}
