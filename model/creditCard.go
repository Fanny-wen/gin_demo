package model

import (
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	CardId string `gorm:"not null"`
	User   User   `gorm:"ForeignKey:UserId;AssociationForeignKey:ID"`
	UserId int
}

/*
	AssociationForeignKey 可以指明BlogType中的哪个字段当做 关联的字段， 可以将ID改成Name
	这样关联的字段就是Name了
*/
