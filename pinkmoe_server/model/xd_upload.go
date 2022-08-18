/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:50
 * @FilePath: /pinkmoe_server/model/xd_upload.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdUploadFile struct {
	global.XD_MODEL
	Name    string    `json:"name" gorm:"comment:文件名"`                  // 文件名
	Url     string    `json:"url" gorm:"comment:文件地址"`                  // 文件地址
	Tag     string    `json:"tag" gorm:"comment:文件标签"`                  // 文件标签
	Key     string    `json:"key" gorm:"comment:编号"`                    // 编号
	PostId  string    `json:"postId" gorm:"comment:文章ID;default:null"`  // 文章ID
	GoodsId string    `json:"goodsId" gorm:"comment:商品ID;default:null"` // 商品ID
	Type    string    `json:"type" gorm:"comment:图片类型;default:'post'"`  // 图片类型
	Uuid    uuid.UUID `json:"uuid" gorm:"comment:用户ID"`                 // 用户ID
}
