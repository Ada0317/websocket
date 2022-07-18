package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io"
	"os"
)

func main() {
	router := gin.Default() //启动默认的路由引擎实例服务   携带基础中间件启动 logger 和 recovery (启动了一个服务器

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameRule", ruleName)
	}

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
		usr := c.PostForm("usr")

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

	type User struct {
		Name string `json:"name" form:"name" uri:"name" binding:"required,NameRule"` //如果加上了binding标签但是不传参就会报错
		Sexy bool   `json:"sexy" form:"sexy" uri:"sexy" binding:"required"`
	}

	//router.POST("TestBind/:name/:sexy", func(c *gin.Context) {  err := c.ShouldBindUri(usr)
	router.POST("TestBind", func(c *gin.Context) {
		usr := new(User)
		err := c.ShouldBindJSON(usr)
		//err := c.ShouldBindQuery(usr)
		//err := c.ShouldBindUri(usr) //  ":/name/:sexy" 的形式

		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"err": err,
			})
		} else {
			c.JSON(200, gin.H{
				"usr": usr,
			})
		}
	})

	//单个文件传输
	router.POST("UploadSingleFile", func(c *gin.Context) {
		file, _ := c.FormFile("file") //文件流和 err

		//c.SaveUploadedFile(file, "./"+file.Filename) //保存文件到本地
		//name := c.PostForm("name")
		in, _ := file.Open()
		defer in.Close()

		out, _ := os.Create("./" + file.Filename)
		defer out.Close()

		io.Copy(out, in)

		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))

		c.File("./" + file.Filename)

		//c.JSON(200, gin.H{
		//	"filename": file.Filename,
		//	"filesize": file.Size,
		//})
	})

	//多个文件传输
	router.POST("UploadMultiFile", func(c *gin.Context) {
		form, _ := c.MultipartForm() //
		FileList := form.File["file"]

		for _, item := range FileList {
			//c.SaveUploadedFile(item, "./"+item.Filename)

			in, _ := item.Open()
			defer in.Close()
			out, _ := os.Create("./" + item.Filename)
			defer out.Close()
			io.Copy(out, in)

			c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", item.Filename))

			c.File("./" + item.Filename)

		}

		//name := c.PostForm("name")
		//c.JSON(200, gin.H{
		//	"filename": file.Filename,
		//	"filesize": file.Size,
		//})
	})

	router.Run(":10101") //1024到65536(16位二进制的最大数
}

func ruleName(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(string); ok {
		if len(date) <= 3 {
			return false
		} else {
			return true
		}
	}
	return false
}

//func ff(fl validator.FieldLevel) bool {
//	if i, ok := fl.Field().Interface().(int); ok {
//
//	}
//}
