/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:52
 * @FilePath: /pinkmoe_server/model/request/upload.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

import uuid "github.com/satori/go.uuid"

// DeleteImgFileParams api分页条件查询及排序结构体
type DeleteImgFileParams struct {
	Url string `json:"url" form:"url" gorm:"comment:文件地址"` // 文件地址
}

type XdUploadFileParams struct {
	PostId string    `json:"postId" form:"postId" gorm:"comment:文章ID;default:null"` // 文章ID
	Uuid   uuid.UUID `json:"uuid" form:"uuid" gorm:"comment:用户ID"`                  // 用户ID
	Type   string    `json:"type" form:"type" gorm:"comment:图片类型;default:'post'"`   // 图片类型
	PageInfo
	OrderKey string `json:"orderKey" form:"orderKey"` // 排序
	Desc     bool   `json:"desc" form:"desc"`         // 排序方式:升序false(默认)|降序true
}
