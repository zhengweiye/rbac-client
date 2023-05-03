package rbac

import (
	"encoding/json"
	"fmt"
)

type MenuService interface {
	FindMenus(clientId string, userId int) []RpcMenusVo
	FindFuncs(clientId string, userId int) []string
}
type MenuServiceImpl struct {
}

func NewMenuService() MenuService {
	return MenuServiceImpl{}
}

func (m MenuServiceImpl) FindMenus(clientId string, userId int) []RpcMenusVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/menu/findMenus")

	var param = make(map[string]any)
	param["clientId"] = clientId
	param["userId"] = userId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}

	var result []RpcMenusVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (m MenuServiceImpl) FindFuncs(clientId string, userId int) []string {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/menu/findFuncs")

	var param = make(map[string]any)
	param["clientId"] = clientId
	param["userId"] = userId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []string
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}
