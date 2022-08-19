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

import (
	"server/dao/mysql"
	"server/model/response"
)

func ConsoleGet() (console response.Console, err error) {
	day := mysql.GetLoginLogByTime(1)
	twoDay := mysql.GetLoginLogByTime(2)
	week := mysql.GetLoginLogByTime(7)
	twoWeek := mysql.GetLoginLogByTime(14)
	amount := mysql.GetLoginLogByTime(9999)

	// 访问量
	if week == 0 || day == 0 {
		if day == 0 {
			console.Visits = response.Visits{
				Amount:    int(amount),
				DayVisits: int(day),
				Decline:   float64((twoWeek - week) / week),
				Rise:      float64((twoDay - day) / 1),
			}
		}
		if week == 0 {
			console.Visits = response.Visits{
				Amount:    int(amount),
				DayVisits: int(day),
				Decline:   float64((twoWeek - week) / 1),
				Rise:      float64((twoDay - day) / day),
			}
		}
	} else {
		console.Visits = response.Visits{
			Amount:    int(amount),
			DayVisits: int(day),
			Decline:   float64((twoWeek - week) / week),
			Rise:      float64((twoDay - day) / day),
		}
	}

	_, WeekSale := mysql.GetOrderByTime(7, "download")
	_, TwoWeekSale := mysql.GetOrderByTime(14, "download")
	_, Amount := mysql.GetOrderByTime(9999, "download")
	// 销售额
	if WeekSale == 0 {
		console.Saleroom = response.Saleroom{
			Amount:       Amount,
			Degree:       float64((TwoWeekSale - WeekSale) / 1),
			WeekSaleroom: WeekSale,
		}
	} else {
		console.Saleroom = response.Saleroom{
			Amount:       Amount,
			Degree:       float64((TwoWeekSale - WeekSale) / WeekSale),
			WeekSaleroom: WeekSale,
		}
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

	today, yesterday := mysql.GetLoginLogTrend()
	// 流量趋势
	console.FluxTrend = response.FluxTrend{
		Today:     today,
		Yesterday: yesterday,
	}
	return
}
