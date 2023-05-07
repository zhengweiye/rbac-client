package rbac

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"io"
	"time"
)

var wsGatewayClient = websocket.Dialer{}
var wsGatewayConn *websocket.Conn

func ConnToGatewayWs(ws Ws, timeTicker *time.Ticker) {
	wsUrl := fmt.Sprintf(ws.GatewayUrl, ws.GatewayHost, ws.GatewayPort, uuid.NewV4().String(), ws.GatewayRouteName)
	conn, res, err := wsGatewayClient.Dial(wsUrl, nil)

	if err != nil {
		var content string
		if res != nil {
			dataBytes, err := io.ReadAll(res.Body)
			if err == nil {
				content = string(dataBytes)
			}
		}

		fmt.Println("连接Gateway失败：", err, content)
		reconnToGateway(ws) // TODO 连接错误时，重连
		return
	} else {
		fmt.Println("连接Gateway成功：", res)
	}

	// 连接赋值
	wsGatewayConn = conn

	// 停止定时器
	if timeTicker != nil {
		timeTicker.Stop()
	}

	// 关闭连接（receiveMsgFromRbac退出for循环时）
	defer reconnToGateway(ws) // TODO 连接断开时，重连

	// 发送消息
	go heartBeatToGateway()

	// 接受消息
	receiveMsgFromGateway()
}

func reconnToGateway(ws Ws) {
	if wsGatewayConn != nil {
		wsGatewayConn.Close()
		wsGatewayConn = nil
	}

	// TODO断开重连
	var timeTicker *time.Ticker
	timeTicker = time.NewTicker(30 * time.Second)
	for {
		<-timeTicker.C
		timeTicker.Stop()
		ConnToGatewayWs(ws, timeTicker)

		//TODO 需求：重连成功，那么停止定时器
		//TODO 永远不会执行到这里，因为连接成功之后往下执行ConnToRbacWs后面的代码，导致堵塞了
		/*if wsRbacConn != nil {
			fmt.Println("重连Rbac服务成功")
			break
		}*/
	}
}

func heartBeatToGateway() {
	t := time.Tick(15 * time.Second)
	for {
		<-t

		if wsGatewayConn != nil {
			//TODO 写法一
			err := wsGatewayConn.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				fmt.Println("heartBeatToRbac：发送心跳失败...")
			}
			//TODO 写法二
			//pingHandler := wsRbacConn.PingHandler()
			//pingHandler("ping")
		}
	}
}

func receiveMsgFromGateway() {
	for {
		msgType, msgBytes, err := wsGatewayConn.ReadMessage()
		if err != nil {
			break //TODO 断开连接
		}
		if msgType == 2 {
			var syncUrl SyncUrl
			err = json.Unmarshal(msgBytes, &syncUrl)
			if err == nil { // TODO 解析json错误，不需要断开连接
				handleGatewayMsg(syncUrl)
			}
		}
	}
}

func handleGatewayMsg(syncUrl SyncUrl) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("WebSocketGateway##handleGatewayMsg执行错误：", err)
		}
	}()

	if syncUrl.Type == "include" {
		syncHandleService.IncludeUrls(syncUrl.Urls)

	} else if syncUrl.Type == "exclude" {
		syncHandleService.ExcludeUrls(syncUrl.Urls)
	}
}
