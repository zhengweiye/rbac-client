package rbac

type RpcDeptVo struct {
	DeptId       int    `json:"deptId"`       // 部门ID
	DeptPid      int    `json:"deptPid"`      // 父部门ID
	DeptName     string `json:"deptName"`     // 部门名称
	HeadUserId   int    `json:"headUserId"`   // 部门负责人ID
	HeadUserName string `json:"headUserName"` // 部门负责人姓名
	ChildrenIds  string `json:"childrenIds"`  // 子部门ID
	DeptType     string `json:"deptType"`     // 部门类型
	Enabled      bool   `json:"enabled"`      // 是否启用
}
