/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:02
 * @FilePath: /pinkmoe_server/dao/mysql/loginLog.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"server/global"
	"server/model"
)

func GetOrderByTime(days int, types string) (total int64, price int) {
	db := global.XD_DB.Model(&model.XdOrder{})

	db = db.Where("type = ?", types)

	db = db.Where("to_days(now()) - to_days(updated_at) <= ?", days)

	var orders []model.XdOrder
	if err := db.Find(&orders).Error; err != nil {
		return 0, 0
	}
	price = 0
	for _, order := range orders {
		price += order.Price
	}
	total = int64(len(orders))
	return total, price
}
