/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:42
 * @FilePath: /pinkmoe_server/model/xd_comment_smile.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import "server/global"

type XdCommentSmile struct {
	global.XD_MODEL
	Name string `json:"name" form:"name" gorm:"comment:评论表情名称"` // 评论表情名称
	Url  string `json:"url" form:"url" gorm:"comment:评论表情链接"`   // 评论表情链接
	Type string `json:"type" form:"type" gorm:"comment:评论表情类型"` // 评论表情类型
}
