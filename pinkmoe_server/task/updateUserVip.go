/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:52:31
 * @FilePath: /pinkmoe_server/task/updateUserVip.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package task

import (
	"server/dao/mysql"
	"server/model"
	"time"

	"go.uber.org/zap"
)

type updateUserVip struct{}

// Run 检查在线用户
func (t updateUserVip) Run() {
	zap.L().Info("更新用户会员状态定时更新任务:")
	var (
		tasks []model.XdTask
		err   error
	)
	err, tasks = mysql.GetTaskList()
	if err != nil {
		return
	}
	if tasks == nil {
		return
	}
	for _, task := range tasks {
		if task.TimeOut.Unix() < time.Now().Unix() && task.Slug == "vip" {
			mysql.UpdateUserAuthority("2333", task.Uuid)
			mysql.DeleteTask(&task)
		}
	}
}
