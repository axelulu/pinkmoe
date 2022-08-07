/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:16
 * @FilePath: /pinkmoe_server/model/xd_post_collect.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdPostCollect struct {
	ID             uint          `gorm:"primarykey"` // 主键ID
	CreatedAt      global.XdTime // 创建时间
	UpdatedAt      global.XdTime // 更新时间
	PostId         string        `json:"postId" form:"postId" gorm:"not null;comment:文章ID"` // 文章ID
	Uuid           uuid.UUID     `json:"uuid" form:"uuid" gorm:"comment:用户ID"`              // 用户ID
	UuidRelation   XdUser        `gorm:"foreignKey:Uuid;references:UUID"`
	PostIdRelation XdPost        `gorm:"foreignKey:PostId;references:PostId"`
}
