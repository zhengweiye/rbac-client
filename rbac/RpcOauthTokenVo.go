package rbac

type RpcLoginTokenVo struct {
	AccessToken  string            `json:"accessToken"`
	RefreshToken string            `json:"refreshToken"`
	TokenType    string            `json:"tokenType"`
	ExpiresIn    int               `json:"expiresIn"`
	Scope        string            `json:"scope"`
	UserType     string            `json:"userType"`
	UserInfo     RpcLoginTokenUser `json:"userInfo"`
}
type RpcLoginTokenUser struct {
	UserId   int    `json:"userId"`
	RealName string `json:"realName"`
}

type RpcTokenVo struct {
	AccessToken RpcLoginTokenVo `json:"accessToken"`
	ClientId    int             `json:"clientId"`
	Authorities []string        `json:"authorities"`
}
