package rbac

type RpcDeptAndUserTreeVo struct {
	Id        int                    `json:"id"`        //ID
	Name      string                 `json:"name"`      // 名称
	Type      int                    `json:"type"`      // 类型（0-部门，1-用户）
	JobName   string                 `json:"jobName"`   // 职务
	Telephone string                 `json:"telephone"` // 手机号
	Children  []RpcDeptAndUserTreeVo `json:"children"`  // 子节点
}
