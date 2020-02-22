package controllers

import (
	"mbook/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	cateModel := new(models.Category)
	if cateList, err := cateModel.GetCateGoryList(-1, 1); err == nil {
		c.Data["Cates"] = cateList
	} else {
		beego.Error(err.Error())
	}
	c.TplName = "home/list.html"
}
