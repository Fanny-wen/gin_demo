package initialize

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB // 定义一个全局对象DB
)

func InitMysqlConnect() (DB *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", GC.Mysql.User, GC.Mysql.Password, GC.Mysql.Host, GC.Mysql.Port, GC.Mysql.DbName, GC.Mysql.Charset, GC.Mysql.ParseTime, GC.Mysql.Loc)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//设置最大的空闲连接数
	db.DB().SetMaxIdleConns(10)
	// 设置最大连接数
	db.DB().SetMaxOpenConns(100)
	//自动生成数据库表 ，生成的表明默认是struct名称的复数形式。如：user -》users
	// 如果不想是复数形式，需设定：
	db.SingularTable(true)
	//defer db.Close()
	fmt.Println("连接数据库成功！")
	return db, nil
}
