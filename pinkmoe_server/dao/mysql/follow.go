/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:50
 * @FilePath: /pinkmoe_server/dao/mysql/follow.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func GetFollowList(info *request.SearchFollowParams) (err error, post []model.XdFollow, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdFollow{})

	if info.ToUid != "" {
		db = db.Where("to_uid = ?", info.ToUid)
	}

	if info.FormUid != "" {
		db = db.Where("form_uid = ?", info.FormUid)
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorUserFollowListGet, nil, 0
	}
	if err = db.Limit(limit).Offset(offset).Preload("FormUidRelation").Preload("ToUidRelation").Find(&post).Error; err != nil {
		return response.ErrorUserFollowListGet, nil, 0
	}
	return
}

func GetFollowCount(uuid uuid.UUID) (err error, total int64) {
	if err = global.XD_DB.Model(model.XdFollow{}).Where("form_uid = ?", uuid).Count(&total).Error; err != nil {
		return response.ErrorUserFollowCountGet, 0
	}
	return
}

func GetFansCount(uuid uuid.UUID) (err error, total int64) {
	if err = global.XD_DB.Model(model.XdFollow{}).Where("to_uid = ?", uuid).Count(&total).Error; err != nil {
		return response.ErrorUserFansCountGet, 0
	}
	return
}

func GetFollowStatus(followId uuid.UUID, userId uuid.UUID) (err error, isFollow bool) {
	if !errors.Is(global.XD_DB.Where("to_uid = ?", followId).Where("form_uid = ?", userId).First(&model.XdFollow{}).Error, gorm.ErrRecordNotFound) {
		return response.ErrorUserFollowStatusGet, true
	}
	return err, false
}

func CreateFollow(followId uuid.UUID, userId uuid.UUID) (err error) {
	if !errors.Is(global.XD_DB.Where("to_uid = ?", followId).Where("form_uid = ?", userId).First(&model.XdFollow{}).Error, gorm.ErrRecordNotFound) {
		return response.ErrorUserFollowed
	}
	if err = global.XD_DB.Create(&model.XdFollow{
		FormUid: userId,
		ToUid:   followId,
	}).Error; err != nil {
		return response.ErrorUserFollow
	}
	return err
}

func DeleteFollow(followId uuid.UUID, userId uuid.UUID) (err error) {
	if err = global.XD_DB.Where("to_uid = ?", followId).Where("form_uid = ?", userId).Delete(&model.XdFollow{}).Error; err != nil {
		return response.ErrorUserUnFollow
	}
	return err
}
