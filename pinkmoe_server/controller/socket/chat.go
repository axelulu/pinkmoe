/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-31 06:07:01
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:31
 * @FilePath: /pinkmoe_server/controller/socket/chat.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package socket

import (
	"encoding/json"
	"server/dao/mysql"
	"server/model"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type ChatStruct struct {
	Type    string       `json:"type"`
	ChatMsg model.XdChat `json:"chatMsg"`
}

type ChatMsgStruct struct {
	Content        string       `json:"content"`
	Media          string       `json:"media"`
	Cmd            string       `json:"cmd"`
	Amount         string       `json:"amount"`
	Memo           string       `json:"memo"`
	Url            string       `json:"url"`
	Pic            string       `json:"pic"`
	SendId         string       `json:"sendId" form:"sendId"`
	UserId         uuid.UUID    `json:"userId" form:"userId"`
	UserIdRelation model.XdUser `json:"userIdRelation"`
}

func SendChatMsg(chatMsg ChatMsgStruct) {
	SendUid, _ := uuid.FromString(chatMsg.SendId)
	err, chat := mysql.CreateChat(&model.XdChat{
		UserId:  chatMsg.UserId,
		SendId:  SendUid,
		Cmd:     0,
		Media:   0,
		Content: chatMsg.Content,
		Pic:     "",
		Url:     "",
		Memo:    "",
		Amount:  0,
	})
	if err != nil {
		return
	}
	if err, chat.UserIdRelation = mysql.GetUserInfoByUserId(chatMsg.UserId.String()); err != nil {
		return
	}
	marshal, err := json.Marshal(ChatStruct{
		Type:    "chat",
		ChatMsg: chat,
	})
	if err != nil {
		return
	}
	// 发给用户
	if cc, ok := clientMap[chatMsg.SendId]; ok {
		err = cc.Conn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			println("用户已经下线")
		}
	}
	// 发给自己
	if cc, ok := clientMap[chatMsg.UserId.String()]; ok {
		err = cc.Conn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			println("自己已经下线")
		}
	}
}
