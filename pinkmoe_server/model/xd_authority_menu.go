/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:09
 * @FilePath: /pinkmoe_server/model/xd_authority_menu.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

type XdMenu struct {
	XdBaseMenu
	MenuId      string   `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string   `json:"-" gorm:"comment:角色ID"`
	Children    []XdMenu `json:"children" gorm:"-"`
}

func (s XdMenu) TableName() string {
	return "authority_menu"
}
