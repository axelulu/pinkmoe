/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:38
 * @FilePath: /pinkmoe_server/model/xd_task.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdTask struct {
	global.XD_MODEL
	Name    string        `json:"name" form:"name" gorm:"comment:任务名称"`         // 任务名称
	Slug    string        `json:"slug" form:"slug" gorm:"comment:任务标识"`         // 任务标识
	Type    string        `json:"type" form:"type" gorm:"comment:任务类型"`         // 任务类型
	Uuid    uuid.UUID     `json:"uuid" form:"uuid" gorm:"comment:任务用户ID"`       // 任务用户ID
	TimeOut global.XdTime `json:"timeOut" form:"timeOut" gorm:"comment:任务过期时间"` // 任务过期时间
}
