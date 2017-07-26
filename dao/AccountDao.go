package dao

import (
	"git.gumpcome.com/go_kit/dbkit"
	"git.gumpcome.com/gumpoa/constant"
	"git.gumpcome.com/go_kit/strkit"
	"git.gumpcome.com/gumpoa/models"
)

// @Title 添加账号
// @return 是否成功(bool)、插入后的ID(int64)、报错信息(error)
func AddAccountDao(params *map[string]interface{}) (bool, int64, error) {

	// 1. 获得连接
	db, err := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	if err != nil {
		return false, 0, err
	}

	// 2. 保存数据
	return dbkit.SaveInMysql(db, constant.ACCOUNT_TABLE, *params)
}



// @Title 删除账号
// @return 是否成功(bool)、插入后的ID(int64)、报错信息(error)
func DeleteAccountDao(params *map[string]interface{}) (bool, error) {

	// 1. 获得连接
	db, err := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	if err != nil {
		return false, err
	}

	// 2. 生成删除的SQL语句
	sql, data := dbkit.CreateDeleteMysqlSQL(constant.ACCOUNT_TABLE, *params)

	// 3. 删除数据
	return dbkit.DeleteInMysql(db, sql, data...)
}


// @Title 账号分页查找
// @return 分页结果(dbkitPage)、报错信息(error)
func PageAccountDao(filter *models.PageAccountFilter) (dbkit.Page, error) {

	// 1. 查询语句
	selectSql := `SELECT account.id, account.user_name, account.login_name, role.name AS role_name, account.create_time`

	// 2. 条件语句
	data := make([]interface{}, 0)
	sqlExceptSelect := strkit.StringBuilder{}
	sqlExceptSelect.Append("FROM account, role WHERE account.role_id=role.id")

	// 分页查找
	db, _ := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	return dbkit.PaginateInMysql(db, filter.PageNumber, filter.PageSize, selectSql, sqlExceptSelect.ToString(), []string{"account.id"}, data...)
}


// @Title 根据账号信息查询数据库
func LoginDao(account *models.Account) (int, error) {
	// 1. 准备sql和数据
	sql := `SELECT count(*) AS count FROM account WHERE login_name=? AND login_pwd=?`
	db, _ := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	data := []interface{}{
		account.LoginName,
		account.LoginPwd,
	}
	var count int

	// 2. 查询数据库
	rows, err := db.Query(sql, data...)
	if err != nil {
		return 0, err
	}

	// 3. 处理结果集
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return count, err
}


// @Title 查询账号的权限
// @return 查询结果、报错信息(error)
func AccessQueryByAccountDao(login_name string) ([]map[string]interface{}, error) {

	// 1. 获得连接
	db, err := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	if err != nil {
		return []map[string]interface{}{}, err
	}

	// 2. SQL语句
	sql := `SELECT access.id, access.code, access.name
	FROM access, role_access, account WHERE account.login_name=? AND role_access.role_id=account.role_id
	AND access.id = role_access.access_id LIMIT 100;`

	// 3. 查询
	return dbkit.FindInMysql(db, sql, []string{"id", "code"}, []interface{}{login_name}...)
}