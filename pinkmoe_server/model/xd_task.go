/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:38
 * @FilePath: /pinkmoe_server/model/xd_task.go
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

type XdTask struct {
	global.XD_MODEL
	Name    string        `json:"name" form:"name" gorm:"comment:任务名称"`         // 任务名称
	Slug    string        `json:"slug" form:"slug" gorm:"comment:任务标识"`         // 任务标识
	Type    string        `json:"type" form:"type" gorm:"comment:任务类型"`         // 任务类型
	Uuid    uuid.UUID     `json:"uuid" form:"uuid" gorm:"comment:任务用户ID"`       // 任务用户ID
	TimeOut global.XdTime `json:"timeOut" form:"timeOut" gorm:"comment:任务过期时间"` // 任务过期时间
}
