package main

import (
	"encoding/json"
	"git.fenggese.com/go_kit/dbkit"
	"git.fenggese.com/go_kit/logkit"
	"git.fenggese.com/go_manage_system/constant"
	_ "git.fenggese.com/go_manage_system/routers"
	"github.com/astaxie/beego"
	"net/http"
)

func main() {

	// 初始化日志
	logkit.InitLog()

	// 初始化MySQL数据库连接信息
	dbUserName := beego.AppConfig.String("mysql_user")
	dbUserPwd := beego.AppConfig.String("mysql_pwd")
	dbHost := beego.AppConfig.String("mysql_host")
	dbName := beego.AppConfig.String("mysql_db_name")
	dbMaxIdle, _ := beego.AppConfig.Int("mysql_maxidle")
	dbMaxActive, _ := beego.AppConfig.Int("mysql_maxactive")
	dbCfgName := constant.MYSQL_CFNAME
	dbkit.InitMysql(dbUserName, dbUserPwd, dbHost, dbName, dbCfgName, dbMaxIdle, dbMaxActive)

	// 404
	beego.ErrorHandler("404", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"code": 404,
			"msg":  "not found",
		}
		result, _ := json.Marshal(data)
		w.Write(result)
	})

	beego.Run()
}
