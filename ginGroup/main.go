package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use()
	//中间件 洋葱模型
	v1 := router.Group("v1").Use(middle1(), middle2()) //路由更加清晰 路由管理更加便捷 可以group加中间件

	v1.GET("testGet", func(c *gin.Context) {
		fmt.Println("我在方法get方法里")
		c.JSON(200, gin.H{
			"success": "1",
		})
	})
	router.Run(":9090")
}

func middle1() gin.HandlerFunc { //中间件 框架
	return func(c *gin.Context) {
		fmt.Println("我在方法middle1前")
		c.Next()
		fmt.Println("我在方法middle1后")
	}
}

func middle2() gin.HandlerFunc { //中间件 框架
	return func(c *gin.Context) {
		fmt.Println("我在方法middle2前")
		c.Next()
		fmt.Println("我在方法middle2后")
	}
}
