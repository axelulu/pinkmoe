/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:48
 * @FilePath: /pinkmoe_server/model/xd_topic.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import "server/global"

type XdTopic struct {
	global.XD_MODEL
	Value         string   `json:"value" form:"value" gorm:"comment:话题标识"` // 话题标识
	Label         string   `json:"label" form:"label" gorm:"comment:话题名称"` // 话题名称
	Icon          string   `json:"icon" form:"icon" gorm:"comment:话题图标"`   // 话题图标
	Sort          int      `json:"sort" form:"sort" gorm:"comment:话题排序"`   // 话题排序
	PostRelations []XdPost `json:"post" gorm:"many2many:xd_topic_relation"`
}
