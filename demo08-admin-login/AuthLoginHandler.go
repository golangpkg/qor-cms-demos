package main

import (
	"github.com/qor/auth"
	"github.com/qor/auth/claims"
	"strings"
	"github.com/astaxie/beego/logs"
	"github.com/qor/session/manager"
)

// 使用写死的用户名密码登陆。
var DefaultLoginHandler = func(context *auth.Context) (*claims.Claims, error) {
	var (
		req       = context.Request
		w         = context.Writer
		adminUser AdminUser
	)
	req.ParseForm()
	loginName := strings.TrimSpace(req.Form.Get("login"))
	password := strings.TrimSpace(req.Form.Get("password"))
	adminUser.UserName = loginName
	logs.Info("get DefaultAuthorizeHandler , loginName: %v , password : %v  ", loginName, password)

	if loginName == "admin" && password == "admin" {
		claims := adminUser.ToClaims()
		//登陆成功，注册session。
		manager.SessionManager.Add(w, req, "AdminUser", adminUser.UserName)
		logs.Info("##################### login success #####################")
		return claims, nil
	} else {
		return nil, auth.ErrInvalidPassword
	}
}


var DefaultRegisterHandler = func(context *auth.Context) (*claims.Claims, error) {
	return nil, auth.ErrInvalidAccount
}
