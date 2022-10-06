/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:45
 * @FilePath: /pinkmoe_server/logic/admin/setting.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package adminLogic

import (
	"server/dao/mysql"
	"server/model"
)

func GetSiteSetting() (err error, item map[string]interface{}) {
	item = make(map[string]interface{})
	var p model.XdSetting
	p.Slug = "website_basic"
	err, item["basic"] = mysql.GetSettingItem(p)
	p.Slug = "website_footer"
	err, item["footer"] = mysql.GetSettingItem(p)
	p.Slug = "user_reward"
	err, item["user_reward"] = mysql.GetSettingItem(p)
	p.Slug = "user_shop"
	err, item["user_shop"] = mysql.GetSettingItem(p)
	p.Slug = "user_level"
	err, item["user_level"] = mysql.GetSettingItem(p)
	p.Slug = "user_search"
	err, item["user_search"] = mysql.GetSettingItem(p)
	return
}

func GetSettingItem(p model.XdSetting) (err error, item interface{}) {
	err, item = mysql.GetSettingItem(p)
	return
}

func CreateSetting(p model.XdSetting) (err error) {
	err = mysql.CreateSetting(&p)
	return
}

func UpdateSetting(p model.XdSetting) (err error) {
	err = mysql.UpdateSetting(&p)
	return
}
