/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:49
 * @FilePath: /pinkmoe_server/logic/admin/topic.go
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

func GetTopicList(p request.SearchTopicParams) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetTopicList(p)
	return
}

func CreateTopic(p model.XdTopic) (err error) {
	err = mysql.CreateTopic(p)
	return
}

func UpdateTopic(p model.XdTopic) (err error) {
	err = mysql.UpdateTopic(p)
	return
}

func DeleteTopic(p model.XdTopic) (err error) {
	err = mysql.DeleteTopic(p)
	return
}
