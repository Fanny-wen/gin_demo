package service

import (
	"fmt"
	"gin_demo/initialize"
	"gin_demo/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint       `json:"id" form:"id"`     //表字段名为：id,主键
	Name      string     `json:"name" form:"name"` //模型标签 name字段不能为空
	Age       int        `json:"age" form:"age"`   //年龄age字段
	IsAdmin   bool       `json:"is_admin" form:"is_admin"`
	UUID      string     `json:"uuid"`
	CreatedAt time.Time  `json:"created_at" form:"created_at" formatter:"2006-01-02 15:04:05"`
	UpdatedAt time.Time  `json:"updated_at" form:"updated_at" formatter:"2006-01-02 15:04:05"`
	DeletedAt *time.Time `json:"deleted_at" form:"deleted_at" formatter:"2006-01-02 15:04:05"`
}

// UserTable 使用动态表名
func UserTable(user *User) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("users")
	}
}

// BeforeCreate 钩子, 创建记录时会调用这些方法 BeforeSave, BeforeCreate, AfterSave, AfterCreate
/*
	开始事务
	BeforeSave
	BeforeCreate
	关联前的 save
	插入记录至 db
	关联后的 save
	AfterCreate
	AfterSave
	提交或回滚事务
*/
func (user *User) BeforeCreate(db *gorm.DB) error {
	user.UUID = uuid.New().String()
	return nil
}
func (user *User) AfterCreate(db *gorm.DB) error {
	if user.IsAdmin == true {
		if createDb := db.Table("admin_users").Create(&model.AdminUser{
			UserId: user.Id,
		}); createDb.Error != nil {
			return createDb.Error
		}
	}
	return nil
}

// 查询钩子 AfterFind
/*
	从 db 中加载数据
	Preloading (eager loading)
	AfterFind
*/

// 更新钩子, 更新记录时会调用这些方法 BeforeSave, BeforeUpdate, AfterUpdate, AfterSave
/*
	开始事务
	BeforeSave
	BeforeUpdate
	// 关联前的 save
	// 更新 db
	// 关联后的 save
	AfterUpdate
	AfterSave
	提交或回滚事务
*/

func (user *User) BeforeUpdate(db *gorm.DB) error {
	fmt.Println(db.Statement.Changed("IsAdmin"))
	if ok := db.Statement.Changed("IsAdmin"); ok {
		if user.IsAdmin == true {
			fmt.Println("admin_users delete")
			db.Unscoped().Where("user_id = ?", user.Id).Delete(&model.AdminUser{})
		} else {
			fmt.Println("admin_users create")
			db.Table("admin_users").Create(&model.AdminUser{
				UserId: user.Id,
			})
		}
	}
	return nil
}

/*
	开始事务
	BeforeDelete
	删除 db 中的数据
	AfterDelete
	提交或回滚事务
*/

// InsertUser 插入数据
func (user *User) InsertUser(data interface{}) (err error) {
	// 用选定字段的来创建, DB.Select()
	//createDb := initialize.DB.Scopes(UserTable(user)).Select("Name").Create(data)
	// 创建时排除选定字段, DB.Omit()
	//createDb := initialize.DB.Omit("Age", "UpdatedAt").Create(data)
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。db.Create(user)也可以
	//createDb := initialize.DB.Create(data)
	createDb := initialize.DB.Scopes(UserTable(user)).Create(data)
	err = createDb.Error
	if err != nil {
		fmt.Println("新增数据错误,err", err)
		return err
	}
	return nil
}

/*
	GORM 提供 First, Take, Last 方法，以便从数据库中检索单个对象
	获取第一条记录（主键升序）
	db.First(&user)
	获取一条记录，没有指定排序字段
	db.Take(&user)
	获取最后一条记录（主键降序）
	db.Last(&user)
*/
// GetUser 获取一条数据
func (user *User) GetUser(id int) (err error) {
	db := initialize.DB.First(user, id)
	err = db.Error
	if err != nil {
		fmt.Println("查询数据错误,err:", err)
		return err
	}
	return nil
}

// GetListUser 获取列表数据
func (user *User) GetListUser(condition *map[string]interface{}, users *[]User, page, size int) (err error) {
	Db := initialize.DB.Model(model.User{}).Where(*condition).Offset((page - 1) * size).Limit(size).Find(&users)
	if err = Db.Error; err != nil {
		fmt.Println("列表数据错误,err:", err)
		return err
	}
	return nil
}

/*
	update 更新单个字段
	updates 更新多个字段, 通过 `struct` 更新多个字段，不会更新零值字段; 通过 `map` 更新多个字段，零值字段也会更新
	db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
*/

// UpdateUser 更新数据
//func (user *User) UpdateUser(value map[string]interface{}) (err error) {
//func (user *User) UpdateUser(u *User) (err error) {
func (user *User) UpdateUser(u map[string]interface{}) (err error) {
	fmt.Println("===========")
	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", u)
	fmt.Println("===========")
	Db := initialize.DB.Model(user).Select("name", "age", "is_admin").Updates(u)
	err = Db.Error
	if err != nil {
		fmt.Println("更新数据错误,err:", err)
		return err
	}
	return nil
}

// DeleteUser 删除一条数据
func (user *User) DeleteUser() (err error) {
	Db := initialize.DB.Table("users").Delete(user)
	err = Db.Error
	if err != nil {
		fmt.Println("删除数据错误,err:", err)
		return err
	}
	return nil
}

// GetAdminUser 获取一条数据
func (user *User) GetAdminUser(id int) (err error) {
	Db := initialize.DB.Table("admin_user").First(user, id)
	err = Db.Error
	if err != nil {
		fmt.Println("查询数据错误,err:", err)
		return err
	}
	return nil
}

// GetListAdminUser 获取列表数据
func (user *User) GetListAdminUser(users *[]User, page, size int) (err error) {
	Db := initialize.DB.Table("admin_user").Model(model.User{}).Offset((page - 1) * size).Limit(size).Find(&users)
	err = Db.Error
	if err != nil {
		fmt.Println("列表数据错误,err:", err)
		return err
	}
	return nil
}
