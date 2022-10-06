/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:58
 * @FilePath: /pinkmoe_server/dao/mysql/key.go
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
)

func GetKeyList(info request.SearchKeyParams) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdKey{})

	if info.Key != "" {
		db = db.Where("label = ?", info.Key)
	}

	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}

	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}

	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorKeyListGet, nil, 0
	}

	var key []model.XdKey
	if err = db.Limit(limit).Offset(offset).Find(&key).Error; err != nil {
		return response.ErrorKeyListGet, nil, 0
	}
	return err, key, total
}

func GetKeyItem(info *model.XdKey) (err error, key model.XdKey) {
	db := global.XD_DB.Model(&model.XdKey{})

	if info.Key != "" {
		db = db.Where("label = ?", info.Key)
	}

	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}

	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}

	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}

	if err = db.First(&key).Error; err != nil {
		return response.ErrorKeyListGet, key
	}
	return err, key
}

func UpdateKeyItem(u *model.XdKey) (err error) {
	db := global.XD_DB.Model(&model.XdKey{})
	if err = db.Where("ID = ?", u.ID).Update("status", u.Status).Error; err != nil {
		return response.ErrorKeyUpdate
	}
	return err
}

func CreateKey(p *model.XdKey) (err error) {
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorKeyCreate
	}
	return err
}

func DeleteKey(p model.XdKey) (err error) {
	if err = global.XD_DB.Delete(&p).Error; err != nil {
		return response.ErrorKeyDelete
	}
	return
}
