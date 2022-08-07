/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:29
 * @FilePath: /pinkmoe_server/model/request/key.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

// SearchKeyParams api分页条件查询及排序结构体
type SearchKeyParams struct {
	PageInfo
	Key    string `json:"key" form:"key" gorm:"comment:卡密内容"`       // 卡密内容
	Value  int    `json:"value" form:"value" gorm:"comment:卡密面额"`   // 卡密面额
	Num    int    `json:"num" form:"num" gorm:"comment:卡密面额"`       // 卡密面额
	Type   string `json:"type" form:"type" gorm:"comment:卡密类型"`     // 卡密类型
	Status string `json:"status" form:"status" gorm:"comment:卡密状态"` // 卡密状态
}
