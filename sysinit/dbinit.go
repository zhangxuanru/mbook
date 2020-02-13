/*
@Time : 2020/2/13 18:52
@Author : zxr
@File : dbinit
@Software: GoLand
*/
package sysinit

import (
	_ "github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var defaultAlias = "default"

func dbInit(alias string) {
	dbAlias := alias //default
	if alias == "w" || alias == defaultAlias || len(alias) <= 0 {
		dbAlias = defaultAlias
		alias = "w"
	}
}
