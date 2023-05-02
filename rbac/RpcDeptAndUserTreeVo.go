package rbac

type RpcDeptAndUserTreeVo struct {
	Id        int                    `json:"id"`
	Name      string                 `json:"name"`
	Type      int                    `json:"type"` // todo 0-部门，1-用户
	JobName   string                 `json:"jobName"`
	Telephone string                 `json:"telephone"`
	Children  []RpcDeptAndUserTreeVo `json:"children"`
}
