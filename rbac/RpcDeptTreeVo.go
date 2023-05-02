package rbac

type RpcDeptTreeVo struct {
	Id       int             `json:"id"`
	Code     string          `json:"code"`
	Name     string          `json:"name"`
	Nocheck  bool            `json:"nocheck"`
	Children []RpcDeptTreeVo `json:"children"`
}
