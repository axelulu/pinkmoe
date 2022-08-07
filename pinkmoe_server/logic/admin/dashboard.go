/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:15
 * @FilePath: /pinkmoe_server/logic/admin/dashboard.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package adminLogic

import "server/model/response"

func ConsoleGet() (console response.Console, err error) {
	// 访问量
	console.Visits = response.Visits{
		Amount:    13555,
		DayVisits: 15,
		Decline:   452,
		Rise:      546456,
	}
	// 销售额
	console.Saleroom = response.Saleroom{
		Amount:       546,
		Degree:       345,
		WeekSaleroom: 234,
	}
	// 订单量
	console.OrderLarge = response.OrderLarge{
		Amount:    342,
		Decline:   1231,
		Rise:      123,
		WeekLarge: 121,
	}
	// 成交额
	console.Volume = response.Volume{
		Amount:    11,
		Decline:   111,
		Rise:      1111,
		WeekLarge: 1111,
	}
	return
}
