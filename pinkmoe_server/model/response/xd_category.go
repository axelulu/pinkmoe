/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:32
 * @FilePath: /pinkmoe_server/model/response/xd_category.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package response

import "server/global"

type XdCategory struct {
	global.XD_MODEL
	ParentId uint         `json:"parentId" form:"parentId" gorm:"comment:父菜单ID"` // 父菜单ID
	Slug     string       `json:"slug" form:"slug" gorm:"unique;comment:菜单标识"`   // 菜单标识
	Name     string       `json:"name" form:"name" gorm:"comment:菜单名称"`          // 菜单名称
	Icon     string       `json:"icon" form:"icon" gorm:"comment:菜单图标"`          // 菜单图标
	Sort     int          `json:"sort" form:"sort" gorm:"comment:排序标记"`          // 排序标记
	Children []XdCategory `json:"children" gorm:"-"`
}
