package rbac

type RpcLoginTokenVo struct {
	Error            string            `json:"error"`
	ErrorDescription string            `json:"error_description"`
	AccessToken      string            `json:"access_token"`
	TokenType        string            `json:"token_type"`
	RefreshToken     string            `json:"refresh_token"`
	ExpiresIn        int               `json:"expires_in"`
	Scope            string            `json:"scope"`
	UserInfo         RpcLoginTokenUser `json:"user_info"`
}
type RpcLoginTokenUser struct {
	UserId   int    `json:"userId"`
	RealName string `json:"realName"`
}

type RpcTokenVo struct {
	UserId      int      `json:"userId"`
	UserName    string   `json:"userName"`
	Principal   string   `json:"principal"`
	UserType    string   `json:"userType"`
	Authorities []string `json:"authorities"`
}
