package rbac

import (
	"encoding/json"
	"fmt"
)

type IconService interface {
	FindIconList() []RpcIconVo
}
type IconServiceImpl struct {
}

func NewIconService() IconService {
	return IconServiceImpl{}
}

func (i IconServiceImpl) FindIconList() []RpcIconVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/icon/findIconList")

	resBytes := HttpPost(url, nil)

	var result []RpcIconVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}
