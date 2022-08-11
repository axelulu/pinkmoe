/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:02
 * @FilePath: /pinkmoe_server/logic/admin/authority.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package adminLogic

import (
	"server/dao/mysql"
	"server/model"
	"server/model/request"
	"server/model/response"
)

func GetAdminAuthorityList(p request.PageInfo) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetAuthorityList(p, false)
	return
}

func GetAuthorityList(p request.PageInfo) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetAuthorityList(p, true)
	return
}

func GetMenuAuthority(p request.GetAuthorityId) (err error, menus []model.XdMenu) {
	err, menus = mysql.GetMenuAuthority(p)
	return
}

func CreateAuthority(p model.XdAuthority) (err error) {
	err, authority := mysql.CreateAuthority(p)
	if err != nil {
		return response.ErrorAuthorityCreate
	}
	var auth model.XdAuthority
	auth.AuthorityId = authority.AuthorityId
	auth.XdBaseMenus = mysql.DefaultMenu()
	if err = mysql.UpdateMenuAuthority(auth); err != nil {
		return response.ErrorAuthorityMenuCreate
	}
	err = mysql.UpdateCasbin(authority.AuthorityId, mysql.DefaultCasbin())
	return
}

func UpdateMenuAuthority(p request.UpdateMenuAuthorityInfo) (err error) {
	var auth model.XdAuthority
	auth.AuthorityId = p.AuthorityId
	auth.XdBaseMenus = p.Menus
	err = mysql.UpdateMenuAuthority(auth)
	return
}

func UpdateAuthority(p model.XdAuthority) (err error) {
	u := model.XdAuthority{AuthorityName: p.AuthorityName, VipStart: p.VipStart, AuthorityWeight: p.AuthorityWeight, AuthorityColor: p.AuthorityColor}
	mysql.DeleteAuthorityParams(p.AuthorityId)
	for _, param := range p.AuthorityParams {
		mysql.CreateAuthorityParams(&model.XdAuthorityParams{
			AuthorityId:         p.AuthorityId,
			AuthorityParamId:    param.AuthorityParamId,
			AuthorityParamKey:   param.AuthorityParamKey,
			AuthorityParamValue: param.AuthorityParamValue,
			AuthorityParamDay:   param.AuthorityParamDay,
		})
	}
	err = mysql.UpdateAuthority(u, p.AuthorityId)
	return
}

func DeleteAuthority(p model.XdAuthority) (err error) {
	err = mysql.DeleteAuthority(p)
	return
}
