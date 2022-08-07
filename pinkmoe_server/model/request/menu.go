/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:38
 * @FilePath: /pinkmoe_server/model/request/menu.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

import (
	"server/model"
)

type UpdateMenuAuthorityInfo struct {
	Menus       []model.XdBaseMenu `json:"menus"`
	AuthorityId string             `json:"authorityId"` // 角色ID
}
