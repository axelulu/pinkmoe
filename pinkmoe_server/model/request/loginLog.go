/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:32
 * @FilePath: /pinkmoe_server/model/request/loginLog.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

// SearchLoginLogParams api分页条件查询及排序结构体
type SearchLoginLogParams struct {
	PageInfo
	UserName string `json:"userName" form:"userName" gorm:"comment:登录账号"`  // 登录账号
	Ip       string `json:"ip" form:"ip" gorm:"comment:登录IP地址"`            // 登录IP地址
	Explorer string `json:"explorer" form:"explorer" gorm:"comment:浏览器类型"` // 浏览器类型
}
