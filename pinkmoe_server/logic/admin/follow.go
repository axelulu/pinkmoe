/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:18
 * @FilePath: /pinkmoe_server/logic/admin/follow.go
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

	uuid "github.com/satori/go.uuid"
)

func GetFollowList(p request.SearchFollowParams) (err error, list []model.XdFollow, total int64) {
	err, list, total = mysql.GetFollowList(&p)
	return
}

func GetFollowStatus(p request.ReqFollow, userId uuid.UUID) (err error, isFollow bool) {
	fromString, _ := uuid.FromString(p.ToUid)
	if fromString == userId {
		return response.ErrorUserFolloweOwn, true
	}
	err, isFollow = mysql.GetFollowStatus(fromString, userId)
	return
}

func CreateFollow(p request.ReqFollow, userId uuid.UUID) (err error) {
	fromString, _ := uuid.FromString(p.ToUid)
	if fromString == userId {
		return response.ErrorUserFolloweOwn
	}
	err = mysql.CreateFollow(fromString, userId)
	return
}

func DeleteFollow(p request.ReqFollow, userId uuid.UUID) (err error) {
	fromString, _ := uuid.FromString(p.ToUid)
	if fromString == userId {
		return response.ErrorUserUnFolloweOwn
	}
	err = mysql.DeleteFollow(fromString, userId)
	return
}
