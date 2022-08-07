/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-30 15:18:36
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:29
 * @FilePath: /pinkmoe_server/model/xd_chat_relation.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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
