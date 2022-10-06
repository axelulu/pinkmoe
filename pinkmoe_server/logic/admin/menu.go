/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:31
 * @FilePath: /pinkmoe_server/logic/admin/menu.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package adminLogic

import (
	"server/dao/mysql"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/util"

	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) (userMenu []model.XdMenu, err error) {
	userId, err := util.GetCurrentUserID(c)
	if err != nil {
		return nil, response.ErrorUserNotLogin
	}
	var user model.XdUser
	if err, user = mysql.GetUserInfoByUserId(userId.String()); err != nil {
		return nil, err
	}
	if err, userMenu = mysql.GetMenuTree(user.AuthorityId); err != nil {
		return nil, err
	}
	return
}

func GetMenuList() (err error, list interface{}, total int64) {
	err, list, total = mysql.GetInfoList()
	return
}

func CreateMenu(p model.XdBaseMenu) (err error) {
	err = mysql.CreateMenu(&p)
	return
}

func UpdateMenu(p model.XdBaseMenu) (err error) {
	err = mysql.UpdateMenu(&p)
	return
}

func DeleteMenu(p request.GetById) (err error) {
	err = mysql.DeleteMenu(p.ID)
	return
}
