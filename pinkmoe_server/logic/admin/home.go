/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:21
 * @FilePath: /pinkmoe_server/logic/admin/home.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package adminLogic

import (
	"encoding/json"
	"server/dao/mysql"
	"server/model"
	"server/model/response"
)

func GetHomeList() (err error, list response.Home) {
	var p model.XdSetting
	var homeConfig response.HomeConfig
	p.Slug = "website_home"
	err, home := mysql.GetSettingItem(p)
	err = json.Unmarshal([]byte(home.Value), &homeConfig)
	if err != nil {
		return err, list
	}
	err, list.Popular = mysql.GetPostByPostIds(homeConfig.Popular)
	err, list.Recommend = mysql.GetPostByPostIds(homeConfig.Recommend)
	for _, cm := range homeConfig.Cms {
		var categoryPosts response.Content
		_, categoryPost := mysql.GetCategoryById(cm.Category)
		categoryPosts.Name = categoryPost.Name
		categoryPosts.IconType = categoryPost.IconType
		categoryPosts.Icon = categoryPost.Icon
		categoryPosts.Slug = categoryPost.Slug
		categoryPosts.Sort = categoryPost.Sort
		categoryPosts.Style = cm.Style
		_, categoryPosts.Topic = mysql.GetTopicByTopicValues(cm.Topic)
		_, categoryPosts.Post = mysql.GetPostByCategoryId(cm.Category)
		list.Content = append(list.Content, categoryPosts)
	}
	return
}
