/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:52:26
 * @FilePath: /pinkmoe_server/task/init.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
	_, err = c.AddJob("* * */1 * *", updateUserVip{})
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
