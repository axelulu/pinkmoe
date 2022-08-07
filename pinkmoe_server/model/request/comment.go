/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:18
 * @FilePath: /pinkmoe_server/model/request/comment.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

// SearchCommentParams api分页条件查询及排序结构体
type SearchCommentParams struct {
	PageInfo
	PostId   string `json:"postId" form:"postId" gorm:"comment:评论文章ID"` // 评论文章ID
	FormUid  string `json:"formUid" form:"formUid" gorm:"comment:评论作者"` // 评论作者
	Content  string `json:"content" form:"content" gorm:"comment:评论内容"` // 评论内容
	Type     string `json:"type" form:"type" gorm:"comment:评论类型"`       // 评论类型
	Status   string `json:"status" form:"status" gorm:"comment:评论状态"`   // 评论状态
	OrderKey string `json:"orderKey"`                                   // 排序
	Desc     bool   `json:"desc"`                                       // 排序方式:升序false(默认)|降序true
}
