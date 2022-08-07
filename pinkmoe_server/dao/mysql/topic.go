/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:28
 * @FilePath: /pinkmoe_server/dao/mysql/topic.go
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

func GetTopicList(info request.SearchTopicParams) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdTopic{})

	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}

	if info.Value != "" {
		db = db.Where("value LIKE ?", "%"+info.Value+"%")
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorTopicListGet, nil, 0
	}

	var topic []model.XdTopic
	if err = db.Limit(limit).Offset(offset).Find(&topic).Error; err != nil {
		return response.ErrorTopicListGet, nil, 0
	}
	return err, topic, total
}

func GetTopicByTopicValues(topicValues []string) (err error, list []model.XdTopic) {
	for _, topicValue := range topicValues {
		var topic model.XdTopic
		if err := global.XD_DB.Where("value = ?", topicValue).First(&topic).Error; err != nil {
			return response.ErrorTopicListGet, nil
		}
		list = append(list, topic)
	}
	return err, list
}

func CreateTopic(p model.XdTopic) (err error) {
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorTopicCreate
	}
	return err
}

func UpdateTopic(u model.XdTopic) (err error) {
	if err = global.XD_DB.Where("id = ?", u.ID).Updates(&model.XdTopic{
		Value: u.Value,
		Label: u.Label,
		Icon:  u.Icon,
		Sort:  u.Sort,
	}).Error; err != nil {
		return response.ErrorTopicUpdate
	}
	return err
}

func DeleteTopic(p model.XdTopic) (err error) {
	var topic model.XdTopicRelation
	if err = global.XD_DB.Where("xd_topic_id = ?", p.ID).First(&topic).Error; topic.XdTopicId == p.ID {
		return response.ErrorPostTopicDelete
	}
	if err = global.XD_DB.Delete(&p).Error; err != nil {
		return response.ErrorTopicDelete
	}
	return
}
