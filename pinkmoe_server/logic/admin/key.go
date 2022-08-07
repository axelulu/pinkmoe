/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:25
 * @FilePath: /pinkmoe_server/logic/admin/key.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package adminLogic

import (
	"server/dao/mysql"
	"server/model"
	"server/model/request"
	"server/util"
)

func GetKeyList(p request.SearchKeyParams) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetKeyList(p)
	return
}

func CreateKey(p request.SearchKeyParams) (err error) {
	for i := 0; i < p.Num; i++ {
		err = mysql.CreateKey(&model.XdKey{
			Key:    string(util.RandUp(128)),
			Value:  p.Value,
			Type:   p.Type,
			Status: "1",
		})
	}
	return
}

func DeleteKey(p model.XdKey) (err error) {
	err = mysql.DeleteKey(p)
	return
}
