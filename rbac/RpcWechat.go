package rbac

type RpcWechatAccessToken2Vo struct {
	RpcWechatAccessTokenVo        // 微信信息
	SessionKey             string `json:"session_key"` // session_key
}

type RpcWechatAccessTokenVo struct {
	AccessToken  string `json:"access_token"`  // 微信token
	ExpiresIn    int64  `json:"expires_in"`    // token过期时间
	RefreshToken string `json:"refresh_token"` // 微信刷新token
	Openid       string `json:"openid"`        // 微信openid
	Unionid      string `json:"unionid"`       // 微信unionid
	Scope        string `json:"scope"`         // 范围
}
