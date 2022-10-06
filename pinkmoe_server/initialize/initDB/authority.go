/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:50
 * @FilePath: /pinkmoe_server/initialize/initDB/authority.go
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

var Authority = new(authority)

type authority struct{}

func (a *authority) TableName() string {
	return "xd_authorities"
}

func (a *authority) Initialize() error {
	entities := []model.XdAuthority{
		{AuthorityId: "2333", AuthorityName: "普通用户", AuthorityWeight: 9999, AuthorityColor: "rgba(255, 0, 0, 1)", VipStart: 1},
		{AuthorityId: "9527", AuthorityName: "管理员", AuthorityWeight: 1, AuthorityColor: "rgba(255, 157, 0, 1)", VipStart: 0},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil {
		return errors.Wrapf(err, "%s表数据初始化失败!", a.TableName())
	}
	return nil
}

func (a *authority) CheckDataExist() bool {
	if errors.Is(global.XD_DB.Where("authority_id = ?", "9527").First(&model.XdAuthority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
