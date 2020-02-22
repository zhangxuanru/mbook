/*
@Time : 2020/2/17 18:57
@Author : zxr
@File : book_category
@Software: GoLand
*/
package models

//图书分类对应关系
type BookCategory struct {
	Id         int //自增主键
	BookId     int //书籍id
	CategoryId int //分类id
}

func (m *BookCategory) TableName() string {
	return TNBookCategory()
}
