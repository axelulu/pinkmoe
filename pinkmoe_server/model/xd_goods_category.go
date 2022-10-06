/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:25
 * @FilePath: /pinkmoe_server/model/xd_category.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package model

import "server/global"

type XdGoodsCategory struct {
	global.XD_MODEL
	ParentId uint              `json:"parentId" form:"parentId" gorm:"comment:父菜单ID"` // 父菜单ID
	Slug     string            `json:"slug" form:"slug" gorm:"unique;comment:菜单标识"`   // 菜单标识
	Name     string            `json:"name" form:"name" gorm:"comment:菜单名称"`          // 菜单名称
	Icon     string            `json:"icon" form:"icon" gorm:"comment:菜单图标"`          // 菜单图标
	Sort     int               `json:"sort" form:"sort" gorm:"comment:排序标记"`          // 排序标记
	Children []XdGoodsCategory `json:"children" gorm:"-"`
}
