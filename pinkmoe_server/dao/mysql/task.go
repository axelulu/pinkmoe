/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:24
 * @FilePath: /pinkmoe_server/dao/mysql/task.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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
