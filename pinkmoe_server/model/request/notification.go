/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:40
 * @FilePath: /pinkmoe_server/model/request/notification.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
