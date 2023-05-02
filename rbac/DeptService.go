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
	FindDept(deptId int) []RpcDeptVo
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
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/dept/findOwnDeptByUserId")

	var param = make(map[string]any)
	param["userId"] = userId

	resBytes := HttpPost(url, param)

	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindOwnDeptUserByUserId(userId int) []RpcDeptAndUserTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/dept/findOwnDeptUserByUserId")

	var param = make(map[string]any)
	param["userId"] = userId

	resBytes := HttpPost(url, param)

	var result []RpcDeptAndUserTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindThisByDeptId(deptId int) []RpcDeptTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/dept/findThisByDeptId")

	var param = make(map[string]any)
	param["deptId"] = deptId

	resBytes := HttpPost(url, param)

	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindThisByDeptIds(deptIds []int) []RpcDeptTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/dept/findThisByDeptIds")

	var param = make(map[string]any)
	param["deptIds"] = deptIds

	resBytes := HttpPost(url, param)

	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindThisAndChildByDeptId(deptId int) []RpcDeptTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/dept/findThisAndChildByDeptId")

	var param = make(map[string]any)
	param["deptId"] = deptId

	resBytes := HttpPost(url, param)

	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindThisAndChildByDeptIds(deptIds []int) []RpcDeptTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/dept/findThisAndChildByDeptIds")

	var param = make(map[string]any)
	param["deptIds"] = deptIds

	resBytes := HttpPost(url, param)

	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindChildByDeptId(deptId int) []RpcDeptTreeVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/dept/findChildByDeptId")

	var param = make(map[string]any)
	param["deptId"] = deptId

	resBytes := HttpPost(url, param)

	var result []RpcDeptTreeVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) FindDept(deptId int) []RpcDeptVo {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/dept/findDept")

	var param = make(map[string]any)
	param["deptId"] = deptId

	resBytes := HttpPost(url, param)

	var result []RpcDeptVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) CreateDept(deptName string, deptPid int) int {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/dept/createDept")

	var param = make(map[string]any)
	param["deptName"] = deptName
	param["deptPid"] = deptPid

	resBytes := HttpPost(url, param)

	var result int
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (d DeptServiceImpl) UpdateDept(deptId int, deptName string, deptPid int, headUserId int) {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/dept/updateDept")

	var param = make(map[string]any)
	param["deptId"] = deptId
	param["deptName"] = deptName
	param["deptPid"] = deptPid
	param["headUserId"] = headUserId

	HttpPost(url, param)
}

func (d DeptServiceImpl) DeleteDept(deptId int) {
	url := fmt.Sprintf("http://%s:%d/%s", RbacIp, RbacPort, "rpc/dept/deleteDept")

	var param = make(map[string]any)
	param["deptId"] = deptId

	HttpPost(url, param)
}
