/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-30 11:13:00
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:34
 * @FilePath: /pinkmoe_server/controller/socket/socket.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package socket

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"server/model/response"
	jwt "server/util"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
)

var upgrader = websocket.Upgrader{}

type SocketMsg struct {
	Type    string        `json:"type"`
	ChatMsg ChatMsgStruct `json:"chatMsg"`
}

type ReqChat struct {
	Token   string `json:"token" form:"token"`
	SendUid string `json:"sendUid" form:"sendUid"`
}

//读写锁
var rwlocker sync.RWMutex

// Node 本核心在于形成userid和Node的映射关系
type Node struct {
	Conn *websocket.Conn
	//并行转串行,
	DataQueue chan []byte
	GroupSets set.Interface
}

//映射关系表
var clientMap = make(map[string]*Node, 0)

func InitSocket(c *gin.Context) {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.L().Error("chat err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	//todo 获得conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	var p ReqChat
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("ReqChat with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}

	currentUserInfo, err := jwt.ParseToken(p.Token)
	if err == nil {
		rwlocker.Lock()
		clientMap[currentUserInfo.UserID.String()] = node
		rwlocker.Unlock()
		//todo 完成发送逻辑,con
		go sendproc(node)
		//todo 完成接收逻辑
		go recvproc(node)
	} else {
		return
	}
}

//ws接收协程
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		//dispatch(data)
		//把消息广播到局域网
		broadMsg(data)
	}
}

type sendData struct {
	userId uuid.UUID
	params ReqChat
	data   []byte
}

//用来存放发送的要广播的数据
var udpsendchan chan []byte = make(chan []byte, 1024)

//todo 将消息广播到局域网
func broadMsg(data []byte) {
	udpsendchan <- data
}

func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

func init() {
	go udpsendproc()
	go udprecvproc()
}

//todo 完成udp数据的发送协程
func udpsendproc() {
	//todo 使用udp协议拨号
	con, err := net.DialUDP("udp", nil,
		&net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: 3000,
		})
	defer con.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	//todo 通过得到的con发送消息
	for {
		select {
		case data := <-udpsendchan:
			_, err = con.Write(data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

//todo 完成upd接收并处理功能
func udprecvproc() {
	//todo 监听udp广播端口
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		log.Println(err.Error())
	}
	//TODO 处理端口发过来的数据
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			log.Println(err.Error())
			return
		}
		//直接数据处理
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	var socketMsg SocketMsg
	if err := json.Unmarshal(data, &socketMsg); err != nil {
		return
	}
	switch socketMsg.Type {
	case "chat":
		SendChatMsg(socketMsg.ChatMsg)
		break
	case "notice":
		NoticeMsg()
		break
	// 处理心跳响应 , heartbeat为与客户端约定的值
	case "heartbeat":
		// 发给用户
		if cc, ok := clientMap[socketMsg.ChatMsg.UserId.String()]; ok {
			err := cc.Conn.WriteMessage(websocket.TextMessage, []byte(`{"type": "heartbeat"}`))
			if err != nil {
				println("心跳发送失败")
			}
		}
		break
	}
}
