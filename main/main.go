package main

import (
	"github.com/gin-gonic/gin"
	"github.com/helloGin/controller"
	"net/http"
)

func main()  {

	// Engin
	router := gin.Default()

	//router := gin.New()
	router.LoadHTMLGlob("template/*")
	router.GET("/hello", hello)

	// 路由组
	user := router.Group("/user")
	{	// 请求参数在请求路径上
		user.GET("/get/:id/:username",controller.QueryById)
		user.GET("/query",controller.QueryParam)
		user.POST("/insert",controller.InsertNewUser)
		user.GET("/form",controller.RenderForm)
		user.POST("/form/post",controller.PostForm)
		//可以自己添加其他，一个请求的路径对应一个函数

		// ...
	}

	file := router.Group("/file")
	{
		// 跳转上传文件页面
		file.GET("/view",controller.RenderView)
		// 根据表单上传
		file.POST("/insert",controller.FormUpload)
		file.POST("/multiUpload",controller.MultiUpload)
		// base64上传
		file.POST("/upload",controller.Base64Upload)
	}

	// 指定地址和端口号
	router.Run(":9090")
}

func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")

	context.JSON(http.StatusOK,gin.H{
		"code":200,
		"success":true,
	})
}

