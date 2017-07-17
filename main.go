package main

import (
	"fmt"
	"gitlab.alipay-inc.com/lark/hello-go/utils"
)

func main() {
	newName := utils.FormatName("自成")
	monkey, _ := BuildMonkey("Foo", 5, 0)
	fmt.Println("Hello world.", newName, monkey)
}
