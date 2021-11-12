package User

import (
	"fmt"
	"gin_demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DetailUserHandler 学生信息详情
func DetailUserHandler(c *gin.Context) {

}

// ListUserHandler 学生信息列表
func ListUserHandler(c *gin.Context) {

}

// CreateUserHandler 新增学生
func CreateUserHandler(c *gin.Context) {
	var u *service.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	fmt.Printf("User: %+v\n", u)
	// 使用Validate验证
	//validate := validator.New()
	//if err := validate.Struct(u); err != nil {
	//	fmt.Printf("%+v\n", err)
	//	c.JSON(500, gin.H{"message": err.Error()})
	//	return
	//}

	//连接数据库
	//err := initialize.InitMysqlConnect(initialize.DB)
	//if err != nil {
	//	fmt.Println("初始化数据库失败,err", err)
	//	c.JSON(500, gin.H{"message": err.Error()})
	//	return
	//}

	//调用user表的插入操作
	if err := u.InsertUser(); err != nil {
		fmt.Println("新增数据错误,err", err.Error())
		panic("新增数据错误")
	}
	//新增结束后会返回对应结构的所有数据，比如id
	fmt.Println("新增数据成功，新增user机构体主键是", u.Id)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"user": u,
		},
	})
}

// DeleteUserHandler 删除学生
func DeleteUserHandler(c *gin.Context) {

}

// UpdateUserHandler 更新学生信息
func UpdateUserHandler(c *gin.Context) {

}
