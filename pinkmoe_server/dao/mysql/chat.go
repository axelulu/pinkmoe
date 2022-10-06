/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:46
 * @FilePath: /pinkmoe_server/dao/mysql/chat.go
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

func GetChatList(p request.SearchChatParams) (err error, chat []model.XdChat, total int64) {
	limit := p.PageSize
	offset := p.PageSize * (p.Page - 1)
	db := global.XD_DB.Model(&model.XdChat{})

	db = db.Where("(send_id = ? AND user_id = ?) OR (send_id = ? AND user_id = ?)", p.SendId, p.UserId, p.UserId, p.SendId)

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorChatListGet, nil, 0
	}
	if err = db.Limit(limit).Offset(offset).Preload("SendIdRelation").Preload("UserIdRelation").Preload("UserIdRelation.Authority").Find(&chat).Error; err != nil {
		return response.ErrorChatListGet, nil, 0
	}
	return err, chat, 0
}

func GetChatRelationList(p request.SearchChatRelationParams) (err error, chatRelation []model.XdChatRelation, total int64) {
	limit := p.PageSize
	offset := p.PageSize * (p.Page - 1)
	db := global.XD_DB.Model(&model.XdChatRelation{})

	db = db.Where("user_id = ?", p.UserId)

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorChatListGet, nil, 0
	}
	if err = db.Limit(limit).Offset(offset).Preload("SendIdRelation").Preload("UserIdRelation").Find(&chatRelation).Error; err != nil {
		return response.ErrorChatListGet, nil, 0
	}
	return err, chatRelation, 0
}

func GetChatRelationByUserId(userId uuid.UUID) (err error, chatRelation []model.XdChatRelation) {
	db := global.XD_DB.Model(&model.XdChatRelation{})

	db = db.Where("user_id = ?", userId)

	if err = db.Find(&chatRelation).Error; err != nil {
		return response.ErrorChatRelationGet, chatRelation
	}
	return err, chatRelation
}

func CreateChatRelation(p *model.XdChatRelation) (err error) {
	if p.SendId == p.UserId {
		return response.ErrorSettingCreate
	}
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorChatRelationCreate
	}
	return err
}

func DeleteChatRelation(userId uuid.UUID, sendId uuid.UUID) (err error) {
	if err = global.XD_DB.Where("user_id = ?", userId).Where("send_id = ?", sendId).Delete(&model.XdChatRelation{}).Error; err != nil {
		return response.ErrorChatRelationDelete
	}
	return err
}

func CreateChat(p *model.XdChat) (err error, chat model.XdChat) {
	if p.SendId == p.UserId {
		return response.ErrorChatCreate, chat
	}
	chat = model.XdChat{
		UserId:  p.UserId,
		SendId:  p.SendId,
		Cmd:     p.Cmd,
		Media:   p.Media,
		Content: p.Content,
		Pic:     p.Pic,
		Url:     p.Url,
		Memo:    p.Memo,
		Amount:  p.Amount,
	}
	if err = global.XD_DB.Create(&chat).Error; err != nil {
		return response.ErrorChatCreate, chat
	}
	return err, chat
}
