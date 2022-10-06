/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:52:33
 * @FilePath: /pinkmoe_server/task/userOnline.go
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
