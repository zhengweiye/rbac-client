package rbac

type RpcMenusVo struct {
	Id       int          `json:"id"`       // 菜单ID
	MenuCode string       `json:"menuCode"` // 菜单编码
	MenuName string       `json:"menuName"` // 菜单名称
	MenuIcon string       `json:"menuIcon"` // 菜单图标
	MenuPath string       `json:"menuPath"` // 菜单路由
	MenuSort int          `json:"menuSort"` // 菜单序号
	IFrame   bool         `json:"iFrame"`   // 是否外部连接
	Children []RpcMenusVo `json:"children"` // 子菜单
}
