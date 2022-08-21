/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:56
 * @FilePath: /pinkmoe_server/initialize/initDB/category.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package initDB

import (
	"server/global"
	"server/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Category = new(category)

type category struct{}

func (u *category) TableName() string {
	return "xd_categories"
}

func (u *category) Initialize() error {
	entities := []model.XdCategory{
		{Slug: "news", Name: "每日新闻", Icon: "newspaper", ParentId: 0, Sort: 0},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (u *category) CheckDataExist() bool {
	if errors.Is(global.XD_DB.Where("slug = ?", "news").First(&model.XdCategory{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
