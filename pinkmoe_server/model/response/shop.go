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

type ShopConfig struct {
	Popular   []string  `json:"popular"`
	Recommend []string  `json:"recommend"`
	Cms       []ShopCms `json:"cms"`
}

type ShopCms struct {
	Category string `json:"category"`
	Style    int    `json:"style"`
}

type ShopSettings struct {
	ParsingUrl string `json:"parsingUrl"`
}

type ShopContent struct {
	Slug     string          `json:"slug"`     // 菜单标识
	Name     string          `json:"name"`     // 菜单名称
	Icon     string          `json:"icon"`     // 菜单图标
	IconType string          `json:"iconType"` // 菜单图标类型
	Sort     int             `json:"sort"`     // 排序标记
	Style    int             `json:"style"`    // 分类模块显示类型
	Goods    []model.XdGoods `json:"goods"`    // 文章
}

type Shop struct {
	Popular      []model.XdGoods         `json:"popular"`
	ShopCategory []model.XdGoodsCategory `json:"shopCategory"`
	Content      []ShopContent           `json:"content"`
}
