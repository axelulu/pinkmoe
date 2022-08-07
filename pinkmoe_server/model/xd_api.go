/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:07
 * @FilePath: /pinkmoe_server/model/xd_api.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import "server/global"

type XdApi struct {
	global.XD_MODEL
	Path        string `json:"path" form:"path" gorm:"comment:api路径"`                 // api路径
	Description string `json:"description" form:"description" gorm:"comment:api中文描述"` // api中文描述
	ApiGroup    string `json:"apiGroup" form:"apiGroup" gorm:"comment:api组"`          // api组
	Method      string `json:"method" form:"method" gorm:"default:POST;comment:方法"`   // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}
