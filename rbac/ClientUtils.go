package rbac

var rbacIp string
var rbacPort int

var clientList []Client

type Client struct {
	ClientId     string
	ClientSecret string
}

func Init(ip string, port int, clients []Client) {
	rbacIp = ip
	rbacPort = port
	clientList = clients
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
