package rbac

import (
	"encoding/json"
	"fmt"
)

type WeChatService interface {
	GetAccessTokenOfApp(clientId, wechatCode string) *RpcWechatAccessToken2Vo
	GetAccessTokenOfPc(clientId, wechatCode string) *RpcWechatAccessTokenVo
	GetPhone(clientId, wechatCode, iv, encryptedData string) string
	GetQrcodeUrl(clientId, redirect_uri string) string
	SendSmsCode(clientId, grantType, phone string)
}
type WeChatServiceImpl struct {
}

func NewWeChatService() WeChatService {
	return WeChatServiceImpl{}
}

func (w WeChatServiceImpl) GetAccessTokenOfApp(clientId, wechatCode string) *RpcWechatAccessToken2Vo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/wechat/getAccessTokenOfApp")

	var param = make(map[string]any)
	param["clientId"] = clientId
	param["wechatCode"] = wechatCode

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result RpcWechatAccessToken2Vo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return &result
}

func (w WeChatServiceImpl) GetAccessTokenOfPc(clientId, wechatCode string) *RpcWechatAccessTokenVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/wechat/getAccessTokenOfPc")

	var param = make(map[string]any)
	param["clientId"] = clientId
	param["wechatCode"] = wechatCode

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}

	var result RpcWechatAccessTokenVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return &result
}

func (w WeChatServiceImpl) GetPhone(clientId, wechatCode, iv, encryptedData string) string {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/wechat/getPhone")

	var param = make(map[string]any)
	param["clientId"] = clientId
	param["wechatCode"] = wechatCode
	param["iv"] = iv
	param["encryptedData"] = encryptedData

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return ""
	}
	var result string
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (w WeChatServiceImpl) GetQrcodeUrl(clientId, redirect_uri string) string {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/wechat/getQrcodeUrl")

	var param = make(map[string]any)
	param["clientId"] = clientId
	param["redirect_uri"] = redirect_uri

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return ""
	}
	var result string
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (w WeChatServiceImpl) SendSmsCode(clientId, grantType, phone string) {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/wechat/sendSmsCode")

	var param = make(map[string]any)
	param["clientId"] = clientId
	param["grantType"] = grantType
	param["phone"] = phone

	HttpPost(url, param)
}
