package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/admin"
	"net/http"
	"database/sql"
	"fmt"
)

//声明全局变量
var (
	// 注册数据库，可以是sqlite3 也可以是 mysql 换下驱动就可以了。
	DB, _ = gorm.Open("sqlite3", "demo.db")
	// 初始化admin 还有其他的，比如API
	Admin = admin.New(&admin.AdminConfig{DB: DB, SiteName: "Qor Admin "})
)

func init() {
	//地址
	type Address struct {
		gorm.Model
		Code string
		Name string
	}
	//角色
	type Role struct {
		gorm.Model
		Name string
	}
	//用户
	type User struct {
		gorm.Model
		Email     string
		Password  string
		Name      sql.NullString
		Gender    string
		Role      Role
		Addresses []Address
		Active     bool
	}
	DB.AutoMigrate(&Address{}, &Role{}, &User{}) //自动创建表。

	address := Admin.AddResource(&Address{}, &admin.Config{Name: "地址管理", Menu: []string{"组织管理"}})
	role := Admin.AddResource(&Role{}, &admin.Config{Name: "角色管理", Menu: []string{"组织管理"}})
	// Add it to Admin
	user := Admin.AddResource(&User{}, &admin.Config{Name: "用户管理", Menu: []string{"组织管理"}})
	// 显示字段
	//user.IndexAttrs("Name", "Gender", "Role", "Password", "CreateAt")
	// 不显示字段
	user.IndexAttrs("-Password", "-Role", "-Addresses")
	user.Meta(&admin.Meta{Name: "Email", Label: "邮箱"})
	user.Meta(&admin.Meta{Name: "Name", Label: "名称"})
	user.Meta(&admin.Meta{Name: "Gender", Label: "性别", Config: &admin.SelectOneConfig{Collection: []string{"Male", "Female", "Unknown"}}})
	////分组显示
	//user.Scope(&admin.Scope{Name: "Male", Group: "Gender", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
	//	return db.Where("gender = ?", "Male")
	//}})
	//user.Scope(&admin.Scope{Name: "Female", Group: "Gender", Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
	//	return db.Where("gender = ?", "Female")
	//}})

	user.Action(&admin.Action{
		Name: "enable",
		Handler: func(actionArgument *admin.ActionArgument) error {
			// `FindSelectedRecords` => in bulk action mode, will return all checked records, in other mode, will return current record
			for _, record := range actionArgument.FindSelectedRecords() {
				actionArgument.Context.DB.Model(record.(*User)).Update("Active", true)
			}
			return nil
		},
	})
	// Filter users by gender
	user.Filter(&admin.Filter{
		Name: "性别",
		Config: &admin.SelectOneConfig{
			Collection: []string{"Male", "Female", "Unknown"},
		},
	})
	// Filter products by collection
	user.Filter(&admin.Filter{
		Name:   "角色",
		Config: &admin.SelectOneConfig{RemoteDataResource: role},
	})

	fmt.Println(user)
	fmt.Println(role)
	fmt.Println(address)
}

func main() {
	// 启动服务
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	http.ListenAndServe(":9000", mux)
}
