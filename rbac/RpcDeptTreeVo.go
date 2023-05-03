package rbac

type RpcDeptTreeVo struct {
	Id       int             `json:"id"`       // ID
	Code     string          `json:"code"`     // 部门编码
	Name     string          `json:"name"`     // 部门名称
	Nocheck  bool            `json:"nocheck"`  // 是否可勾选
	Children []RpcDeptTreeVo `json:"children"` // 子部门
}
