package dao

import (
	"git.fenggese.com/go_kit/dbkit"
	"git.fenggese.com/go_kit/logiccode"
	"git.fenggese.com/go_kit/strkit"
	"git.fenggese.com/go_manage_system/constant"
	"git.fenggese.com/go_manage_system/models"
	"github.com/astaxie/beego"
)

// @Title 添加角色
// @return 是否成功(bool)、插入后的ID(int64)、报错信息(error)
func AddRoleDao(params *map[string]interface{}) (bool, int64, error) {

	// 1. 获得连接
	db, err := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	if err != nil {
		return false, 0, err
	}

	// 2. 保存数据
	return dbkit.SaveInMysql(db, constant.ROLE_TABLE, *params)
}

// @Title 绑定权限
// @return 是否成功(bool)、插入后的ID(int64)、报错信息(error)
func BindAccessDao(params *map[string]interface{}) (bool, int64, error) {

	// 1. 获得连接
	db, err := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	if err != nil {
		return false, 0, err
	}

	// 2. 保存数据
	return dbkit.SaveInMysql(db, constant.ROLE_ACCESS_TABLE, *params)
}

// @Title 解除角色权限
// @return 是否成功(bool)、插入后的ID(int64)、报错信息(error)
func UnbindAccessDao(params *map[string]interface{}) (bool, error) {

	// 1. 获得连接
	db, err := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	if err != nil {
		return false, err
	}

	// 2. 生成删除的SQL语句
	sql, data := dbkit.CreateDeleteMysqlSQL(constant.ROLE_ACCESS_TABLE, *params)

	// 3. 删除数据
	return dbkit.DeleteInMysql(db, sql, data...)
}

// @Title 角色分页查找
// @return 分页结果(dbkitPage)、报错信息(error)
func PageRoleDao(filter *models.PageRoleFilter) (dbkit.Page, error) {

	// 1. 查询语句
	selectSql := `SELECT id, name, create_time`

	// 2. 条件语句
	data := make([]interface{}, 0)
	sqlExceptSelect := strkit.StringBuilder{}
	sqlExceptSelect.Append("FROM role WHERE 1=1")

	// 分页查找
	db, _ := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	return dbkit.PaginateInMysql(db, filter.PageNumber, filter.PageSize, selectSql, sqlExceptSelect.ToString(), []string{"id"}, data...)
}

// @Title 查询角色的权限
// @return 查询结果、报错信息(error)
func AccessQueryByRoleDao(role_id int64) ([]map[string]interface{}, error) {

	// 1. 获得连接
	db, err := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	if err != nil {
		return []map[string]interface{}{}, err
	}

	// 2. SQL语句
	sql := `SELECT access.id, access.code, access.name
	FROM access, role_access WHERE access.id = role_access.access_id
	AND role_access.role_id = ? LIMIT 100;`

	// 3. 查询
	return dbkit.FindInMysql(db, sql, []string{"id", "code"}, []interface{}{role_id}...)
}

// @Title 删除角色
// @Description 开启事务后，先删角色表，再删角色权限关联表
// @return 是否成功(bool)、报错信息(error)
func DeleteRoleDao(id int) (bool, error) {
	// 1. 数据库连接 和 SQL
	db, _ := dbkit.GetMysqlCon(constant.MYSQL_CFNAME)
	//删除角色
	sql := `DELETE FROM role WHERE id=?`
	beego.BeeLogger.Debug(sql)

	// 2. 开启事务
	tx, err := db.Begin() //开启事务
	if err != nil {
		beego.BeeLogger.Error("%v", err)
		return false, logiccode.DbDeleteErrorCode()
	}
	defer tx.Commit() //提交事务

	// 3. 执行删除角色的SQL
	stmt, err := tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		beego.BeeLogger.Error("%v", err)
		return false, logiccode.DbDeleteErrorCode()
	}
	defer stmt.Close() //释放连接

	result, err := stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		beego.BeeLogger.Error("%v", err)
		return false, logiccode.DbDeleteErrorCode()
	}

	rowsNum, _ := result.RowsAffected() //影响的总行数
	if rowsNum == 0 {
		return false, logiccode.DbDeleteErrorCode()
	}

	// 4. 只想删除角色权限关联表的SQL
	sql = `DELETE FROM role_access WHERE role_id=?`
	beego.BeeLogger.Debug(sql)

	stmt, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		beego.BeeLogger.Error("%v", err)
		return false, logiccode.DbInsertErrorCode()
	}

	// 5. 因为存在没有权限的角色，所以对角色关联表的删除结果不做处理。
	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		beego.BeeLogger.Error("%v", err)
		return false, logiccode.DbDeleteErrorCode()
	}

	return true, nil
}
