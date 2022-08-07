/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:48
 * @FilePath: /pinkmoe_server/initialize/initDB/authority_params.go
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

var AuthorityParams = new(authorityParams)

type authorityParams struct{}

func (a *authorityParams) TableName() string {
	return "xd_authority_params"
}

func (a *authorityParams) Initialize() error {
	entities := []model.XdAuthorityParams{
		{AuthorityId: "2333", AuthorityParamId: "dayPrice", AuthorityParamKey: "日会员", AuthorityParamValue: 1, AuthorityParamDay: 1},
		{AuthorityId: "2333", AuthorityParamId: "weekPrice", AuthorityParamKey: "周会员", AuthorityParamValue: 7, AuthorityParamDay: 7},
		{AuthorityId: "2333", AuthorityParamId: "monthPrice", AuthorityParamKey: "月会员", AuthorityParamValue: 30, AuthorityParamDay: 30},
		{AuthorityId: "2333", AuthorityParamId: "yearPrice", AuthorityParamKey: "年会员", AuthorityParamValue: 365, AuthorityParamDay: 365},
		{AuthorityId: "2333", AuthorityParamId: "download", AuthorityParamKey: "下载折扣", AuthorityParamValue: 100, AuthorityParamDay: 0},
		{AuthorityId: "2333", AuthorityParamId: "music", AuthorityParamKey: "音乐折扣", AuthorityParamValue: 100, AuthorityParamDay: 0},
		{AuthorityId: "2333", AuthorityParamId: "video", AuthorityParamKey: "视频折扣", AuthorityParamValue: 100, AuthorityParamDay: 0},
		{AuthorityId: "9527", AuthorityParamId: "dayPrice", AuthorityParamKey: "日会员", AuthorityParamValue: 10, AuthorityParamDay: 1},
		{AuthorityId: "9527", AuthorityParamId: "weekPrice", AuthorityParamKey: "周会员", AuthorityParamValue: 60, AuthorityParamDay: 7},
		{AuthorityId: "9527", AuthorityParamId: "monthPrice", AuthorityParamKey: "月会员", AuthorityParamValue: 100, AuthorityParamDay: 30},
		{AuthorityId: "9527", AuthorityParamId: "yearPrice", AuthorityParamKey: "年会员", AuthorityParamValue: 400, AuthorityParamDay: 365},
		{AuthorityId: "9527", AuthorityParamId: "download", AuthorityParamKey: "下载折扣", AuthorityParamValue: 80, AuthorityParamDay: 0},
		{AuthorityId: "9527", AuthorityParamId: "music", AuthorityParamKey: "音乐折扣", AuthorityParamValue: 80, AuthorityParamDay: 0},
		{AuthorityId: "9527", AuthorityParamId: "video", AuthorityParamKey: "视频折扣", AuthorityParamValue: 80, AuthorityParamDay: 0},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil {
		return errors.Wrapf(err, "%s表数据初始化失败!", a.TableName())
	}
	return nil
}

func (a *authorityParams) CheckDataExist() bool {
	if errors.Is(global.XD_DB.Where("authority_id = ?", "9527").First(&model.XdAuthorityParams{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
