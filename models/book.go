/*
@Time : 2020/2/17 18:57
@Author : zxr
@File : book
@Software: GoLand
*/
package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Book struct {
	BookId         int       `orm:"pk;auto" json:"book_id"`
	BookName       string    `orm:"size(500)" json:"book_name"`       //名称
	Identify       string    `orm:"size(100);unique" json:"identify"` //唯一标识
	OrderIndex     int       `orm:"default(0)" json:"order_index"`
	Description    string    `orm:"size(1000)" json:"description"`       //图书描述
	Cover          string    `orm:"size(1000)" json:"cover"`             //封面地址
	Editor         string    `orm:"size(50)" json:"editor"`              //编辑器类型: "markdown"
	Status         int       `orm:"default(0)" json:"status"`            //状态:0 正常 ; 1 已删除
	PrivatelyOwned int       `orm:"default(0)" json:"privately_owned"`   // 是否私有: 0 公开 ; 1 私有
	PrivateToken   string    `orm:"size(500);null" json:"private_token"` // 私有图书访问Token
	MemberId       int       `orm:"size(100)" json:"member_id"`
	CreateTime     time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //创建时间
	ModifyTime     time.Time `orm:"type(datetime);auto_now_add" json:"modify_time"`
	ReleaseTime    time.Time `orm:"type(datetime);" json:"release_time"` //发布时间
	DocCount       int       `json:"doc_count"`                          //文档数量
	CommentCount   int       `orm:"type(int)" json:"comment_count"`
	Vcnt           int       `orm:"default(0)" json:"vcnt"`              //阅读次数
	Collection     int       `orm:"column(star);default(0)" json:"star"` //收藏次数
	Score          int       `orm:"default(40)" json:"score"`            //评分
	CntScore       int       //评分人数
	CntComment     int       //评论人数
	Author         string    `orm:"size(50)"`                      //来源
	AuthorURL      string    `orm:"column(author_url);size(1000)"` //来源链接
}

func (m *Book) TableName() string {
	return TNBook()
}

func NewBook() *Book {
	return &Book{}
}

func (m *Book) HomeData(page int, size int, cid int, fields ...string) (books []Book, total int, err error) {
	sqlFmt := "SELECT %v FROM %s b LEFT JOIN %s c on b.book_id=c.book_id where c.category_id=%d"
	if len(fields) == 0 {
		fields = append(fields, "book_id", "book_name", "identify", "conver", "order_index")
	}
	fieldStr := "b." + strings.Join(fields, ",b.")
	sql := fmt.Sprintf(sqlFmt, fieldStr, m.TableName(), TNBookCategory(), cid)
	totalStr := "COUNT(*) as cnt"
	totalSql := fmt.Sprint(sqlFmt, totalStr, m.TableName(), TNBookCategory(), cid)

	o := orm.NewOrm()
	var params []orm.Params
	if _, err := o.Raw(totalSql).Values(&params); err == nil && len(params) > 0 {
		total, _ = strconv.Atoi(params[0]["cnt"].(string))
	}
	_, err = o.Raw(sql).QueryRows(&books)
	return
}
