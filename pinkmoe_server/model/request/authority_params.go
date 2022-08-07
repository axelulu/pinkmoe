/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 11:49:20
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:06
 * @FilePath: /pinkmoe_server/model/request/authority_params.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

type XdAuthorityParams struct {
	AuthorityId         string `json:"authorityId" gorm:"comment:角色ID"`          // 角色ID
	AuthorityParamId    string `json:"authorityParamId" gorm:"comment:角色扩展ID"`   // 角色扩展ID
	AuthorityParamKey   string `json:"authorityParamKey" gorm:"comment:角色扩展名"`   // 角色扩展名
	AuthorityParamValue int    `json:"authorityParamValue" gorm:"comment:角色扩展值"` // 角色扩展值
	AuthorityParamDay   int    `json:"authorityParamDay" gorm:"comment:角色扩展值2"`  // 角色扩展值2
	AuthorityNum        int    `json:"authorityNum" gorm:"comment:角色开通数量"`       // 角色开通数量
}
