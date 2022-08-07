/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:23
 * @FilePath: /pinkmoe_server/model/response/home.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package response

import "server/model"

type HomeConfig struct {
	Popular   []string `json:"popular"`
	Recommend []string `json:"recommend"`
	Cms       []Cms    `json:"cms"`
}

type Cms struct {
	Category string   `json:"category"`
	Topic    []string `json:"topic"`
	Style    int      `json:"style"`
}

type Content struct {
	Slug     string          `json:"slug"`     // 菜单标识
	Name     string          `json:"name"`     // 菜单名称
	Icon     string          `json:"icon"`     // 菜单图标
	IconType string          `json:"iconType"` // 菜单图标类型
	Sort     int             `json:"sort"`     // 排序标记
	Style    int             `json:"style"`    // 分类模块显示类型
	Topic    []model.XdTopic `json:"topic"`    // 话题
	Post     []model.XdPost  `json:"post"`     // 文章
}

type Home struct {
	Popular   []model.XdPost `json:"popular"`
	Recommend []model.XdPost `json:"recommend"`
	Content   []Content      `json:"content"`
}
