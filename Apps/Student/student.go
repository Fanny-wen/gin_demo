package Student

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// 定义一个添加学生参数结构体
type Student struct {
	Name        string `json:"name" validate:"required"`             //必填
	Email       string `json:"email" validate:"required,email"`      // 必填，并且格式是email
	Age         uint8  `json:"age" validate:"gte=0"`                 // 年龄范围
	Gender      string `json:"gender" validate:"required"` // 性别
	PhoneNumber string `json:"phone_number" validate:"required"`     // 手机号
}

// DetailStudentHandler 学生信息详情
func DetailStudentHandler(c *gin.Context) {

}

// ListStudentHandler 学生信息列表
func ListStudentHandler(c *gin.Context) {

}

// CreateStudentHandler 新增学生
func CreateStudentHandler(c *gin.Context) {
	var s Student
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(500, gin.H{"message": err})
		return
	}
	fmt.Printf("student: %+v\n", s)
	// 使用Validate验证
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"message": "success",
		"data": gin.H{
			"student": s,
		},
	})
}

// DeleteStudentHandler 删除学生
func DeleteStudentHandler(c *gin.Context) {

}

// UpdateStudentHandler 更新学生信息
func UpdateStudentHandler(c *gin.Context) {

}
