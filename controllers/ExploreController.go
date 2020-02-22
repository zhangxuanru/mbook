package controllers

import (
	"math"
	"mbook/models"
	"mbook/utils"
	"strconv"

	"github.com/astaxie/beego"
)

type ExploreController struct {
	BaseController
}

func (c *ExploreController) Index() {
	var (
		cid       int
		cate      models.Category
		cateModel *models.Category
		urePrefix = beego.URLFor("ExploreController.Index")
	)
	cateModel = new(models.Category)
	if cid, _ = c.GetInt("cid"); cid > 0 {
		cate = cateModel.Find(cid)
	} else {
		c.Abort("404")
	}
	page, _ := c.GetInt("page", 1)
	pageSize := 24

	books, totalCount, err := models.NewBook().HomeData(page, pageSize, cid)
	if err != nil {
		beego.Error(err.Error())
		c.Abort("404")
	}
	c.Data["PageHtml"] = ""
	if totalCount > 0 {
		urlSuffix := "&cid=" + strconv.Itoa(cid)
		html := utils.NewPaginations(4, totalCount, pageSize, page, urePrefix, urlSuffix)
		c.Data["PageHtml"] = html
	}
	c.Data["TotalPages"] = int(math.Ceil(float64(totalCount) / float64(pageSize)))
	c.Data["Lists"] = books
	c.Data["Cate"] = cate
	c.Data["Cid"] = cid
	c.TplName = "explore/index.html"
}
