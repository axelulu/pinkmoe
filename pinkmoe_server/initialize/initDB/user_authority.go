/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:10
 * @FilePath: /pinkmoe_server/initialize/initDB/user_authority.go
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
	"gorm.io/gorm"
)

var UserAuthority = new(userAuthority)

type userAuthority struct{}

func (a *userAuthority) TableName() string {
	var entity model.XdUseAuthority
	return entity.TableName()
}

func (a *userAuthority) Initialize() error {
	var user model.XdUser
	global.XD_DB.Where("username = ?", "admin").First(&user)
	entities := []model.XdUseAuthority{
		{XdUserId: user.UUID, XdAuthorityAuthorityId: "9527"},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *userAuthority) CheckDataExist() bool {
	var user model.XdUser
	global.XD_DB.Where("username = ?", "admin").First(&user)
	if errors.Is(global.XD_DB.Where("xd_user_uuid = ? AND xd_authority_authority_id = ?", user.UUID, "9527").First(&model.XdUseAuthority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
