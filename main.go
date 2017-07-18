package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.alipay-inc.com/lark/hello-go/utils"
)

func main() {
	newName := utils.FormatName("自成")
	fmt.Println("Hello world.", newName)

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		monkey, _ := BuildMonkey("Foo", 5, 0)
		ctx.JSON(200, monkey)
	})

	router.Run(":8080")

}
