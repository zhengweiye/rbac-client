package rbac

type RpcUserVo struct {
	UserId   int    `json:"userId"`   // 用户ID
	UserName string `json:"userName"` // 用户姓名
	Phone    string `json:"phone"`    // 手机号
	JobName  string `json:"jobName"`  // 职务
}
