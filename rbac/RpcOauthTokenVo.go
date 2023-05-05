package rbac

type RpcTokenVo struct {
	AccessToken  string            `json:"accessToken"`  // 登陆Token
	RefreshToken string            `json:"refreshToken"` // 刷新Token
	TokenType    string            `json:"tokenType"`    // Token类型
	ExpiresIn    int               `json:"expiresIn"`    // Token过期时间
	Scope        string            `json:"scope"`        // 范围
	UserType     string            `json:"userType"`     // 用户类型
	UserInfo     RpcLoginTokenUser `json:"userInfo"`     // 用户信息
	ClientId     int               `json:"clientId"`     // 客户端ID
	Authorities  []string          `json:"authorities"`  // 功能权限集合
}

type RpcLoginTokenVo struct {
	AccessToken  string            `json:"accessToken"`  // 登陆Token
	RefreshToken string            `json:"refreshToken"` // 刷新Token
	TokenType    string            `json:"tokenType"`    // Token类型
	ExpiresIn    int               `json:"expiresIn"`    // Token过期时间
	Scope        string            `json:"scope"`        // 范围
	UserType     string            `json:"userType"`     // 用户类型
	UserInfo     RpcLoginTokenUser `json:"userInfo"`     // 用户信息
}
type RpcLoginTokenUser struct {
	UserId   int    `json:"userId"`   // 用户ID
	RealName string `json:"realName"` // 用户姓名
}
