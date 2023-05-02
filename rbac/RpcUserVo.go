package rbac

type RpcUserVo struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
	JobName  string `json:"jobName"`
}
