/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 14:56:51
 * @FilePath: /xanaduCms/pinkmoe_server/initialize/initDB/user.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost) //加密处理
	entities := []model.XdUser{
		{UUID: uuid.NewV4(), Username: "admin", Email: "izhaicy@163.com", Password: string(password), NickName: "超级管理员", HeaderImg: "uploads/file/default/background.jpeg", Avatar: "uploads/file/default/avatar.png", AuthorityId: "9527"},
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
