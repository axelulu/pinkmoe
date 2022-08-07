/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:52:26
 * @FilePath: /pinkmoe_server/task/init.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package task

import (
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func Init() {
	c := cron.New()

	_, err := c.AddJob("* */10 * * *", checkUserOnline{})
	_, err = c.AddJob("* * */10 * *", updateUserVip{})
	if err != nil {
		zap.L().Info("开启定时失败")
		return
	}
	if err != nil {
		zap.L().Info("开启定时失败")
		return
	}

	c.Start()

	zap.L().Info("开启定时成功")

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
