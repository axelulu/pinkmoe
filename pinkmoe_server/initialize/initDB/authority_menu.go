/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:45
 * @FilePath: /pinkmoe_server/initialize/initDB/authority_menu.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package initDB

import (
	"server/global"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

func (a *authoritiesMenus) TableName() string {
	var entity AuthorityMenus
	return entity.TableName()
}

func (a *authoritiesMenus) Initialize() error {
	entities := []AuthorityMenus{
		{BaseMenuId: 1, AuthorityId: "2333"},

		{BaseMenuId: 1, AuthorityId: "9527"},
		{BaseMenuId: 2, AuthorityId: "9527"},
		{BaseMenuId: 3, AuthorityId: "9527"},
		{BaseMenuId: 4, AuthorityId: "9527"},
		{BaseMenuId: 5, AuthorityId: "9527"},
		{BaseMenuId: 6, AuthorityId: "9527"},
		{BaseMenuId: 7, AuthorityId: "9527"},
		{BaseMenuId: 8, AuthorityId: "9527"},
		{BaseMenuId: 9, AuthorityId: "9527"},
		{BaseMenuId: 10, AuthorityId: "9527"},
		{BaseMenuId: 11, AuthorityId: "9527"},
		{BaseMenuId: 12, AuthorityId: "9527"},
		{BaseMenuId: 13, AuthorityId: "9527"},
		{BaseMenuId: 14, AuthorityId: "9527"},
		{BaseMenuId: 15, AuthorityId: "9527"},
		{BaseMenuId: 16, AuthorityId: "9527"},
		{BaseMenuId: 17, AuthorityId: "9527"},
		{BaseMenuId: 18, AuthorityId: "9527"},
		{BaseMenuId: 19, AuthorityId: "9527"},
		{BaseMenuId: 20, AuthorityId: "9527"},
		{BaseMenuId: 21, AuthorityId: "9527"},
		{BaseMenuId: 22, AuthorityId: "9527"},
		{BaseMenuId: 23, AuthorityId: "9527"},
		{BaseMenuId: 24, AuthorityId: "9527"},
		{BaseMenuId: 25, AuthorityId: "9527"},
		{BaseMenuId: 26, AuthorityId: "9527"},
		{BaseMenuId: 27, AuthorityId: "9527"},
		{BaseMenuId: 28, AuthorityId: "9527"},
		{BaseMenuId: 29, AuthorityId: "9527"},
		{BaseMenuId: 30, AuthorityId: "9527"},
		{BaseMenuId: 31, AuthorityId: "9527"},
		{BaseMenuId: 32, AuthorityId: "9527"},
		{BaseMenuId: 33, AuthorityId: "9527"},
		{BaseMenuId: 34, AuthorityId: "9527"},
		{BaseMenuId: 35, AuthorityId: "9527"},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *authoritiesMenus) CheckDataExist() bool {
	if errors.Is(global.XD_DB.Where("xd_base_menu_id = ? AND xd_authority_authority_id = ?", 2, "9527").First(&AuthorityMenus{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
