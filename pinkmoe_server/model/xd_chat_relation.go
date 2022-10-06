/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-30 15:18:36
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:29
 * @FilePath: /pinkmoe_server/model/xd_chat_relation.go
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

type XdChatRelation struct {
	global.XD_MODEL
	UserId         uuid.UUID `json:"userId" form:"userId"` //谁发的
	SendId         uuid.UUID `json:"sendId" form:"sendId"` //对端用户ID/群ID
	UserIdRelation XdUser    `json:"userIdRelation" gorm:"foreignKey:UserId;references:UUID;"`
	SendIdRelation XdUser    `json:"sendIdRelation" gorm:"foreignKey:SendId;references:UUID;"`
}
