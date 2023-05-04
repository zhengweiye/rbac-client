package rbac

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func ConnToGatewayWs(ws Ws) {
	di := websocket.Dialer{}
	wsUrl := fmt.Sprintf(ws.RbacUrl, ws.RbacHost, ws.RbacPort)

	conn, _, err := di.Dial(wsUrl, nil)
	if err != nil {
		panic(err)
	}
	// 关闭连接（退出for循环时）
	defer conn.Close()

	// 发送消息
	go heartBeatToGateway(conn)

	// 接受消息
	receiveMsgFromGateway(conn)
}
func heartBeatToGateway(conn *websocket.Conn) {
	t := time.Tick(15 * time.Second)
	for {
		<-t
		conn.WriteMessage(websocket.TextMessage, []byte("ping"))
	}
}

func receiveMsgFromGateway(conn *websocket.Conn) {
	for {
		msgType, msgBytes, err := conn.ReadMessage()
		fmt.Println("msgType=", msgType)
		if err != nil {
			break
		}
		fmt.Println(string(msgBytes))
	}
}
