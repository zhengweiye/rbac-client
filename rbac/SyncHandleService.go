package rbac

type SyncHandleService interface {
	SaveOrUpdateUsers(userList []SyncUser)
	SaveOrUpdateDepts(deptList []SyncDept)
	DelAndSaveUsers(userList []SyncUser)
	DelAndSaveDepts(deptList []SyncDept)
	ExcludeUrls(urls []string)
	IncludeUrls(urls []string)
}

type SyncDept struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Pid        int    `json:"pid"`
	HeadUserId int    `json:"headUserId"`
	DeptType   string `json:"deptType"`
	Enabled    bool   `json:"enabled"`
}

type SyncUser struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	JobName   string `json:"jobName"`
	UserType  string `json:"userType"`
	DeptIds   []int  `json:"deptIds"`
	Enabled   bool   `json:"enabled"`
}

type SyncUrl struct {
	Type      string   `json:"type"`
	RouteName string   `json:"routeName"`
	Urls      []string `json:"urls"`
}
