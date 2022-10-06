/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-28 20:25:35
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:08
 * @FilePath: /pinkmoe_server/logic/admin/chat.go
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

	uuid "github.com/satori/go.uuid"
)

func GetChatList(p request.SearchChatParams, userId uuid.UUID) (err error, item interface{}, total int64) {
	p.UserId = userId.String()
	err, item, total = mysql.GetChatList(p)
	return
}

func GetChatRelationList(p request.SearchChatRelationParams, userId uuid.UUID) (err error, item interface{}, total int64) {
	p.UserId = userId.String()
	err, item, total = mysql.GetChatRelationList(p)
	return
}

func CreateChatRelation(p model.XdChatRelation, userId uuid.UUID) (err error) {
	p.UserId = userId
	err, m := mysql.GetChatRelationByUserId(p.UserId)
	if len(m) <= 0 {
		err = mysql.CreateChatRelation(&p)
		err = mysql.CreateChatRelation(&model.XdChatRelation{
			UserId: p.SendId,
			SendId: userId,
		})
	}
	return
}

func DeleteChatRelation(p model.XdChatRelation, userId uuid.UUID) (err error) {
	err = mysql.DeleteChatRelation(userId, p.SendId)
	err = mysql.DeleteChatRelation(p.SendId, userId)
	return
}

func CreateChat(p model.XdChat, userId uuid.UUID) (err error) {
	p.UserId = userId
	err, _ = mysql.CreateChat(&p)
	return
}
