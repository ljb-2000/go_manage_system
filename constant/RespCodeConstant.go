package constant

import "git.gumpcome.com/go_kit/logiccode"

// 响应业务码

// 通用 0-9
var RESP_CODE_PARAMS_ERROR = logiccode.New(4000, "参数错误")
var RESP_CODE_PARAMS_VALUE_ERROR = logiccode.New(4001, "参数不合法")

// 登录 login 10-19
var RESP_CODE_ADMIN_LOGIN_ERROR = logiccode.New(4010, "管理员账户或密码错误")
var RESP_CODE_LOGIN_ERROR = logiccode.New(4011, "用户账户或密码错误")

// 权限 access 20-29
var RESP_CODE_ACCESS_ADD_ERROR = logiccode.New(4020, "权限添加失败")
var RESP_CODE_ACCESS_DELETE_ERROR = logiccode.New(4021, "权限删除失败")
var RESP_CODE_ACCESS_MENUS_ERROR = logiccode.New(4022, "菜单生成失败")
var RESP_CODE_ACCESS_PAGE_ERROR = logiccode.New(4023, "权限分页查询失败")

// 角色 role 30-39
var RESP_CODE_ROLE_ADD_ERROR = logiccode.New(4030, "角色添加失败")
var RESP_CODE_ROLE_ACCESS_BIND_ERROR = logiccode.New(4031, "角色绑定权限失败")
var RESP_CODE_ROLE_ACCESS_UNBIND_ERROR = logiccode.New(4032, "角色解除权限失败")
var RESP_CODE_ROLE_DELETE_ERROR = logiccode.New(4033, "角色删除失败")
var RESP_CODE_ROLE_PAGE_ERROR = logiccode.New(4034, "角色分页查询失败")


// 账号 account 40-49
var RESP_CODE_ACCOUNT_ADD_ERROR = logiccode.New(4040, "账号添加失败")
var RESP_CODE_ACCOUNT_DELETE_ERROR = logiccode.New(4041, "账号删除失败")
var RESP_CODE_ACCOUNT_PAGE_ERROR = logiccode.New(4042, "账号分页查询失败")
var RESP_CODE_ACCOUNT_ACCESSES_ERROR = logiccode.New(4043, "账号权限查询失败")


