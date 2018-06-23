package main

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	type User struct {
		gorm.Model
		Name   string
		Active bool
		State  string
	}

	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&User{}) //自动创建表。
	// 初始化admin 还有其他的，比如API
	Admin := admin.New(&admin.AdminConfig{DB: DB})

	user := Admin.AddResource(&User{})
	user.Action(&admin.Action{
		Name:  "enable",
		Label: "启/禁用",
		Handler: func(actionArgument *admin.ActionArgument) error {
			// `FindSelectedRecords` => in bulk action mode, will return all checked records, in other mode, will return current record
			for _, record := range actionArgument.FindSelectedRecords() {
				actionArgument.Context.DB.Model(record.(*User)).Update("Active", true)
			}
			return nil
		},
		Modes: []string{"batch", "edit", "show"},
	})

	user.Action(&admin.Action{
		Name:  "url",
		Label: "地址",
		URL: func(record interface{}, context *admin.Context) string {
			if user, ok := record.(*User); ok {
				return fmt.Sprintf("/admin/users/%v.json", user.ID)
			}
			return "#"
		},
		//用啥方法也不能新窗口打开。
		//URLOpenType: "_blank\" onclick=\"function(event){event.stopPropagation();}\" target=\"_blank",
		Modes:       []string{"menu_item", "edit", "show"},
	})

	user.Action(&admin.Action{
		Name:  "preview",
		Label: "预览",
		URL: func(record interface{}, context *admin.Context) string {
			if user, ok := record.(*User); ok {
				return fmt.Sprintf("/admin/users/%v.json", user.ID)
			}
			return "#"
		},
		URLOpenType: "bottomsheet",
		Resource:    user,
		Modes:       []string{"menu_item", "edit", "show"},
	})

	// 发布按钮，显示到右侧上面。
	user.Action(&admin.Action{
		Name:  "publish",
		Label: "发布",
		Handler: func(actionArgument *admin.ActionArgument) error {
			logs.Info("############### publish ###############")
			return nil
		},
		Modes: []string{"show", "collection"},
	})

	// 启动服务
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)
}
