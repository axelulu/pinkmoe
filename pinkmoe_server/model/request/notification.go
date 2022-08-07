/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:40
 * @FilePath: /pinkmoe_server/model/request/notification.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type ReqNotification struct {
	global.XD_MODEL
	Uuid      uuid.UUID `json:"uuid" gorm:"comment:用户ID"`      // 用户ID
	PostId    string    `json:"postId" gorm:"comment:文章ID"`    // 文章ID
	CommentId string    `json:"commentId" gorm:"comment:评论ID"` // 评论ID
	Type      string    `json:"type" gorm:"comment:通知类型"`      // 通知类型
}

// SearchNotificationParams api分页条件查询及排序结构体
type SearchNotificationParams struct {
	ReqNotification
	PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
