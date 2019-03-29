package initials

import "fmt"
//变量和方法首字母大写才能被别的包访问，否则只能被本包访问
const Url = "www.baidu.com"

func Test() {
	fmt.Println("test")
}
