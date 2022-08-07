/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:38
 * @FilePath: /pinkmoe_server/logic/admin/operation.go
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

func GetOperationList(p request.SearchOperationParams) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetOperationList(p)
	return
}

func CreateOperation(p model.XdOperationLog) (err error) {
	err = mysql.CreateOperation(&p)
	return
}

func UpdateOperation(p model.XdOperationLog) (err error) {
	err = mysql.UpdateOperation(&p)
	return
}

func DeleteOperation(p model.XdOperationLog) (err error) {
	err = mysql.DeleteOperation(&p)
	return
}
