package main

import (
	"github.com/qor/auth"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/providers/password"
	//"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	//"github.com/qor/auth/providers/facebook"
	//"github.com/qor/auth/providers/twitter"
	"github.com/qor/session/manager"
	"github.com/jinzhu/gorm"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/admin"
)

//声明全局变量
var (
	// 注册数据库，可以是sqlite3 也可以是 mysql 换下驱动就可以了。
	DB, _ = gorm.Open("sqlite3", "demo.db")

	// 初始化admin 还有其他的，比如API
	Auth  = auth.New(&auth.Config{DB: DB})
	Admin = admin.New(&admin.AdminConfig{DB: DB})
)

func init() {

	DB.AutoMigrate(&auth_identity.AuthIdentity{}, &auth_identity.AuthIdentity{}, ) //自动创建表。
	// Register Auth providers
	// Allow use username/password
	Auth.RegisterProvider(password.New(&password.Config{}))

	// Allow use Google
	Auth.RegisterProvider(google.New(&google.Config{
		ClientID:     "google client id",
		ClientSecret: "google client secret",
	}))

	// 创建admin后台对象资源。
	Admin.AddResource(&auth_identity.AuthIdentity{})
}

func main() {
	// 启动服务
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	mux.Handle("/auth/", Auth.NewServeMux())
	http.ListenAndServe(":9000", manager.SessionManager.Middleware(mux))
	//访问 http://localhost:9000/auth/register 注册账号
	// admin ： http://localhost:9000/admin/auth_identities
}
