/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:53
 * @FilePath: /pinkmoe_server/dao/mysql/jwtBlackList.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/response"

	"gorm.io/gorm"
)

func HasJwtBlack(token string) (success bool) {
	if errors.Is(global.XD_DB.Where("jwt = ?", token).First(&model.XdJwtBlacklist{}).Error, gorm.ErrRecordNotFound) {
		return false
	} else {
		return true
	}
}

func CreateJwtBlack(token string) (err error) {
	if err = global.XD_DB.Create(&model.XdJwtBlacklist{
		Jwt: token,
	}).Error; err != nil {
		return response.ErrorJwtCreate
	}
	return
}
