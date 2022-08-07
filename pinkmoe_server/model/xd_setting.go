/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:36
 * @FilePath: /pinkmoe_server/model/xd_setting.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import "server/global"

type XdSetting struct {
	global.XD_MODEL
	Value string `json:"value" form:"value" gorm:"comment:设置值;type:text"` // 设置值
	Slug  string `json:"slug" form:"slug" gorm:"comment:设置标识"`            // 设置标识
	Name  string `json:"name" form:"name" gorm:"comment:设置名称"`            // 设置名称
}
