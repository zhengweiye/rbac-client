package rbac

type RpcDeptVo struct {
	DeptId       int    `json:"deptId"`
	DeptPid      int    `json:"deptPid"`
	DeptName     string `json:"deptName"`
	HeadUserId   int    `json:"headUserId"`
	HeadUserName string `json:"headUserName"`
	ChildrenIds  string `json:"childrenIds"`
	DeptType     string `json:"deptType"`
	Enabled      bool   `json:"enabled"`
}
