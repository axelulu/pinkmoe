/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:02
 * @FilePath: /pinkmoe_server/model/xd_login_log.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import "server/global"

type XdLoginLog struct {
	global.XD_MODEL
	UserName      string `json:"userName" form:"userName" gorm:"comment:登录账号"`           // 登录账号
	Ip            string `json:"ip" form:"ip" gorm:"comment:登录IP地址"`                     // 登录IP地址
	LoginLocation string `json:"loginLocation" form:"loginLocation" gorm:"comment:登录地点"` // 登录地点
	Explorer      string `json:"explorer" form:"explorer" gorm:"comment:浏览器类型"`          // 浏览器类型
	Os            string `json:"os" form:"os" gorm:"comment:操作系统"`                       // 操作系统
	Status        int    `json:"status" form:"status" gorm:"comment:登录状态（0成功 1失败）"`      // 登录状态（0成功 1失败）
	Msg           string `json:"msg" form:"msg" gorm:"comment:提示消息"`                     // 提示消息
	Module        string `json:"module" form:"module" gorm:"comment:登录模块"`               // 登录模块
}
