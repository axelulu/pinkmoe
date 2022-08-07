/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:57
 * @FilePath: /pinkmoe_server/model/xd_user_online.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdUserOnline struct {
	global.XD_MODEL
	Uuid     uuid.UUID `json:"uuid" form:"uuid" gorm:"comment:用户标识"`                // 用户标识
	Token    string    `json:"token" form:"token" gorm:"type:text;comment:用户token"` // 用户token
	UserName string    `json:"userName" form:"userName" gorm:"comment:用户名"`         // 用户名
	Ip       string    `json:"ip" form:"ip" gorm:"comment:登录ip"`                    // 登录ip
	Explorer string    `json:"explorer" form:"explorer" gorm:"comment:浏览器"`         // 浏览器
	Os       string    `json:"os" form:"os" gorm:"comment:操作系统"`                    // 操作系统
}
