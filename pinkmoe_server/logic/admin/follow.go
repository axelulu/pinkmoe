/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:18
 * @FilePath: /pinkmoe_server/logic/admin/follow.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
