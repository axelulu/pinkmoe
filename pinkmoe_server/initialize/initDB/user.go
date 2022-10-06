/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 17:45:41
 * @FilePath: /xanaduCms/pinkmoe_server/initialize/initDB/user.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package initDB

import (
	"server/global"
	"server/model"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var User = new(user)

type user struct{}

func (u *user) TableName() string {
	return "xd_users"
}

func (u *user) Initialize() error {
	password, _ := bcrypt.GenerateFromPassword([]byte("Admin123"), bcrypt.DefaultCost) //加密处理
	entities := []model.XdUser{
		{UUID: uuid.NewV4(), Username: "admin", Email: "izhaicy@163.com", Password: string(password), NickName: "超级管理员", HeaderImg: "/uploads/file/default/background.jpeg", Avatar: "/uploads/file/default/avatar.png", AuthorityId: "9527"},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (u *user) CheckDataExist() bool {
	if errors.Is(global.XD_DB.Where("username = ?", "admin").First(&model.XdUser{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
