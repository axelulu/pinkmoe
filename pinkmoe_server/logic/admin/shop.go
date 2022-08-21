/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:21
 * @FilePath: /pinkmoe_server/logic/admin/shop.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package adminLogic

import (
	"encoding/json"
	"server/dao/mysql"
	"server/model"
	"server/model/request"
	"server/model/response"
)

func GetShopList() (err error, list response.Shop) {
	var p model.XdSetting
	var shopConfig response.ShopConfig
	p.Slug = "website_shop"
	err, shop := mysql.GetSettingItem(p)
	err = json.Unmarshal([]byte(shop.Value), &shopConfig)
	if err != nil {
		return err, list
	}
	_, list.Popular = mysql.GetGoodsByGoodsIds(shopConfig.Popular)
	_, list.ShopCategory, _ = mysql.GetGoodsCategoryList(request.SearchGoodsCategoryParams{
		PageInfo: request.PageInfo{
			Page:     1,
			PageSize: 999,
		},
	})
	for _, cm := range shopConfig.Cms {
		var categoryPosts response.ShopContent
		_, categoryPost := mysql.GetGoodsCategoryById(cm.Category)
		categoryPosts.Name = categoryPost.Name
		categoryPosts.Icon = categoryPost.Icon
		categoryPosts.Slug = categoryPost.Slug
		categoryPosts.Sort = categoryPost.Sort
		categoryPosts.Style = cm.Style
		_, categoryPosts.Goods, _ = mysql.GetCategoryGoodsList(request.SearchGoodsParams{
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
