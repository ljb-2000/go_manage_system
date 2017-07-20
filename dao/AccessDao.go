package dao

import (
	"git.gumpcome.com/gumpoa/models"
	"git.gumpcome.com/go_kit/dbkit"
	"fmt"
	"git.gumpcome.com/go_kit/logiccode"
)

func AddAccessDao(user *models.Access) (bool, error) {

	sql := `INSERT access SET code=?, name=?`

	// 2. 获得连接
	db, err := dbkit.GetMysqlCon("gumpoa")
	return false, logiccode.DbInsertErrorCode()

	// 3. 创建SQL语句
	stmt, err := db.Prepare(`INSERT access SET code=?, name=?`)
	checkErr(err)

	// 4. 执行SQL
	res, err := stmt.Exec("103", "无视我")
	checkErr(err)

	// 5. 判断结果
	rowCount, err := res.RowsAffected()
	checkErr(err)

	db.Close()
	fmt.Println(rowCount == 1)



}


func DeleteAccessDao(id int64) (bool, error) {

}

func UpdateAccessDao(data map[string]interface{}) (bool, error) {

}

func FindAccessDao(id int) (map[string]interface{}, error) {

}

//func PageAccessDao(filter *models.PageAccessFilter) (dbkit.Page, error) {
//
//}

