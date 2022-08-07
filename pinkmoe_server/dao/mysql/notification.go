/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 10:42:53
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:07
 * @FilePath: /pinkmoe_server/dao/mysql/notification.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"

	uuid "github.com/satori/go.uuid"
)

func GetNotificationList(info *request.SearchNotificationParams, userId uuid.UUID) (err error, post []model.XdNotification, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdNotification{})

	db = db.Where("uuid = ?", userId)

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorGetNotification, nil, 0
	}
	if err = db.Limit(limit).Offset(offset).Order("updated_at desc").Preload("CommentUidRelation").Preload("UserRelation").Preload("CommentRelation").Preload("PostRelation").Find(&post).Error; err != nil {
		return response.ErrorGetNotification, nil, 0
	}
	return
}

func CreateNotification(p *model.XdNotification) (err error) {
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorCreateNotification
	}
	return err
}
