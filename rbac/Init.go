package rbac

var rbacIp string
var rbacPort int
var clientList []Client
var syncHandleService SyncHandleService

type Client struct {
	ClientId     string
	ClientSecret string
}

type Ws struct {
	RbacHost       string
	RbacPort       int
	RbacEnabled    bool
	RbacUrl        string
	ClientId       string
	GatewayHost    string
	GatewayPort    int
	GatewayEnabled bool
	GatewayUrl     string
}

func Init(ip string, port int, clients []Client, ws Ws, handleService SyncHandleService) {
	rbacIp = ip
	rbacPort = port
	clientList = clients

	if ws.RbacEnabled {
		ConnToRbacWs(ws, nil)
	}
	if ws.GatewayEnabled {
		ConnToGatewayWs(ws)
	}
	syncHandleService = handleService
}

func GetClientSecret(clientId string) string {
	for _, item := range clientList {
		if item.ClientId == clientId {
			return item.ClientSecret
		}
	}
	return ""
}

func GetIP() string {
	return rbacIp
}

func GetPort() int {
	return rbacPort
}
