package rbac

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var RbacIp string
var RbacPort int
var clientList []Client

func Init(rbacIp string, rbacPort int, clients []Client) {
	RbacIp = rbacIp
	RbacPort = rbacPort
	clientList = clients
}

func GetClientSecret(clientId string) string {
	for _, item := range clientList {
		if item.clientId == clientId {
			return item.clientSecret
		}
	}
	return ""
}

type Client struct {
	clientId     string
	clientSecret string
}

func HttpPost(url string, param map[string]any) []byte {
	var body *bytes.Reader
	if param != nil {
		reqBytes, err := json.Marshal(param)
		if err != nil {
			panic(err)
		}
		body = bytes.NewReader(reqBytes)
	}

	resp, err := http.Post(url, "application/json", body)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	//TODO 读取io
	resBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//TODO 先解析result && 获取data字段
	if len(resBytes) == 0 {
		panic("返回字节数组为空")
	}
	var m = make(map[string]any)
	json.Unmarshal(resBytes, &m)
	code := m["code"]
	if code == nil {
		panic("没有code字段")
	}
	codeInt := code.(int)
	if codeInt != 200 {
		msg := m["msg"]
		if msg != nil {
			panic("没有msg字段")
		}
		panic(msg)
	}
	data := m["data"]
	if data == nil {
		return nil
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return dataBytes
}
