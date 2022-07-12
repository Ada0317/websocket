package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()                        //启动默认的路由引擎实例服务   携带基础中间件启动 logger 和 recovery (启动了一个服务器
	router.GET("/ping/:id", func(c *gin.Context) { //实例 接受get请求
		id := c.Param("id")
		//usr := c.Query("usr")
		usr := c.DefaultQuery("usr", "zy") //如果不早query中传递usr 则默认取值zy
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{ //返回json结构体信息
			"get":  "pong",
			"id":   id,
			"usr ": usr,
			"pwd":  pwd,
		})
	})
	router.POST("/ping", func(c *gin.Context) { //实例 接受get请求  query和form 联合传参
		query := c.Query("query")
		usr := c.DefaultPostForm("usr", "007")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{ //返回json结构体信息
			"post":  "pong",
			"query": query,
			"usr ":  usr,
			"pwd":   pwd,
		})
	})
	router.DELETE("/ping", func(c *gin.Context) { //实例 接受get请求
		//id := c.Param("id")
		usr := c.DefaultPostForm("usr", "007")
		c.JSON(200, gin.H{ //返回json结构体信息
			"delete": "pong",
			//"id":     id,
			"usr": usr,
		})
	})

	router.PUT("/ping", func(c *gin.Context) { //实例 接受get请求
		usr := c.DefaultPostForm("usr", "zzw")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{ //返回json结构体信息
			"put":  "pong",
			"usr ": usr,
			"pwd":  pwd,
		})
	})

	router.Run(":10101") //1024到65536(16位二进制的最大数
}
