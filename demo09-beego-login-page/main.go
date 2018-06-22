package main

import (
	"github.com/astaxie/beego"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/golangpkg/qor-cms-demos/demo09-beego-login-page/models"
	"github.com/golangpkg/qor-cms-demos/demo09-beego-login-page/controllers"
	"github.com/astaxie/beego/logs"
	"github.com/qor/qor"
	"github.com/qor/session/manager"
)

// //########################## 定义admin 权限 ##########################
type AdminAuth struct {
}

func (AdminAuth) LoginURL(c *admin.Context) string {
	logs.Info(" user not login ")
	return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
	logs.Info(" user  logout ")
	return "/auth/logout"
}

//从session中获得当前用户。
func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	adminUserName := manager.SessionManager.Get(c.Request, controllers.USER_SESSION_NAME)
	logs.Info("########## adminUser %v", adminUserName)
	if adminUserName != "" {
		userInfo := models.UserInfo{}
		userInfo.UserName = adminUserName
		return &userInfo
	}
	return nil
}

func main() {
	//开启session。配置文件 配置下sessionon = true即可。
	beego.BConfig.WebConfig.Session.SessionOn = true
	// 注册数据库，可以是sqlite3 也可以是 mysql 换下驱动就可以了。
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&models.UserInfo{}) //自动创建表。
	// 初始化admin 还有其他的，比如API
	Admin := admin.New(&admin.AdminConfig{SiteName: "demo", DB: DB, Auth: AdminAuth{}})
	Admin.AddResource(&models.UserInfo{})
	// 启动服务
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	beego.Handler("/admin/*", mux)
	beego.Router("/auth/login", &controllers.UserInfoController{}, "get:LoginIndex")
	beego.Router("/auth/login", &controllers.UserInfoController{}, "post:Login")
	beego.Router("/auth/logout", &controllers.UserInfoController{}, "get:Logout")
	beego.Run()
}
