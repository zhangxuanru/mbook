/*
@Time : 2020/2/17 18:57
@Author : zxr
@File : category
@Software: GoLand
*/
package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id     int
	Pid    int    //分类id
	Title  string `orm:"size(30);unique"`
	Intro  string //介绍
	Icon   string
	Cnt    int  //统计分类下图书
	Sort   int  //排序
	Status bool //状态，true 显示
}

func (m *Category) TableName() string {
	return TNCategory()
}

func (m *Category) GetCateGoryList(pid int, status int) (catList []Category, err error) {
	qs := orm.NewOrm().QueryTable(m.TableName())
	if pid > -1 {
		qs.Filter("pid", pid)
	}
	if status == 0 || status == 1 {
		qs.Filter("status", status)
	}
	_, err = qs.OrderBy("-status", "sort", "title").All(&catList)
	return
}

func (m *Category) Find(cid int) (cate Category) {
	cate.Id = cid
	orm.NewOrm().Read(&cate)
	return
}
