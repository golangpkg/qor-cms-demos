package main

import (
	"fmt"
	"net/http"
	"github.com/qor/admin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//windows需要下载 http://tdm-gcc.tdragon.net/download
)

// 用户
type User struct {
	gorm.Model
	Name string
}

// 产品
type Product struct {
	gorm.Model
	Name        string
	Description string
}

func main() {
	// 注册数据库，可以是sqlite3 也可以是 mysql 换下驱动就可以了。
	//mysql_url := "root:mariadb@mariadb/qor_cms?charset=utf8&parseTime=True&loc=Local"
	//DB, _ := gorm.Open("mysql", mysql_url)
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&User{}, &Product{}) //自动创建表。

	// 初始化admin 还有其他的，比如API
	Admin := admin.New(&admin.AdminConfig{DB: DB})

	// 创建admin后台对象资源。
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})

	// 启动服务
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)
}
