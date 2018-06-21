package main

import (
	"fmt"
	"net/http"
	"github.com/qor/admin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/yaml"
	"path/filepath"
	//"github.com/qor/l10n"
	"github.com/qor/l10n"
	"os"
)

// 用户
type User struct {
	gorm.Model
	Name string
}

func main() {
	// 注册数据库，可以是sqlite3 也可以是 mysql 换下驱动就可以了。
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&User{}, ) //自动创建表。

	i18nPath := os.Getenv("GOPATH") + "/src/github.com/golangpkg/qor-cms-demos/demo06-i18n"
	println("i18nPath: ", i18nPath)
	I18n := i18n.New(
		//database.New(DB),                                       // load translations from the database
		yaml.New(filepath.Join(i18nPath, "config/locales")), // load translations from the YAML files in directory `config/locales`
	)
	I18n.SaveTranslation(&i18n.Translation{Key: "qor_i18n.form.saved", Locale: "en-US", Value: "保存"})

	l10n.Global = "zh-CN"
	I18n.T("en-US", "demo.greeting") // Not exist at first
	I18n.T("en-US", "demo.hello")    // Exists in the yml file
	fmt.Println(I18n)

	// 初始化admin 还有其他的，比如API
	Admin := admin.New(&admin.AdminConfig{DB: DB, I18n: I18n})

	// 创建admin后台对象资源。
	Admin.AddResource(&User{})
	Admin.AddResource(I18n)

	// 启动服务
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)
}
