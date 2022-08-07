/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:42
 * @FilePath: /pinkmoe_server/model/request/operation.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

// SearchOperationParams api分页条件查询及排序结构体
type SearchOperationParams struct {
	PageInfo
	Ip     string `json:"ip" form:"ip" gorm:"comment:请求ip"`         // 请求ip
	Method string `json:"method" form:"method" gorm:"comment:请求方法"` // 请求方法
	Status int    `json:"status" form:"status" gorm:"comment:请求状态"` // 请求状态
	UserID string `json:"userId" form:"userId" gorm:"comment:用户id"` // 用户id
}
