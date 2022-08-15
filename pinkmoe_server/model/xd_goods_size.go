/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:32
 * @FilePath: /pinkmoe_server/model/xd_post.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import (
	"server/global"
)

type XdGoodsSize struct {
	CreatedAt global.XdTime  // 创建时间
	UpdatedAt global.XdTime  // 更新时间
	DeletedAt *global.XdTime `gorm:"index" json:"-"`                                                         // 删除时间
	Value     string         `json:"value" form:"value" gorm:"comment:规格配置"`                                 // 规格配置
	GoodsId   string         `json:"goodsId" form:"goodsId" gorm:"not null;unique;primary_key;comment:商品ID"` // 商品ID
}
