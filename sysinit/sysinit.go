package sysinit

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

func sysInit() {
	uploads := filepath.Join("./", "uploads")
	fmt.Println("uploads", uploads)
	beego.BConfig.WebConfig.StaticDir["/uploads"] = uploads
	registerFunctions()
}

func registerFunctions() {
	beego.AddFuncMap("cdnjs", func(p string) string {
		cdn := beego.AppConfig.DefaultString("cdnjs", "")
		if len(cdn) > 0 && strings.HasPrefix(p, "/") && strings.HasSuffix(cdn, "/") {
			return cdn + string(p[1:])
		}
		return p
	})
	beego.AddFuncMap("cdncss", func(p string) string {
		cdn := beego.AppConfig.DefaultString("cdncss", "")
		if strings.HasPrefix(p, "/") && strings.HasSuffix(cdn, "/") {
			return cdn + string(p[1:])
		}
		if !strings.HasPrefix(p, "/") && !strings.HasSuffix(cdn, "/") {
			return cdn + "/" + p
		}
		return cdn + p
	})
}
