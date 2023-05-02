package rbac

type RpcMenusVo struct {
	Id       int          `json:"id"`
	MenuCode string       `json:"menuCode"`
	MenuName string       `json:"menuName"`
	MenuIcon string       `json:"menuIcon"`
	MenuPath string       `json:"menuPath"`
	MenuSort int          `json:"menuSort"`
	IFrame   bool         `json:"iFrame"`
	Children []RpcMenusVo `json:"children"`
}
