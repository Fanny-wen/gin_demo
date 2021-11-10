package Upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UploadHandler 上传单个文件
func UploadHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	if file != nil {
		fmt.Printf("filename: %s\n", file.Filename)
		err := c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			fmt.Printf("save uploaded file failed! err: %v\n", err)
			panic("save uploaded file failed!")
		}
		c.JSON(http.StatusOK, gin.H{
			"file":    file.Filename,
			"size":    file.Size,
			"message": "success",
			"status":  200,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
			"status":  400,
		})
	}
}

// UploadMultipleHandler 上传多文件
func UploadMultipleHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err,
				"status":  400,
			})
		}
	}()
	var data []interface{}
	form, err := c.MultipartForm()
	if form == nil && err != nil {
		panic(err)
	}
	files := form.File["file"]
	value := form.Value // Value 属性保存除 文件外的其他body数据
	fmt.Printf("form's value: %v\n", value)

	for _, file := range files {
		fmt.Printf("filename: %s\n", file.Filename)
		err := c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			fmt.Printf("save uploaded file failed! err: %v\n", err)
			panic("save uploaded file failed!")
		}
		info := map[string]interface{}{
			"filename": file.Filename,
			"size":     file.Size,
		}
		data = append(data, info)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
