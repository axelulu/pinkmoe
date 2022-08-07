/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:05
 * @FilePath: /pinkmoe_server/initialize/initDB/post.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package initDB

import (
	"server/global"
	"server/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Post = new(post)

type post struct{}

func (u *post) TableName() string {
	return "xd_posts"
}

func (u *post) Initialize() error {
	var user model.XdUser
	global.XD_DB.Where("username = ?", "admin").First(&user)
	entities := []model.XdPost{
		{PostId: "2165371931828686848", Title: "测试文章", Content: "<p>测试文章</p>", Exerpt: "测试文章", Category: "news", Type: "post", Status: "published", Author: user.UUID},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (u *post) CheckDataExist() bool {
	if errors.Is(global.XD_DB.Where("post_id = ?", "2165371931828686848").First(&model.XdPost{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
