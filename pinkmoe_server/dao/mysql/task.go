/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:24
 * @FilePath: /pinkmoe_server/dao/mysql/task.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package mysql

import (
	"server/global"
	"server/model"
	"server/model/response"
)

func GetTaskList() (err error, task []model.XdTask) {
	db := global.XD_DB.Model(&model.XdTask{})
	if err = db.Find(&task).Error; err != nil {
		return response.ErrorTaskGet, task
	}
	return err, task
}

func GetTask(p *model.XdTask) (err error, task model.XdTask) {
	db := global.XD_DB.Model(&model.XdTask{})
	if p.Uuid.String() != "" {
		db = db.Where("uuid = ?", p.Uuid)
	}

	if p.Type != "" {
		db = db.Where("type = ?", p.Type)
	}

	if err = db.First(&task).Error; err != nil {
		return response.ErrorTaskGet, task
	}
	return err, task
}

func CreateTask(p *model.XdTask) (err error) {
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorTaskCreate
	}
	return err
}

func DeleteTask(p *model.XdTask) (err error) {
	if err = global.XD_DB.Delete(&p).Error; err != nil {
		return response.ErrorTaskCreate
	}
	return err
}
