/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:12
 * @FilePath: /pinkmoe_server/dao/mysql/operation.go
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
)

func GetOperationList(info request.SearchOperationParams) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdOperationLog{})

	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}

	if info.Ip != "" {
		db = db.Where("ip = ?", info.Ip)
	}

	if info.UserID != "" {
		db = db.Where("user_id = ?", info.UserID)
	}

	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorOperationListGet, nil, 0
	}

	var topic []model.XdOperationLog
	if err = db.Limit(limit).Offset(offset).Find(&topic).Error; err != nil {
		return response.ErrorOperationListGet, nil, 0
	}
	return err, topic, total
}

func CreateOperation(p *model.XdOperationLog) (err error) {
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorOperationCreate
	}
	return err
}

func UpdateOperation(p *model.XdOperationLog) (err error) {
	if err = global.XD_DB.Where("id = ?", p.ID).First(&model.XdOperationLog{}).Updates(&model.XdOperationLog{
		Ip:           p.Ip,
		Method:       p.Method,
		Path:         p.Path,
		Status:       p.Status,
		Latency:      p.Latency,
		Agent:        p.Agent,
		ErrorMessage: p.ErrorMessage,
		Body:         p.Body,
		Resp:         p.Resp,
		UserID:       p.UserID,
		User:         p.User,
	}).Error; err != nil {
		return response.ErrorOperationUpdate
	}
	return err
}

func DeleteOperation(p *model.XdOperationLog) (err error) {
	if err = global.XD_DB.Delete(&p).Error; err != nil {
		return response.ErrorOperationDelete
	}
	return
}
