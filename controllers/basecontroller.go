package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/qi/apimanager/models"
	_ "github.com/qi/session/memory"
	"github.com/qi/session/session"
)

var globalSessions *session.Manager

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

type BaseController struct {
	beego.Controller
	User models.User
}

func (c *BaseController) SaveUser(user *models.User) {
	sess := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	user.Id = sess.SessionID()
	sess.Set("user", user)
}

func (controller *BaseController) Prepare() {
	fmt.Println("处理请求之前")
	//不管用户是否登录SessionStart都会产生一个session.未登录用户只有sessionid和并未存储数据的session
	sess := globalSessions.SessionStart(controller.Ctx.ResponseWriter, controller.Ctx.Request)
	a := sess.Get("user")
	if a == nil || a.(*models.User).Id == "" {
		url := beego.URLFor("LoginController.Get")
		if controller.Ctx.Request.URL.Path == url {
			return
		}
		fmt.Println("重定向到登录页：", url)
		controller.Redirect(url, 302)
		return
	}
	controller.User = *a.(*models.User)
}
