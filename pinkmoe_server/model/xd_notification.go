/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:05
 * @FilePath: /pinkmoe_server/model/xd_notification.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdNotification struct {
	global.XD_MODEL
	Uuid       uuid.UUID `json:"uuid" gorm:"comment:用户ID;default:null"`         // 用户ID
	CommentUid uuid.UUID `json:"commentUid" gorm:"comment:评论用户ID;default:null"` // 评论用户ID
	PostId     string    `json:"postId" gorm:"comment:文章ID;default:null"`       // 文章ID
	CommentId  uint      `json:"commentId" gorm:"comment:评论ID;default:null"`    // 评论ID
	Credit     int       `json:"credit" gorm:"comment:奖励积分"`                    // 奖励积分
	Cash       int       `json:"cash" gorm:"comment:奖励现金"`                      // 奖励现金
	Type       string    `json:"type" gorm:"comment:通知类型"`                      // 通知类型
	//PayType            string    `json:"payType" gorm:"comment:充值类型"`                   // 充值类型
	Msg                string    `json:"msg" gorm:"comment:通知信息"`          // 通知信息
	Read               int       `json:"read" gorm:"comment:读取;default:0"` // 读取
	PostRelation       XdPost    `json:"postRelation" gorm:"foreignKey:PostId;references:PostId"`
	CommentRelation    XdComment `json:"commentRelation" gorm:"foreignKey:CommentId;references:ID"`
	UserRelation       XdUser    `json:"userRelation" gorm:"foreignKey:Uuid;references:UUID"`
	CommentUidRelation XdUser    `json:"commentUidRelation" gorm:"foreignKey:CommentUid;references:UUID"`
}
