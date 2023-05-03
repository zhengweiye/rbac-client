package rbac

import (
	"encoding/json"
	"fmt"
)

type DeptService interface {
	FindOwnDeptByUserId(userId int) []RpcDeptTreeVo
	FindOwnDeptUserByUserId(userId int) []RpcDeptAndUserTreeVo
	FindThisByDeptId(deptId int) []RpcDeptTreeVo
	FindThisByDeptIds(deptIds []int) []RpcDeptTreeVo
	FindThisAndChildByDeptId(deptId int) []RpcDeptTreeVo
	FindThisAndChildByDeptIds(deptIds []int) []RpcDeptTreeVo
	FindChildByDeptId(deptId int) []RpcDeptTreeVo
	FindDept(deptId int) *RpcDeptVo
	FindNames(ids []int) []RpcNameVo
	FindName(id int) *RpcNameVo
	CreateDept(deptName string, deptPid int) int
	UpdateDept(deptId int, deptName string, deptPid int, headUserId int)
	DeleteDept(deptId int)
}
type DeptServiceImpl struct {
}

func NewDeptService() DeptService {
	return DeptServiceImpl{}
}

func (d DeptServiceImpl) FindOwnDeptByUserId(userId int) []RpcDeptTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/findOwnDeptByUserId")

	var param = make(map[string]any)
	param["userId"] = userId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindOwnDeptUserByUserId(userId int) []RpcDeptAndUserTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/findOwnDeptUserByUserId")

	var param = make(map[string]any)
	param["userId"] = userId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []RpcDeptAndUserTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindThisByDeptId(deptId int) []RpcDeptTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/findThisByDeptId")

	var param = make(map[string]any)
	param["deptId"] = deptId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindThisByDeptIds(deptIds []int) []RpcDeptTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/findThisByDeptIds")

	var param = make(map[string]any)
	param["deptIds"] = deptIds

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindThisAndChildByDeptId(deptId int) []RpcDeptTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/findThisAndChildByDeptId")

	var param = make(map[string]any)
	param["deptId"] = deptId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindThisAndChildByDeptIds(deptIds []int) []RpcDeptTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/findThisAndChildByDeptIds")

	var param = make(map[string]any)
	param["deptIds"] = deptIds

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindChildByDeptId(deptId int) []RpcDeptTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/findChildByDeptId")

	var param = make(map[string]any)
	param["deptId"] = deptId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindDept(deptId int) *RpcDeptVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/findDept")

	var param = make(map[string]any)
	param["deptId"] = deptId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result RpcDeptVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}
	return &result
}

func (d DeptServiceImpl) FindNames(ids []int) []RpcNameVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/findNames")

	var param = make(map[string]any)
	param["ids"] = ids

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []RpcNameVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func (d DeptServiceImpl) FindName(id int) *RpcNameVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/findName")

	var param = make(map[string]any)
	param["id"] = id

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result RpcNameVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}
	return &result
}

func (d DeptServiceImpl) CreateDept(deptName string, deptPid int) int {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/createDept")

	var param = make(map[string]any)
	param["deptName"] = deptName
	param["deptPid"] = deptPid

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return 0
	}
	var result int
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) UpdateDept(deptId int, deptName string, deptPid int, headUserId int) {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/updateDept")

	var param = make(map[string]any)
	param["deptId"] = deptId
	param["deptName"] = deptName
	param["deptPid"] = deptPid
	param["headUserId"] = headUserId

	HttpPost(url, param)
}

func (d DeptServiceImpl) DeleteDept(deptId int) {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/dept/deleteDept")

	var param = make(map[string]any)
	param["deptId"] = deptId

	HttpPost(url, param)
}
