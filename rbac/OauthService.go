package rbac

import (
	"encoding/json"
	"fmt"
)

type OauthService interface {
	LoginByPassword(clientId, username, password string) RpcLoginTokenVo
	LoginBySms(clientId, phone, code string) RpcLoginTokenVo
	LoginByWechatPc(clientId, openId, phone, smscode string) RpcLoginTokenVo
	LoginByWechatApp(clientId, openId, phone, nickName, avatarUrl string) RpcLoginTokenVo
	RefreshToken(clientId, refreshToken string) RpcLoginTokenVo
	CheckToken(token string) *RpcTokenVo
	ClearToken(token string)
}
type OauthServiceImpl struct {
}

func NewOauthService() OauthService {
	return OauthServiceImpl{}
}

func (o OauthServiceImpl) LoginByPassword(clientId, username, password string) RpcLoginTokenVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/oauth/login")

	var param = make(map[string]any)
	param["client_id"] = clientId
	param["client_secret"] = GetClientSecret(clientId)
	param["grant_type"] = "password"
	param["username"] = username
	param["password"] = password

	resBytes := HttpPost(url, param)

	var result RpcLoginTokenVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (o OauthServiceImpl) LoginBySms(clientId, phone, code string) RpcLoginTokenVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/oauth/login")

	var param = make(map[string]any)
	param["client_id"] = clientId
	param["client_secret"] = GetClientSecret(clientId)
	param["grant_type"] = "sms"
	param["phone"] = phone
	param["code"] = code

	resBytes := HttpPost(url, param)

	var result RpcLoginTokenVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (o OauthServiceImpl) LoginByWechatPc(clientId, openId, phone, smscode string) RpcLoginTokenVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/oauth/login")

	var param = make(map[string]any)
	param["client_id"] = clientId
	param["client_secret"] = GetClientSecret(clientId)
	param["grant_type"] = "wechat_pc"
	param["openId"] = openId
	param["phone"] = phone
	param["code"] = smscode

	resBytes := HttpPost(url, param)

	var result RpcLoginTokenVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (o OauthServiceImpl) LoginByWechatApp(clientId, openId, phone, nickName, avatarUrl string) RpcLoginTokenVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/oauth/login")

	var param = make(map[string]any)
	param["client_id"] = clientId
	param["client_secret"] = GetClientSecret(clientId)
	param["grant_type"] = "wechat_app"
	param["phone"] = phone
	param["openId"] = openId
	param["nickName"] = nickName
	param["avatarUrl"] = avatarUrl

	resBytes := HttpPost(url, param)

	var result RpcLoginTokenVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (o OauthServiceImpl) RefreshToken(clientId, refreshToken string) RpcLoginTokenVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/oauth/refreshToken")

	var param = make(map[string]any)
	param["client_id"] = clientId
	param["client_secret"] = GetClientSecret(clientId)
	param["grant_type"] = "refresh_token"
	param["refresh_token"] = refreshToken

	resBytes := HttpPost(url, param)

	var result RpcLoginTokenVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (o OauthServiceImpl) CheckToken(token string) *RpcTokenVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/oauth/checkToken")

	var param = make(map[string]any)
	param["token"] = token

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}

	var result RpcTokenVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return &result
}

func (o OauthServiceImpl) ClearToken(token string) {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/oauth/clearToken")

	var param = make(map[string]any)
	param["token"] = token

	HttpPost(url, param)
}
