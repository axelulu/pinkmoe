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

func GetCategoryList(p request.SearchCategoryParams) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetCategoryList(p)
	return
}

func CreateCategory(p model.XdCategory) (err error) {
	err = mysql.CreateCategory(p)
	return
}

func UpdateCategory(p model.XdCategory) (err error) {
	err = mysql.UpdateCategory(p)
	return
}

func DeleteCategory(p model.XdCategory) (err error) {
	err = mysql.DeleteCategory(p)
	return
}
