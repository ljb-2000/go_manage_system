package constant

import "git.gumpcome.com/go_kit/logiccode"

// 响应业务码

// 通用 0-9
var RESP_CODE_PARAMS_ERROR = logiccode.New(4000, "参数错误")
var RESP_CODE_PARAMS_VALUE_ERROR = logiccode.New(4001, "参数不合法")

// 账号 account 10-19
var RESP_CODE_ACCOUNT_ERROR = logiccode.New(4010, "账户或密码错误")

// 权限 access 20-29
var RESP_CODE_ACCESS_ADD_ERROR = logiccode.New(4020, "权限添加失败")
var RESP_CODE_ACCESS_DELETE_ERROR = logiccode.New(4021, "权限删除失败")
var RESP_CODE_ACCESS_LIST_ERROR = logiccode.New(4021, "权限菜单生成失败")


