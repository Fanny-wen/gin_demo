package User

import (
	"fmt"
	"gin_demo/initialize"
	"gin_demo/model"
	"gin_demo/service"
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
User
*/

// DetailUserHandler 学生信息详情
func DetailUserHandler(c *gin.Context) {
	var u = &service.User{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := u.GetUser(id); err != nil {
		fmt.Println("查询数据错误,err", err.Error())
		panic("查询数据错误")
	}
	{
		// struct 与 map 之间的转换
		m := util.Struct2Map(*u, 8)
		fmt.Printf("%+v\n", m)
		var u2 service.User
		_ = util.Map2Struct(m, &u2)
		fmt.Printf("%+v\n", u2)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"user": u,
		},
	})
}

// ListUserHandler 学生信息列表
func ListUserHandler(c *gin.Context) {
	var u *service.User
	//u := service.User{}
	var users []service.User
	//users := new([]service.User)
	//users := make([]service.User, size, size)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	condition := make(map[string]interface{}) // ok
	//condition := map[string]interface{}{}     // ok
	//var condition = map[string]interface{}{}  // ok
	//var condition map[string]interface{}      // 这种方法不可取, 只定义, 未初始化, 引用类型的变量默认为nil
	if name := c.Query("name"); name != "" {
		condition["name"] = name
	}
	if age := c.Query("age"); age != "" {
		condition["age"], _ = strconv.Atoi(age)
	}
	if uuid := c.Query("uuid"); uuid != "" {
		condition["uuid"] = uuid
	}

	//调用user表的获取列表数据
	if err := u.GetListUser(&condition, &users, page, size); err != nil {
		fmt.Println("列表数据错误,err", err.Error())
		panic("列表数据错误")
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    users,
		"page":    page,
		"size":    size,
	})
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
	//调用user表的插入操作
	if err := u.InsertUser(u); err != nil {
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
	var u = &service.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := u.GetUser(id); err != nil {
		fmt.Println("获取数据错误,err", err.Error())
		panic("获取数据错误")
	}
	if err := u.DeleteUser(); err != nil {
		fmt.Println("删除数据错误,err", err.Error())
		panic("删除数据错误")
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"user": u,
		},
	})
}

// UpdateUserHandler 更新学生信息
func UpdateUserHandler(c *gin.Context) {
	var u = make(map[string]interface{})  // 接收数据
	var u2 = &service.User{} // 原始数据
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := u2.GetUser(id); err != nil {
		fmt.Println("获取数据错误,err", err.Error())
		panic("获取数据错误")
	}
	if err := u2.UpdateUser(u); err != nil {
		fmt.Println("修改数据错误,err", err.Error())
		panic("修改数据错误")
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"user": u,
		},
	})
}

/*
AdminUser
*/

// DetailAdminUserHandler 学生信息详情
func DetailAdminUserHandler(c *gin.Context) {
	var u = &service.User{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := u.GetAdminUser(id); err != nil {
		fmt.Println("查询数据错误,err", err.Error())
		panic("查询数据错误")
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"user": u,
		},
	})
}

// DeleteAdminUserHandler 删除学生
func DeleteAdminUserHandler(c *gin.Context) {
	var u = &service.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := u.GetAdminUser(id); err != nil {
		fmt.Println("获取数据错误,err", err.Error())
		panic("获取数据错误")
	}
	if err := u.DeleteUser(); err != nil {
		fmt.Println("删除数据错误,err", err.Error())
		panic("删除数据错误")
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"user": u,
		},
	})
}

// ListAdminUserHandler 学生信息列表
func ListAdminUserHandler(c *gin.Context) {
	var u *service.User
	var users []service.User
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	initialize.DB.Table("user").Model(model.User{}).Offset((page - 1) * size).Limit(size).Find(&users)

	//调用user表的获取列表数据
	if err := u.GetListAdminUser(&users, page, size); err != nil {
		fmt.Println("列表数据错误,err", err.Error())
		panic("列表数据错误")
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    users,
		"page":    page,
		"size":    size,
	})
}
