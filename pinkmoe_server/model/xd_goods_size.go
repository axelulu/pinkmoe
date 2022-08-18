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

import "server/global"

type XdGoodsSize struct {
	global.XD_MODEL
	Key     string `json:"key" form:"key" gorm:"comment:规格名称"`                  // 规格名称
	Value   string `json:"value" form:"value" gorm:"comment:规格配置"`              // 规格配置
	GoodsId string `json:"goodsId" form:"goodsId" gorm:"not null;comment:商品ID"` // 商品ID
}
