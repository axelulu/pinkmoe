/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:18
 * @FilePath: /pinkmoe_server/model/xd_authority.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import (
	"server/global"
)

type XdAuthority struct {
	CreatedAt       global.XdTime       // 创建时间
	UpdatedAt       global.XdTime       // 更新时间
	DeletedAt       *global.XdTime      `sql:"index"`
	AuthorityId     string              `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityWeight int                 `json:"authorityWeight" gorm:"comment:角色权重"`                                 // 角色权重
	VipStart        int                 `json:"vipStart" gorm:"comment:允许开通会员"`                                      // 允许开通会员
	AuthorityName   string              `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	AuthorityColor  string              `json:"authorityColor" gorm:"comment:角色颜色标识"`                                // 角色颜色标识
	DataAuthorityId []XdAuthority       `json:"dataAuthorityId" gorm:"many2many:xd_data_authority_id"`
	XdBaseMenus     []XdBaseMenu        `json:"menus" gorm:"many2many:xd_authority_menus;"`
	AuthorityParams []XdAuthorityParams `json:"authorityParams" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色扩展参数"`
}
