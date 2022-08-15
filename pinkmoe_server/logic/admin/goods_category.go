/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:05
 * @FilePath: /pinkmoe_server/logic/admin/category.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package adminLogic

import (
	"server/dao/mysql"
	"server/model"
	"server/model/request"
)

func GetGoodsCategoryList(p request.SearchGoodsCategoryParams) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetGoodsCategoryList(p)
	return
}

func CreateGoodsCategory(p model.XdGoodsCategory) (err error) {
	err = mysql.CreateGoodsCategory(p)
	return
}

func UpdateGoodsCategory(p model.XdGoodsCategory) (err error) {
	err = mysql.UpdateGoodsCategory(p)
	return
}

func DeleteGoodsCategory(p model.XdGoodsCategory) (err error) {
	err = mysql.DeleteGoodsCategory(p)
	return
}
