package model

import (
	"time"
)

type User struct {
	Id   int    `gorm:"primary_key" json:"id"` //表字段名为：id,主键
	Name string `gorm:"not null" json:"name"`  //模型标签 name字段不能为空
	Age  int    `json:"age"`                   //年龄age字段
	//对应表中的create_time字段并且不为空
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}
