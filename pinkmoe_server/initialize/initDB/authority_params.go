/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:48
 * @FilePath: /pinkmoe_server/initialize/initDB/authority_params.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
		{AuthorityId: "2333", AuthorityParamId: "downloadDiscounts", AuthorityParamKey: "下载折扣", AuthorityParamValue: 100, AuthorityParamDay: 0},
		{AuthorityId: "2333", AuthorityParamId: "musicDiscounts", AuthorityParamKey: "音乐折扣", AuthorityParamValue: 100, AuthorityParamDay: 0},
		{AuthorityId: "2333", AuthorityParamId: "videoDiscounts", AuthorityParamKey: "视频折扣", AuthorityParamValue: 100, AuthorityParamDay: 0},
		{AuthorityId: "9527", AuthorityParamId: "dayPrice", AuthorityParamKey: "日会员", AuthorityParamValue: 10, AuthorityParamDay: 1},
		{AuthorityId: "9527", AuthorityParamId: "weekPrice", AuthorityParamKey: "周会员", AuthorityParamValue: 60, AuthorityParamDay: 7},
		{AuthorityId: "9527", AuthorityParamId: "monthPrice", AuthorityParamKey: "月会员", AuthorityParamValue: 100, AuthorityParamDay: 30},
		{AuthorityId: "9527", AuthorityParamId: "yearPrice", AuthorityParamKey: "年会员", AuthorityParamValue: 400, AuthorityParamDay: 365},
		{AuthorityId: "9527", AuthorityParamId: "downloadDiscounts", AuthorityParamKey: "下载折扣", AuthorityParamValue: 80, AuthorityParamDay: 0},
		{AuthorityId: "9527", AuthorityParamId: "musicDiscounts", AuthorityParamKey: "音乐折扣", AuthorityParamValue: 80, AuthorityParamDay: 0},
		{AuthorityId: "9527", AuthorityParamId: "videoDiscounts", AuthorityParamKey: "视频折扣", AuthorityParamValue: 80, AuthorityParamDay: 0},
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
