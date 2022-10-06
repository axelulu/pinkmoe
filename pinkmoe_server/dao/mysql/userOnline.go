/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:35
 * @FilePath: /pinkmoe_server/dao/mysql/userOnline.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package mysql

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func GetUserOnlineList(info *request.SearchUserOnlineParams) (err error, list []model.XdUserOnline, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdUserOnline{})
	if info.UserName != "" {
		db = db.Where("user_name LIKE ?", "%"+info.UserName+"%")
	}
	if info.Ip != "" {
		db = db.Where("ip = ?", info.Ip)
	}

	if info.Explorer != "" {
		db = db.Where("explorer LIKE ?", "%"+info.Explorer+"%")
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorUserOnlineListGet, nil, 0
	}
	var comment []model.XdUserOnline
	if err = db.Limit(limit).Offset(offset).Find(&comment).Error; err != nil {
		return response.ErrorUserOnlineListGet, nil, 0
	}
	return err, comment, total
}

func CreateUserOnline(p *model.XdUserOnline) (err error) {
	var userOnline model.XdUserOnline
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if errors.Is(tx.Where("uuid = ?", p.Uuid).First(&userOnline).Error, gorm.ErrRecordNotFound) {
			if err = tx.Create(&p).Error; err != nil {
				return response.ErrorUserOnlineCreate
			}
		} else {
			if err = tx.Where("uuid = ?", p.Uuid).First(&model.XdUserOnline{}).Updates(&p).Error; err != nil {
				return response.ErrorUserOnlineCreate
			}
		}
		return nil
	}); err != nil {
		return response.ErrorUserOnlineCreate
	}
	return err
}

func UpdateUserOnline(u *model.XdUserOnline) (err error) {
	if err = global.XD_DB.Where("uuid = ?", u.Uuid).First(&model.XdUserOnline{}).Updates(&model.XdUserOnline{
		Uuid:     u.Uuid,
		Token:    u.Token,
		UserName: u.UserName,
		Ip:       u.Ip,
		Explorer: u.Explorer,
		Os:       u.Os,
	}).Error; err != nil {
		return response.ErrorUserOnlineUpdate
	}
	return err
}

func DeleteUserOnline(Uuid uuid.UUID) (err error) {
	if err = global.XD_DB.Where("uuid = ?", Uuid).Delete(&model.XdUserOnline{}).Error; err != nil {
		return response.ErrorUserOnlineDelete
	}
	return
}

func DeleteUserOnlineByToken(token string) (err error) {
	if err = global.XD_DB.Where("token = ?", token).Delete(&model.XdUserOnline{}).Error; err != nil {
		return response.ErrorUserOnlineDelete
	}
	return
}
