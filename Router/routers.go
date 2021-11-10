package Router

import (
	"gin_demo/Apps/Notfindpage"
	"gin_demo/Apps/Redirect"
	"gin_demo/Apps/Upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

//var Engine = gin.Default()
var Engine = gin.New()

func init() {
	Engine.LoadHTMLGlob("./Assets/static/*")

	Engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": map[string]interface{}{
				"message": "success",
				"status":  http.StatusOK,
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
}
