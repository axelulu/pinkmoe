/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:41
 * @FilePath: /pinkmoe_server/logic/admin/goods.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package adminLogic

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"server/dao/mysql"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
)

func GetAdminGoodsList(p request.SearchGoodsParams) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetAdminGoodsList(p, "")
	return
}

func GetGoodsByGoodsId(p request.SearchGoodsParams, userId uuid.UUID) (err error, list map[string]interface{}) {
	list = make(map[string]interface{})
	err, goods := mysql.GetGoodsByGoodsId(p.GoodsId)
	err, list["followCount"] = mysql.GetFollowCount(goods.Author)
	err, list["followStatus"] = mysql.GetFollowStatus(goods.Author, userId)
	err, list["fansCount"] = mysql.GetFansCount(goods.Author)
	err, list["goodsCount"] = mysql.GetAuthorGoodsCount(goods.Author)
	err, list["commentCount"] = mysql.GetAuthorCommentCount(goods.Author)
	list["goods"] = goods
	err, list["authorGoods"], _ = mysql.GetGoodsList(request.SearchGoodsParams{
		Author: goods.Author.String(),
		PageInfo: request.PageInfo{
			PageSize: 12,
			Page:     1,
		},
	}, "0")
	err, list["users"], _ = mysql.GetUserInfoList(request.SearchUserParams{
		PageInfo: request.PageInfo{
			Page:     1,
			PageSize: 9,
		},
		OrderKey: "credit",
		Desc:     false,
	})
	err, list["comments"], _ = mysql.GetCommentList(request.SearchCommentParams{
		PageInfo: request.PageInfo{
			Page:     1,
			PageSize: 5,
		},
		OrderKey: "",
		Desc:     false,
	})
	return
}

func GetGoodsList(p request.SearchGoodsParams, userId string) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetGoodsList(p, userId)
	return
}

func GetUserGoodsList(p request.SearchGoodsParams, userId uuid.UUID) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetUserGoodsList(p, userId)
	return
}

func GetCategoryGoodsList(p request.SearchGoodsParams) (err error, list map[string]interface{}, total int64) {
	list = make(map[string]interface{})
	err, list["goods"], total = mysql.GetCategoryGoodsList(p)
	err, list["category"] = mysql.GetGoodsCategoryTreeById(p.Category)
	return
}

func GetSearchGoodsList(p request.SearchGoodsParams) (err error, list []model.XdGoods, total int64) {
	err, list, total = mysql.GetCategoryGoodsList(p)
	return
}

func CreateGoods(p request.CreateGoodsParams) (err error) {
	err = mysql.CreateGoods(p)
	return
}

func GoodsViewUpdate(p request.CreateGoodsParams) (err error) {
	var goods model.XdGoods
	if err = global.XD_DB.Model(&model.XdGoods{}).Where("goods_id = ?", p.GoodsId).First(&goods).Error; err != nil {
		return response.ErrorGoodsUpdate
	}
	if err = global.XD_DB.Model(&model.XdGoods{}).Where("goods_id = ?", p.GoodsId).UpdateColumns(model.XdGoods{
		View: goods.View + 1,
		//UpdatedAt: found.UpdatedAt,
	}).Update("view", goods.View+1).Error; err != nil {
		return response.ErrorGoodsUpdate
	}
	return err
}

func UpdateGoods(p request.CreateGoodsParams) (err error) {
	err = mysql.UpdateGoods(p)
	return
}

func UpdateGoodsStatus(p request.CreateGoodsParams) (err error) {
	if err = mysql.UpdateGoodsStatus(p); err != nil {
		return err
	}
	// 添加10积分并通知
	if p.Status == "published" {
		// 初始化积分通知
		err, userReward := mysql.GetSettingItem(model.XdSetting{
			Slug: "user_reward",
		})
		var reward response.ResUserReward
		if err = json.Unmarshal([]byte(userReward.Value), &reward); err != nil {
			return response.ErrorGoodsUpdate
		}
		if reward.PublishGoodsType == "cash" {
			err = mysql.UpdateUserCash(reward.PublishGoodsCredit, p.Author)
			if err == nil {
				CreateNotification(p.Author, "", p.GoodsId, 0, reward.PublishGoodsCredit, 0, "publishGoods", "")
			}
		} else {
			err = mysql.UpdateUserCredit(reward.PublishGoodsCredit, p.Author)
			if err == nil {
				CreateNotification(p.Author, "", p.GoodsId, reward.PublishGoodsCredit, 0, 0, "publishGoods", "")
			}
		}
	}
	return
}

func DeleteGoods(p model.XdGoods) (err error) {
	err = mysql.DeleteGoods(p)
	return
}
