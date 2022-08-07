/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:47
 * @FilePath: /pinkmoe_server/model/request/topic.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

// SearchTopicParams api分页条件查询及排序结构体
type SearchTopicParams struct {
	PageInfo
	Value    string `json:"value" form:"value" gorm:"comment:话题标识"` // 话题标识
	Label    string `json:"label" form:"label" gorm:"comment:话题名称"` // 话题名称
	OrderKey string `json:"orderKey" form:"orderKey"`               // 排序
	Desc     bool   `json:"desc" form:"desc"`                       // 排序方式:升序false(默认)|降序true
}
