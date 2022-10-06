/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:08
 * @FilePath: /pinkmoe_server/model/xd_operation_log.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package model

import (
	"server/global"
	"time"

	uuid "github.com/satori/go.uuid"
)

type XdOperationLog struct {
	global.XD_MODEL
	Ip           string        `json:"ip" form:"ip" gorm:"comment:请求ip"`                       // 请求ip
	Method       string        `json:"method" form:"method" gorm:"comment:请求方法"`               // 请求方法
	Path         string        `json:"path" form:"path" gorm:"comment:请求路径"`                   // 请求路径
	Status       int           `json:"status" form:"status" gorm:"comment:请求状态"`               // 请求状态
	Latency      time.Duration `json:"latency" form:"latency" gorm:"comment:延迟"`               // 延迟
	Agent        string        `json:"agent" form:"agent" gorm:"comment:代理"`                   // 代理
	ErrorMessage string        `json:"errorMessage" form:"errorMessage" gorm:"comment:错误信息"`   // 错误信息
	Body         string        `json:"body" form:"body" gorm:"type:text;comment:请求Body"`       // 请求Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;comment:响应Body"`       // 响应Body
	UserID       uuid.UUID     `json:"userId" form:"userId" gorm:"comment:用户id;default:null;"` // 用户id
	User         XdUser        `json:"user"`
}
