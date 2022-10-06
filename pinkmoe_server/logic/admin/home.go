/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:21
 * @FilePath: /pinkmoe_server/logic/admin/home.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package adminLogic

import (
	"encoding/json"
	"server/dao/mysql"
	"server/model"
	"server/model/request"
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
	_, list.Popular = mysql.GetPostByPostIds(homeConfig.Popular)
	_, list.Recommend = mysql.GetPostByPostIds(homeConfig.Recommend)
	for _, cm := range homeConfig.Cms {
		var categoryPosts response.Content
		_, categoryPost := mysql.GetCategoryById(cm.Category)
		categoryPosts.Name = categoryPost.Name
		categoryPosts.Icon = categoryPost.Icon
		categoryPosts.Slug = categoryPost.Slug
		categoryPosts.Sort = categoryPost.Sort
		categoryPosts.Style = cm.Style
		_, categoryPosts.Topic = mysql.GetTopicByTopicValues(cm.Topic)
		_, categoryPosts.Post, _ = mysql.GetCategoryPostList(request.SearchPostParams{
			PageInfo: request.PageInfo{
				Page:     1,
				PageSize: 12,
			},
			Category: cm.Category,
			OrderKey: "updated_at",
			Desc:     false,
		})
		list.Content = append(list.Content, categoryPosts)
	}
	return
}
