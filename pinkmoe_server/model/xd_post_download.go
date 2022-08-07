/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:19
 * @FilePath: /pinkmoe_server/model/xd_post_download.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import "server/global"

type XdPostDownload struct {
	ID         uint          `gorm:"primarykey"` // 主键ID
	CreatedAt  global.XdTime // 创建时间
	UpdatedAt  global.XdTime // 更新时间
	PostId     string        `json:"postId" form:"postId" gorm:"not null;comment:文章ID"` // 文章ID
	Price      int           `json:"price" form:"price" gorm:"comment:链接价格"`            // 链接价格
	PriceType  string        `json:"priceType" form:"priceType" gorm:"comment:链接价格类型"`  // 链接价格类型
	Name       string        `json:"name" form:"name" gorm:"type:text;comment:下载链接名称"`  // 下载链接名称
	Url        string        `json:"url" form:"url" gorm:"comment:下载链接"`                // 下载链接
	ExtractPwd string        `json:"extractPwd" form:"extractPwd" gorm:"comment:提取密码"`  // 提取密码
	UnpackPwd  string        `json:"unpackPwd" form:"unpackPwd" gorm:"comment:解压密码"`    // 解压密码
}
