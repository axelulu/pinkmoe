/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:15
 * @FilePath: /pinkmoe_server/model/request/chat.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

// SearchChatParams api分页条件查询及排序结构体
type SearchChatParams struct {
	ReqChatRelation
	PageInfo
	OrderKey string `json:"orderKey" form:"orderKey"` // 排序
	Desc     bool   `json:"desc" form:"desc"`         // 排序方式:升序false(默认)|降序true
}

// ReqChatRelation api分页条件查询及排序结构体
type ReqChatRelation struct {
	SendId string `json:"sendId" form:"sendId"` // 排序
	UserId string `json:"userId" form:"userId"` // 排序
}

// SearchChatRelationParams api分页条件查询及排序结构体
type SearchChatRelationParams struct {
	ReqChatRelation
	PageInfo
	OrderKey string `json:"orderKey" form:"orderKey"` // 排序
	Desc     bool   `json:"desc" form:"desc"`         // 排序方式:升序false(默认)|降序true
}
