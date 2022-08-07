/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-12 20:47:11
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:34
 * @FilePath: /pinkmoe_server/logic/admin/notification.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package adminLogic

import (
	"server/dao/mysql"
	"server/global"
	"server/model"
	"server/model/request"

	uuid "github.com/satori/go.uuid"
)

func GetNotificationList(p request.SearchNotificationParams, userId uuid.UUID) (err error, list []model.XdNotification, total int64) {
	err, list, total = mysql.GetNotificationList(&p, userId)
	return
}

func CreateNotification(uuids uuid.UUID, commentUid string, postId string, credit int, cash int, commentId uint, types string, msg string) {
	commentUids, _ := uuid.FromString(commentUid)
	if commentId == 0 {
		_ = mysql.CreateNotification(&model.XdNotification{
			XD_MODEL: global.XD_MODEL{},
			Uuid:     uuids,
			PostId:   postId,
			Credit:   credit,
			Cash:     cash,
			Type:     types,
			Msg:      msg,
		})
	} else if postId == "" {
		_ = mysql.CreateNotification(&model.XdNotification{
			XD_MODEL:   global.XD_MODEL{},
			Uuid:       uuids,
			CommentId:  commentId,
			CommentUid: commentUids,
			Credit:     credit,
			Cash:       cash,
			Type:       types,
			Msg:        msg,
		})
	} else if commentUid == "" {
		_ = mysql.CreateNotification(&model.XdNotification{
			XD_MODEL:  global.XD_MODEL{},
			Uuid:      uuids,
			CommentId: commentId,
			Credit:    credit,
			Cash:      cash,
			Type:      types,
			Msg:       msg,
		})
	} else {
		_ = mysql.CreateNotification(&model.XdNotification{
			XD_MODEL:   global.XD_MODEL{},
			Uuid:       uuids,
			CommentUid: commentUids,
			PostId:     postId,
			CommentId:  commentId,
			Credit:     credit,
			Cash:       cash,
			Type:       types,
			Msg:        msg,
		})
	}
	return
}
