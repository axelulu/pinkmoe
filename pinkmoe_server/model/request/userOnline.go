/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:57
 * @FilePath: /pinkmoe_server/model/request/userOnline.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

// SearchUserOnlineParams api分页条件查询及排序结构体
type SearchUserOnlineParams struct {
	PageInfo
	UserName string `json:"userName" form:"userName" gorm:"comment:用户名"` // 用户名
	Ip       string `json:"ip" form:"ip" gorm:"comment:登录ip"`            // 登录ip
	Explorer string `json:"explorer" form:"explorer" gorm:"comment:浏览器"` // 浏览器
}
