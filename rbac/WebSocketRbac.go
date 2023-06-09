package rbac

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"io"
	"time"
)

var wsRbacClient = websocket.Dialer{}
var wsRbacConn *websocket.Conn

func ConnToRbacWs(ws Ws, timeTicker *time.Ticker) {
	wsUrl := fmt.Sprintf(ws.RbacUrl, ws.RbacHost, ws.RbacPort, uuid.NewV4().String(), ws.ClientId)
	conn, res, err := wsRbacClient.Dial(wsUrl, nil)

	if err != nil {
		var content string
		if res != nil {
			dataBytes, err := io.ReadAll(res.Body)
			if err == nil {
				content = string(dataBytes)
			}
		}

		fmt.Println("连接Rbac失败：", err, content)
		reconnToRbac(ws) // TODO 连接错误时，重连
		return
	} else {
		fmt.Println("连接Rbac成功：", res)
	}

	// 连接赋值
	wsRbacConn = conn

	// 停止定时器
	if timeTicker != nil {
		timeTicker.Stop()
	}

	// 关闭连接（receiveMsgFromRbac退出for循环时）
	defer reconnToRbac(ws) // TODO 连接断开时，重连

	// 发送消息
	go heartBeatToRbac()

	// 接受消息
	receiveMsgFromRbac()
}

func reconnToRbac(ws Ws) {
	if wsRbacConn != nil {
		wsRbacConn.Close()
		wsRbacConn = nil
	}

	// TODO断开重连
	var timeTicker *time.Ticker
	timeTicker = time.NewTicker(30 * time.Second)
	for {
		<-timeTicker.C
		timeTicker.Stop()
		ConnToRbacWs(ws, timeTicker)

		//TODO 需求：重连成功，那么停止定时器
		//TODO 永远不会执行到这里，因为连接成功之后往下执行ConnToRbacWs后面的代码，导致堵塞了
		/*if wsRbacConn != nil {
			fmt.Println("重连Rbac服务成功")
			break
		}*/
	}
}

func heartBeatToRbac() {
	t := time.Tick(15 * time.Second)
	for {
		<-t

		if wsRbacConn != nil {
			//TODO 写法一
			err := wsRbacConn.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				fmt.Println("heartBeatToRbac：发送心跳失败...")
			}
			//TODO 写法二
			//pingHandler := wsRbacConn.PingHandler()
			//pingHandler("ping")
		}
	}
}

func receiveMsgFromRbac() {
	for {
		msgType, msgBytes, err := wsRbacConn.ReadMessage()
		if err != nil {
			break //TODO 断开连接
		}
		if msgType == 2 {
			var mapData = make(map[string]any)
			err = json.Unmarshal(msgBytes, &mapData)
			if err == nil { // TODO 解析json错误，不需要断开连接
				handleRbacMsg(mapData)
			}
		}
	}
}

func handleRbacMsg(mapData map[string]any) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("WebSocketRbac##handleRbacMsg执行错误：", err)
		}
	}()

	operType, ok1 := mapData["operType"]
	tableName, ok2 := mapData["tableName"]
	list, ok3 := mapData["list"]

	if ok1 && ok2 && ok3 {
		if operType == "push_data" && tableName == "sys_user" {
			dataBytes, err := json.Marshal(list)
			if err == nil {
				var userList []SyncUser
				err = json.Unmarshal(dataBytes, &userList)
				if err == nil {
					syncHandleService.SaveOrUpdateUsers(userList)
				}
			}

		} else if operType == "push_data" && tableName == "sys_dept" {
			dataBytes, err := json.Marshal(list)
			if err == nil {
				var deptList []SyncDept
				err = json.Unmarshal(dataBytes, &deptList)
				if err == nil {
					syncHandleService.SaveOrUpdateDepts(deptList)
				}
			}

		} else if operType == "repush_data" && tableName == "sys_user" {
			dataBytes, err := json.Marshal(list)
			if err == nil {
				var userList []SyncUser
				err = json.Unmarshal(dataBytes, &userList)
				if err == nil {
					syncHandleService.DelAndSaveUsers(userList)
				}
			}

		} else if operType == "repush_data" && tableName == "sys_dept" {
			dataBytes, err := json.Marshal(list)
			if err == nil {
				var deptList []SyncDept
				err = json.Unmarshal(dataBytes, &deptList)
				if err == nil {
					syncHandleService.DelAndSaveDepts(deptList)
				}
			}
		}
	}
}
