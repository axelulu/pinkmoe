/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:20
 * @FilePath: /pinkmoe_server/dao/mysql/setting.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"server/global"
	"server/model"
	"server/model/response"
)

func GetSettingItem(p model.XdSetting) (err error, setting model.XdSetting) {
	db := global.XD_DB.Model(&model.XdSetting{})
	if p.Slug != "" {
		db = db.Where("slug = ?", p.Slug)
	}

	if err = db.First(&setting).Error; err != nil {
		return response.ErrorSettingGet, setting
	}
	return err, setting
}

func CreateSetting(p *model.XdSetting) (err error) {
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorSettingCreate
	}
	return err
}

func UpdateSetting(p *model.XdSetting) (err error) {
	if err = global.XD_DB.Model(&model.XdSetting{}).Where("slug = ?", p.Slug).Update("value", p.Value).Error; err != nil {
		return response.ErrorSettingUpdate
	}
	return err
}
