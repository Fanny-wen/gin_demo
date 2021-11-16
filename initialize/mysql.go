package initialize

import (
	"fmt"
	"gin_demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB *gorm.DB // 定义一个全局对象DB
)

func InitMysqlConnect() (DB *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", GC.Mysql.User, GC.Mysql.Password, GC.Mysql.Host, GC.Mysql.Port, GC.Mysql.DbName, GC.Mysql.Charset, GC.Mysql.ParseTime, GC.Mysql.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	//设置最大的空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 自动生成数据库表 ，生成的表明默认是struct名称的复数形式。如：user -》users
	// 如果不想是复数形式，需设定：
	//db.SingularTable(true)
	//defer db.Close()
	fmt.Println("连接数据库成功！")

	//自动检查 Product 结构是否变化，变化则进行迁移
	_ = db.AutoMigrate(&model.User{}, &model.AdminUser{})

	return db, nil
}
