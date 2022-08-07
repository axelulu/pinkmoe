/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:04
 * @FilePath: /pinkmoe_server/model/xd_address.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdAddress struct {
	CreatedAt global.XdTime  // 创建时间
	UpdatedAt global.XdTime  // 更新时间
	DeletedAt *global.XdTime `gorm:"index" json:"-"`                                            // 删除时间
	AddressId string         `json:"addressId" gorm:"not null;unique;primary_key;comment:订单ID"` // 订单ID
	Address   int64          `json:"address" gorm:"comment:地址信息"`                               // 地址信息
	Uuid      uuid.UUID      `json:"uuid" gorm:"comment:用户ID"`                                  // 用户ID
}
