package main

import (
	"github.com/zhengweiye/rbac-client/rbac"
)

func main() {
	ws := rbac.Ws{
		RbacHost:         "192.168.64.131",
		RbacPort:         8080,
		RbacUrl:          "ws://%s:%d/ws/syncData?connId=%s&applicationName=%s",
		ClientId:         "fastflow",
		RbacEnabled:      true,
		GatewayHost:      "192.168.1.143",
		GatewayPort:      80,
		GatewayEnabled:   true,
		GatewayUrl:       "ws://%s:%d/ws/urlAuthen?connId=%s&routeName=%s",
		GatewayRouteName: "route_law_base",
	}
	//rbac.ConnToRbacWs(ws, nil)
	rbac.ConnToGatewayWs(ws, nil)
}
