package router

import (
	"gin_demo/Middlewares"
	"gin_demo/api/Notfindpage"
	"gin_demo/api/Redirect"
	"gin_demo/api/Student"
	"gin_demo/api/Upload"
	"gin_demo/api/User"
	"gin_demo/initialize"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func NewRouter() *gin.Engine {
	var Engine = gin.New()
	gin.ForceConsoleColor()
	f, err := os.OpenFile("./logs/gin.log", os.O_WRONLY|os.O_APPEND|os.O_APPEND, 0777)
	if err != nil {
		panic("./logs/gin.log not find")
	}
	gin.DefaultWriter = io.MultiWriter(f)

	Engine.LoadHTMLGlob("./static/*")
	Engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	Engine.Use(gin.Recovery())
	Engine.Use(Middlewares.LogMiddleware())

	Engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": map[string]interface{}{
				"message": "success",
				"code":    http.StatusOK,
			},
		})
	})
	Engine.GET("/config", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": map[string]interface{}{
				"message": "success",
				"code":    http.StatusOK,
				"config":  initialize.GC,
				"time":    time.Now().Format("2006-01-02 15:04:05"),
			},
		})
	})

	Engine.Any("/404", Notfindpage.PageNotFindHandler)

	Engine.NoRoute(Redirect.RedirectHandler)

	upload := Engine.Group("upload")
	{
		upload.POST("", Upload.UploadHandler)
		upload.POST("/multiple", Upload.UploadMultipleHandler)
	}

	student := Engine.Group("student")
	{
		student.GET("/detail/:id", Student.DetailStudentHandler)
		student.GET("/list", Student.ListStudentHandler)
		student.POST("/create", Student.CreateStudentHandler)
		student.DELETE("/delete/:id", Student.DeleteStudentHandler)
		student.PUT("/update/:id", Student.UpdateStudentHandler)
	}
	user := Engine.Group("user")
	{
		user.GET("/:id", User.DetailUserHandler)
		user.GET("/list", User.ListUserHandler)
		user.POST("/create", User.CreateUserHandler)
		user.DELETE("/:id", User.DeleteUserHandler)
		user.PUT("/:id", User.UpdateUserHandler)
	}
	admin := Engine.Group("admin")
	{
		admin.GET("/:id", User.DetailAdminUserHandler)
		admin.GET("/list", User.ListAdminUserHandler)
		admin.DELETE("/:id", User.DeleteAdminUserHandler)
		admin.PUT("/:id", User.UpdateUserHandler)
	}
	return Engine
}
