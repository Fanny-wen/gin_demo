package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string       `gorm:"not null;size:20"` //模型标签 name字段不能为空
	Age     *int         `gorm:"not null"`         //年龄age字段
	IsAdmin sql.NullBool `gorm:"default:false"`
	UUID    string       `gorm:"not null"`
}

// TableName 设置 `User` 的表名为 `user`
//func (User) TableName() string {
//	return "user"
//}

type AdminUser struct {
	gorm.Model
	User   User `gorm:"ForeignKey:UserId;AssociationForeignKey:ID"`
	UserId uint
}

// TableName 设置 `User` 的表名为 `user`
//func (AdminUser) TableName() string {
//	return "admin_user"
//}
