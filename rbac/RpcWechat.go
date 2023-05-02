package rbac

type RpcWechatAccessTokenVo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Unionid      string `json:"unionid"`
	Scope        string `json:"scope"`
}

type RpcWechatAccessToken2Vo struct {
	RpcWechatAccessTokenVo
	SessionKey string `json:"session_key"`
}
