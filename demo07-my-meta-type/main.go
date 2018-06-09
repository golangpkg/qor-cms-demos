package main

import (
	"fmt"
	"net/http"
	"github.com/qor/admin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//windows需要下载 http://tdm-gcc.tdragon.net/download
	"time"
	"github.com/qor/slug"
)

// 用户
type Product struct {
	gorm.Model
	Name         string
	Password     string
	Description  string
	Description2 string
	Year         string
	ReleaseDate  time.Time
	NameWithSlug slug.Slug
}

func main() {
	// 注册数据库，可以是sqlite3 也可以是 mysql 换下驱动就可以了。
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&Product{}, ) //自动创建表。

	// 初始化admin 还有其他的，比如API
	Admin := admin.New(&admin.AdminConfig{SiteName: "demo", DB: DB})

	// 创建admin后台对象资源。
	product := Admin.AddResource(&Product{})
	//属性配置：https://doc.getqor.com/admin/fields.html
	product.Meta(&admin.Meta{Name: "ReleaseDate", Type: "date"})
	product.Meta(&admin.Meta{Name: "Password", Type: "password"})
	//product.Meta(&admin.Meta{Name: "Description", Type: "rich_editor"})

	//https://doc.getqor.com/admin/extend_admin.html#create-new-meta-types
	//自定义组件：
	product.Meta(&admin.Meta{Name: "Description2", Type: "kindeditor"})

	fmt.Println(product)

	// 启动服务
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)
}
