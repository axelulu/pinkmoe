/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:45
 * @FilePath: /pinkmoe_server/model/xd_comment.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdComment struct {
	global.XD_MODEL
	PostId          string      `json:"postId" form:"postId" gorm:"comment:评论文章ID"`            // 评论文章ID
	ParentId        uint        `json:"parentId" form:"parentId" gorm:"comment:父评论ID"`         // 父评论ID
	FormUid         uuid.UUID   `json:"formUid" form:"formUid" gorm:"comment:评论作者"`            // 评论作者
	ToUid           uuid.UUID   `json:"toUid" form:"toUid" gorm:"default:null;comment:评论目标用户"` // 评论目标用户
	Content         string      `json:"content" form:"content" gorm:"comment:评论内容"`            // 评论内容
	Type            string      `json:"type" form:"type" gorm:"comment:评论类型"`                  // 评论类型
	Status          string      `json:"status" form:"status" gorm:"comment:评论状态"`              // 评论状态
	Children        []XdComment `json:"children" gorm:"-"`
	FormUidRelation XdUser      `gorm:"foreignKey:FormUid;references:UUID"`
	ToUidRelation   XdUser      `gorm:"foreignKey:ToUid;references:UUID"`
}
