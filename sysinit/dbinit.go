/*
@Time : 2020/2/13 18:52
@Author : zxr
@File : dbinit
@Software: GoLand
*/
package sysinit

import (
	"fmt"
	"log"

	_ "mbook/models"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var defaultAlias = "default"

func dbInit(aliases ...string) {
	if len(aliases) > 0 {
		for _, alias := range aliases {
			registerDatabase(alias)
		}
	} else {
		registerDatabase("w")
	}
}

func registerDatabase(alias string) {
	if len(alias) == 0 {
		fmt.Println("alias init error")
		return
	}
	dbAlias := alias //default
	if alias == "w" || alias == defaultAlias {
		dbAlias = defaultAlias
		alias = "w"
	}
	ormConnectDb(dbAlias, alias)
}

func ormConnectDb(dbAlias string, alias string) {
	dbPrefix := "db"
	dbHost := beego.AppConfig.String(dbPrefix + "_" + alias + "_host")
	dbPort := beego.AppConfig.String(dbPrefix + "_" + alias + "_port")
	dbName := beego.AppConfig.String(dbPrefix + "_" + alias + "_database")
	dbUserName := beego.AppConfig.String(dbPrefix + "_" + alias + "_username")
	dbPassWord := beego.AppConfig.String(dbPrefix + "_" + alias + "_password")
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&timeout=5sreadTimeout=5s",
		dbUserName, dbPassWord, dbHost, dbPort, dbName)
	if err := orm.RegisterDataBase(dbAlias, "mysql", dbSource); err != nil {
		log.Println("conn err :", err)
		return
	}
	isDebug := beego.AppConfig.String("runmode") == "dev"
	if alias == "w" {
		orm.RunSyncdb(dbAlias, false, isDebug)
	}
	orm.Debug = isDebug
}
