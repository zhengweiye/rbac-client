package main

import (
	"github.com/zhengweiye/rbac-client/rbac"
)

func main() {
	ws := rbac.Ws{
		RbacHost: "192.168.64.131",
		//RbacHost: "127.0.0.1",
		RbacPort: 8080,
		RbacUrl:  "ws://%s:%d/ws/syncData?connId=%s&applicationName=%s",
		ClientId: "fastflow",
	}
	rbac.ConnToRbacWs(ws, nil)
}
