package rbac

type RpcUserDetailVo struct {
	UserId       int           `json:"userId"`       // 用户ID
	RealName     string        `json:"realName"`     // 用户姓名
	Phone        string        `json:"phone"`        // 手机号
	SuperManager bool          `json:"superManager"` // 是否超级管理员
	UserType     string        `json:"userType"`     // 用户类型
	JobName      string        `json:"jobName"`      // 职务
	WxNickName   string        `json:"wxNickName"`   // 微信昵称
	WxAvatarUrl  string        `json:"wxAvatarUrl"`  // 微信头像地址
	Roles        []RpcUserRole `json:"roles"`        // 所属角色集合
	Depts        []RpcUserDept `json:"depts"`        // 所属部门集合
}
type RpcUserRole struct {
	RoleId   int    `json:"roleId"`   // 角色ID
	RoleCode string `json:"roleCode"` // 角色编码
	RoleName string `json:"roleName"` // 角色名称
}
type RpcUserDept struct {
	DeptId   int    `json:"deptId"`   // 部门ID
	DeptCode string `json:"deptCode"` // 部门编码
	DeptName string `json:"deptName"` // 部门名称
	DeptPid  int    `json:"deptPid"`  // 部门父ID
}
