package rbac

import (
	"encoding/json"
	"fmt"
)

type UserService interface {
	FindUserInfo(userId int) *RpcUserDetailVo
	FindUserListByDeptId(deptId int) []RpcUserVo
	FindUserListByRoleId(roleId int) []RpcUserVo
	FindUserList(roleId int, deptIds []int) []RpcUserVo
	FindNames(ids []int) []RpcNameVo
	FindName(id int) *RpcNameVo

	FindDataScopeOfDept(userId int) []int
	FindDataScopeOfUser(userId int) []int

	CreateUser(phone, password, realName, jobName, userType string, deptId int) int
	UpdateWxInfo(userId int, nickName, avatarUrl string)
	UpdatePwd(userId int, password string)
	UpdateName(userId int, realName string)
	UpdateDepts(userId int, deptIds []int)
	DeleteById(userId int)
	DeleteByTypeAndPhone(userType string, phone string)
}

type UserServiceImpl struct {
}

func NewUserService() UserService {
	return UserServiceImpl{}
}

func (u UserServiceImpl) FindUserInfo(userId int) *RpcUserDetailVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/findUserInfo")

	var param = make(map[string]any)
	param["userId"] = userId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}

	var result RpcUserDetailVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return &result
}

func (u UserServiceImpl) FindUserListByDeptId(deptId int) []RpcUserVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/findUserListByDeptId")

	var param = make(map[string]any)
	param["deptId"] = deptId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []RpcUserVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (u UserServiceImpl) FindUserListByRoleId(roleId int) []RpcUserVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/findUserListByRoleId")

	var param = make(map[string]any)
	param["roleId"] = roleId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []RpcUserVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (u UserServiceImpl) FindUserList(roleId int, deptIds []int) []RpcUserVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/findUserList")

	var param = make(map[string]any)
	param["roleId"] = roleId
	param["deptIds"] = deptIds

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []RpcUserVo
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (u UserServiceImpl) FindNames(ids []int) []RpcNameVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/findNames")

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

func (u UserServiceImpl) FindName(id int) *RpcNameVo {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/findName")

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

func (u UserServiceImpl) FindDataScopeOfDept(userId int) []int {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/findDataScopeOfDept")

	var param = make(map[string]any)
	param["userId"] = userId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []int
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (u UserServiceImpl) FindDataScopeOfUser(userId int) []int {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/findDataScopeOfUser")

	var param = make(map[string]any)
	param["userId"] = userId

	resBytes := HttpPost(url, param)
	if resBytes == nil {
		return nil
	}
	var result []int
	err := json.Unmarshal(resBytes, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (u UserServiceImpl) CreateUser(phone, password, realName, jobName, userType string, deptId int) int {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/createUser")

	var param = make(map[string]any)
	param["phone"] = phone
	param["password"] = password
	param["realName"] = realName
	param["jobName"] = jobName
	param["userType"] = userType
	param["deptId"] = deptId

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

func (u UserServiceImpl) UpdateWxInfo(userId int, nickName, avatarUrl string) {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/updateUserWxInfo")

	var param = make(map[string]any)
	param["userId"] = userId
	param["nickName"] = nickName
	param["avatarUrl"] = avatarUrl

	HttpPost(url, param)
}

func (u UserServiceImpl) UpdatePwd(userId int, password string) {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/updateUserPwd")

	var param = make(map[string]any)
	param["userId"] = userId
	param["password"] = password

	HttpPost(url, param)
}

func (u UserServiceImpl) UpdateName(userId int, realName string) {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/updateUserName")

	var param = make(map[string]any)
	param["userId"] = userId
	param["realName"] = realName

	HttpPost(url, param)
}

func (u UserServiceImpl) UpdateDepts(userId int, deptIds []int) {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/updateUserDepts")

	var param = make(map[string]any)
	param["userId"] = userId
	param["deptIds"] = deptIds

	HttpPost(url, param)
}

func (u UserServiceImpl) DeleteById(userId int) {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/deleteUserById")

	var param = make(map[string]any)
	param["userId"] = userId

	HttpPost(url, param)
}

func (u UserServiceImpl) DeleteByTypeAndPhone(userType string, phone string) {
	url := fmt.Sprintf("http://%s:%d/%s", GetIP(), GetPort(), "rpc/user/deleteUserByTypeAndPhone")

	var param = make(map[string]any)
	param["userType"] = userType
	param["phone"] = phone

	HttpPost(url, param)
}
