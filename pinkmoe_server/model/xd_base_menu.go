/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:20
 * @FilePath: /pinkmoe_server/model/xd_base_menu.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package model

import "server/global"

type XdBaseMenu struct {
	global.XD_MODEL
	ParentId     uint                              `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Path         string                            `json:"path" gorm:"comment:路由path"`        // 路由path
	Name         string                            `json:"name" gorm:"comment:路由name"`        // 路由name
	Hidden       bool                              `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component    string                            `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Sort         int                               `json:"sort" gorm:"comment:排序标记"`          // 排序标记
	Meta         `json:"meta" gorm:"comment:附加属性"` // 附加属性
	XdAuthoritys []XdAuthority                     `json:"authoritys" gorm:"many2many:xd_authority_menus;"`
	Children     []XdBaseMenu                      `json:"children" gorm:"-"`
}

type Meta struct {
	Title string `json:"title" gorm:"comment:菜单名"` // 菜单名
	Icon  string `json:"icon" gorm:"comment:菜单图标"` // 菜单图标
}
