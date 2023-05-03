package rbac

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"time"
)

// 创建客户端
var client = &http.Client{
	//请求超时时间
	Timeout: time.Second * 30,
	// 创建连接池
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second, // 长连接超时时间
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接
		IdleConnTimeout:       90 * time.Second, // 空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, // tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  // 100-continue状态码超时时间
	},
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

	resp, err := client.Post(url, "application/json", body)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	//TODO 读取io！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！
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

	switch code.(type) {
	case int:
		codeInt := code.(int)
		if codeInt != 200 {
			msg := m["msg"]
			if msg != nil {
				panic("没有msg字段")
			}
			panic(msg)
		}
	case int32:
		codeInt := code.(int32)
		if codeInt != 200 {
			msg := m["msg"]
			if msg != nil {
				panic("没有msg字段")
			}
			panic(msg)
		}
	case int64:
		codeInt := code.(int64)
		if codeInt != 200 {
			msg := m["msg"]
			if msg != nil {
				panic("没有msg字段")
			}
			panic(msg)
		}
	case float32:
		codeInt := code.(float32)
		if codeInt != 200 {
			msg := m["msg"]
			if msg != nil {
				panic("没有msg字段")
			}
			panic(msg)
		}
	case float64:
		codeInt := code.(float64)
		if codeInt != 200 {
			msg := m["msg"]
			if msg != nil {
				panic("没有msg字段")
			}
			panic(msg)
		}
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
