package rbac

type RpcUserDetailVo struct {
	UserId       int           `json:"userId"`
	RealName     string        `json:"realName"`
	Phone        string        `json:"phone"`
	SuperManager bool          `json:"superManager"`
	UserType     string        `json:"userType"`
	JobName      string        `json:"jobName"`
	WxNickName   string        `json:"wxNickName"`
	WxAvatarUrl  string        `json:"wxAvatarUrl"`
	Roles        []RpcUserRole `json:"roles"`
	Depts        []RpcUserDept `json:"depts"`
}
type RpcUserRole struct {
	RoleId   int    `json:"roleId"`
	RoleCode string `json:"roleCode"`
	RoleName string `json:"roleName"`
}
type RpcUserDept struct {
	DeptId   int    `json:"deptId"`
	DeptCode string `json:"deptCode"`
	DeptName string `json:"deptName"`
	DeptPid  int    `json:"deptPid"`
}
