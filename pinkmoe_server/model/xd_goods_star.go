/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:26
 * @FilePath: /pinkmoe_server/model/xd_post_star.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdGoodsStar struct {
	ID        uint          `gorm:"primarykey"` // 主键ID
	CreatedAt global.XdTime // 创建时间
	UpdatedAt global.XdTime // 更新时间
	GoodsId   string        `json:"goodsId" form:"goodsId" gorm:"not null;comment:商品ID"` // 商品ID
	Uuid      uuid.UUID     `json:"uuid" gorm:"comment:用户ID"`                            // 用户ID
}
