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
	"github.com/qor/qor"
	"fmt"
	//"github.com/qor/auth_themes/clean"
	"github.com/qor/roles"
)

type AdminAuth struct{}

func (AdminAuth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	currentUser := Auth.GetCurrentUser(c.Request)
	if currentUser != nil {
		qorCurrentUser, ok := currentUser.(qor.CurrentUser)
		if !ok {
			fmt.Printf("User %#v haven't implement qor.CurrentUser interface\n", currentUser)
		}
		return qorCurrentUser
	}
	return nil
}

type User struct {
	Id       string
	Username string
	Password string
}

func (u *User) DisplayName() string {
	return u.Username
}

//声明全局变量
var (
	// 注册数据库，可以是sqlite3 也可以是 mysql 换下驱动就可以了。
	DB, _ = gorm.Open("sqlite3", "demo.db")
	// 初始化admin 还有其他的，比如API
	Auth = auth.New(&auth.Config{DB: DB, UserModel: User{},})
	// User model needs to implement qor.CurrentUser interface (https://godoc.org/github.com/qor/qor#CurrentUser) to use it in QOR Admin

	Admin = admin.New(&admin.AdminConfig{DB: DB, Auth: &AdminAuth{}})
)

func init() {

	DB.AutoMigrate(&auth_identity.AuthIdentity{},&User{}) //自动创建表。
	// Register Auth providers
	// Allow use username/password
	Auth.RegisterProvider(password.New(&password.Config{}))

	// Allow use Google
	Auth.RegisterProvider(google.New(&google.Config{
		ClientID:     "google client id",
		ClientSecret: "google client secret",
	}))

	// 创建admin后台对象资源。
	//Admin.AddResource(&auth_identity.AuthIdentity{})
	Admin.AddResource(&User{}, &admin.Config{
		Permission: roles.Deny(roles.Delete, roles.Anyone).Allow(roles.Delete, "admin"),
	})
	Admin.AddMenu(&admin.Menu{Name: "Report", Link: "/admin", Permission: roles.Allow(roles.Read, "admin")})

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
