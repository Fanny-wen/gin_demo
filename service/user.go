package service

import (
	"fmt"
	"gin_demo/initialize"
	"time"
)

type User struct {
	Id   int    `gorm:"primary_key" json:"id"` //表字段名为：id,主键
	Name string `gorm:"not null" json:"name"`  //模型标签 name字段不能为空
	Age  int    `json:"age"`                   //年龄age字段
	//对应表中的create_time字段并且不为空
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

// InsertUser 插入数据
func (user *User) InsertUser() (err error) {
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。db.Create(user)也可以
	createDb := initialize.DB.Table("user").Create(user)
	err = createDb.Error
	if err != nil {
		fmt.Println("新增数据错误,err", err)
		return err
	}
	return nil
}
