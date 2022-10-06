/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-10 22:11:30
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 17:59:07
 * @FilePath: /xanaduCms/pinkmoe_server/model/xd_check_in.go
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

type XdCheckIn struct {
	global.XD_MODEL
	Uuid            uuid.UUID     `json:"uuid" gorm:"comment:用户ID"`              // 用户ID
	Credit          int           `json:"credit" gorm:"comment:奖励积分"`            // 奖励积分
	CheckType       string        `json:"checkType" gorm:"comment:奖励类型"`         // 奖励类型
	Status          int           `json:"status" gorm:"comment:状态"`              // 状态
	LastLoginIp     string        `json:"type" gorm:"comment:登陆ip"`              // 登陆ip
	LastCheckinTime global.XdTime `json:"lastCheckinTime" gorm:"comment:上次签到时间"` // 上次签到时间
	CheckinCount    int           `json:"checkinCount" gorm:"comment:连续签到天数"`    // 连续签到天数
}
