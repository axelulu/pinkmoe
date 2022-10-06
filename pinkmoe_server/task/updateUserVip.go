/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:52:31
 * @FilePath: /pinkmoe_server/task/updateUserVip.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
