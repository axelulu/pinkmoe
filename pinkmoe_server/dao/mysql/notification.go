/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 10:42:53
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:07
 * @FilePath: /pinkmoe_server/dao/mysql/notification.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
