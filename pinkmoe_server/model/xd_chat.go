/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-30 15:17:22
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:33
 * @FilePath: /pinkmoe_server/model/xd_chat.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdChat struct {
	global.XD_MODEL
	UserId         uuid.UUID `json:"userId" form:"userId"`   //谁发的
	SendId         uuid.UUID `json:"sendId" form:"sendId"`   //对端用户ID/群ID
	Cmd            int       `json:"cmd" form:"cmd"`         //群聊还是私聊
	Media          int       `json:"media" form:"media"`     //消息按照什么样式展示
	Content        string    `json:"content" form:"content"` //消息的内容
	Pic            string    `json:"pic" form:"pic"`         //预览图片
	Url            string    `json:"url" form:"url"`         //服务的URL
	Memo           string    `json:"memo" form:"memo"`       //简单描述
	Amount         int       `json:"amount" form:"amount"`   //其他和数字相关的
	UserIdRelation XdUser    `json:"userIdRelation" gorm:"foreignKey:UserId;references:UUID;"`
	SendIdRelation XdUser    `json:"sendIdRelation" gorm:"foreignKey:SendId;references:UUID;"`
}
