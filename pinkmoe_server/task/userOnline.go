/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:52:33
 * @FilePath: /pinkmoe_server/task/userOnline.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package task

import (
	"server/dao/mysql"
	"server/model"
	"server/model/request"
	jwt "server/util"

	"go.uber.org/zap"
)

type checkUserOnline struct{}

// Run 检查在线用户
func (t checkUserOnline) Run() {
	zap.L().Info("检查在线用户定时更新任务:")
	param := &request.SearchUserOnlineParams{
		PageInfo: request.PageInfo{
			PageSize: 50,
			Page:     1,
		},
	}
	var total int64
	for {
		var (
			list []model.XdUserOnline
			err  error
		)
		err, list, total = mysql.GetUserOnlineList(param)
		if err != nil {
			break
		}
		if list == nil {
			break
		}
		for _, v := range list {
			if _, err := jwt.ParseToken(v.Token); err != nil {
				_ = mysql.DeleteUserOnlineByToken(v.Token)
			}
		}
		if param.Page*param.PageSize >= int(total) {
			break
		}
		param.Page++
	}
}
