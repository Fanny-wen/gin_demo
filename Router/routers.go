package Router

import (
	"fmt"
	"gin_demo/Apps/Notfindpage"
	"gin_demo/Apps/Redirect"
	"gin_demo/Apps/Student"
	"gin_demo/Apps/Upload"
	"gin_demo/Middlewares"
	"gin_demo/Services/viper"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

//var Engine = gin.Default()
var Engine = gin.New()

func init() {
	gin.ForceConsoleColor()
	f, err := os.OpenFile("./Assets/logs/gin.log", os.O_WRONLY|os.O_APPEND|os.O_APPEND, 0777)
	if err != nil {
		panic("./Assets/logs/gin.log not find")
	}
	fmt.Printf("%+v\n", f.Name())
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	Engine.LoadHTMLGlob("./Assets/static/*")
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
				"data":    viper.ServerConfig,
				"time": time.Now().Format("2006-01-02 15:04:05"),
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
}
