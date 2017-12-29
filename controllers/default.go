package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Ctx.Output.Body([]byte("欢迎你，" + c.User.Name))
}
