/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:08
 * @FilePath: /pinkmoe_server/initialize/initDB/setting.go
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

var Setting = new(setting)

type setting struct{}

func (a *setting) TableName() string {
	return "xd_settings"
}

func (a *setting) Initialize() error {
	entities := []model.XdSetting{
		{Name: "网站设置", Slug: "website_basic", Value: "{\"title\":\"PinkMoe\",\"keyword\":\"PinkMoe\",\"desc\":\"PinkMoe\",\"color\":\"rgba(255, 178, 178, 1)\",\"icon\":[\"uploads/file/default/avatar.png\"],\"background\":[\"uploads/file/default/background.jpeg\"],\"logo\":[\"uploads/file/default/avatar.png\"],\"createTime\":\"2007.06.30 12:08:55\"}"},
		{Name: "网站底部", Slug: "website_footer", Value: "{\"links\":[{\"name\":\"coderzhaolu\",\"url\":\"https://coderzhaolu.com\"}],\"notice\":\"欢迎欢迎\",\"contact\":\"欢迎欢迎\",\"about\":\"欢迎欢迎\",\"friendsLinks\":[{\"name\":\"coderzhaolu\",\"url\":\"https://coderzhaolu.com\"}]}"},
		{Name: "网站首页", Slug: "website_home", Value: "{\"popular\":[\"2165371931828686848\"],\"recommend\":[\"2165371931828686848\"],\"cms\":[{\"topic\":[\"测试\"],\"style\":1,\"category\":\"news\"}]}"},
		{Name: "积分奖励", Slug: "user_reward", Value: "{\"publishPostType\":\"credit\",\"publishPostCredit\":10,\"commentType\":\"credit\",\"commentCredit\":1,\"regType\":\"credit\",\"regCredit\":100,\"updatePwdType\":\"credit\",\"updatePwdCredit\":-10,\"updateEmailType\":\"credit\",\"updateEmailCredit\":-20,\"updateAvatarType\":\"credit\",\"updateAvatarCredit\":-10,\"checkInType\":\"credit\",\"checkInCreditHead\":20,\"checkInCreditFoot\":100,\"forgetPwdType\":\"credit\",\"forgetPwdCredit\":-10}"},
		{Name: "用户等级", Slug: "user_level", Value: "[{\"levelName\":\"lv1\",\"color\":\"rgba(255, 174, 0, 1)\",\"headExp\":0,\"footExp\":100},{\"levelName\":\"lv2\",\"color\":\"rgba(255, 0, 0, 1)\",\"headExp\":100,\"footExp\":2000},{\"levelName\":\"lv3\",\"color\":\"rgba(179, 255, 0, 1)\",\"headExp\":2000,\"footExp\":10000},{\"levelName\":\"lv4\",\"color\":\"rgba(0, 255, 76, 1)\",\"headExp\":10000,\"footExp\":50000},{\"levelName\":\"lv5\",\"color\":\"rgba(0, 195, 255, 1)\",\"headExp\":50000,\"footExp\":200000},{\"levelName\":\"lv6\",\"color\":\"rgba(0, 0, 255, 1)\",\"headExp\":200000,\"footExp\":500000},{\"levelName\":\"lv7\",\"color\":\"rgba(255, 0, 195, 1)\",\"headExp\":500000,\"footExp\":2000000},{\"levelName\":\"lv8\",\"color\":\"rgba(255, 204, 0, 1)\",\"headExp\":2000000,\"footExp\":9999999999}]"},
		{Name: "用户商城", Slug: "user_shop", Value: "[{\"shopCredit\":10,\"shopDesc\":\"购买100积分\",\"shopName\":\"购买100积分\",\"shopType\":\"credit\",\"shopValue\":100,\"priceType\":\"cash\"},{\"shopCredit\":45,\"shopDesc\":\"购买500积分\",\"shopName\":\"购买500积分\",\"shopType\":\"credit\",\"shopValue\":500,\"priceType\":\"cash\"},{\"shopCredit\":80,\"shopDesc\":\"购买1000积分\",\"shopName\":\"购买1000积分\",\"priceType\":\"cash\",\"shopValue\":1000,\"shopType\":\"credit\"},{\"shopCredit\":\"\",\"shopDesc\":\"充值100现金\",\"shopName\":\"充值100现金\",\"shopType\":\"cash\",\"shopValue\":100,\"priceType\":\"key\"}]"},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *setting) CheckDataExist() bool {
	if errors.Is(global.XD_DB.Where("slug = ?", "website_basic").First(&model.XdSetting{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
