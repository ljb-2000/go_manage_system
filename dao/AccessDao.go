package dao

import (
	"git.gumpcome.com/go_kit/dbkit"
	"git.gumpcome.com/gumpoa/constant"
)

const (
	TableName = "access"
)

// @Title 添加权限
// @return 是否成功(bool)、插入后的ID(int64)、报错信息(error)
func AddAccessDao(params *map[string]interface{}) (bool, int64, error) {

	// 1. 获得连接
	db, err := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	if err != nil {
		return false, 0, err
	}

	// 2. 保存数据
	return dbkit.SaveInMysql(db, TableName, *params)
}

// @Title 删除权限
// @return 是否成功(bool)、插入后的ID(int64)、报错信息(error)
func DeleteAccessDao(params *map[string]interface{}) (bool, error) {

	// 1. 获得连接
	db, err := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	if err != nil {
		return false, err
	}

	// 2. 生成删除的SQL语句
	sql, data := dbkit.CreateDeleteMysqlSQL(TableName, *params)

	// 3. 删除数据
	return dbkit.DeleteInMysql(db, sql, data...)
}
