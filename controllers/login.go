package controllers

import (
	"fmt"

	"github.com/qi/apimanager/models"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	fmt.Println("登录页面")
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	c.Ctx.Request.ParseForm()
	name := c.Ctx.Request.FormValue("username")
	user := &models.User{}
	user.Name = name
	c.SaveUser(user)
	c.Redirect(c.URLFor("MainController.Get"), 302)
}
