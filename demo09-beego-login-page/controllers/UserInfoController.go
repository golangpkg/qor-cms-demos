package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"github.com/qor/session/manager"
)

type UserInfoController struct {
	beego.Controller
}

var (
	adminPassword     = beego.AppConfig.String("adminPassword")
	USER_SESSION_NAME = "userSession"
)

//登录
func (c *UserInfoController) LoginIndex() {
	err := c.GetString("err", "")
	if err != "" {
		c.Data["Err"] = true
	} else {
		c.Data["Err"] = false
	}
	c.TplName = "login.html"
}

//登录
func (c *UserInfoController) Login() {

	userName := c.GetString("UserName", "")
	password := c.GetString("Password", "")
	//获得sessionid。
	if c.CruSession == nil {
		c.StartSession()
	}

	sessionId := c.CruSession.SessionID()
	logs.Info("sessionId %s get userName %s and password %s", sessionId, userName, password)
	if userName == "" {
		userName = "no_user_name"
	}
	c.SetSession("UserName", userName)

	if userName != "admin" || password != adminPassword {
		logs.Info("##################### login Error #####################")
		c.Ctx.Redirect(302, "/auth/login?err=password")
		return
	}
	//将用户对象放到session里面。
	//c.SetSession(USER_SESSION_NAME, userName)
	manager.SessionManager.Add(c.Ctx.ResponseWriter, c.Ctx.Request, USER_SESSION_NAME, userName)
	c.Ctx.Redirect(302, "/admin")
	return
}

//登录
func (c *UserInfoController) Logout() {
	//获得id
	//设置返回对象。
	if c.CruSession == nil {
		c.StartSession()
	}
	sessionId := c.CruSession.SessionID()
	logs.Info("==sessionId %s ==", sessionId)
	//设置 SessionDomain 名称。
	c.DestroySession()
	//设置返回对象。
	c.Ctx.Redirect(302, "/auth/login")
	return
}
