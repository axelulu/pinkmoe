/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:16
 * @FilePath: /pinkmoe_server/dao/mysql/goods.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"encoding/json"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"

	"gorm.io/gorm"
)

func GetCategoryGoodsList(info request.SearchGoodsParams) (err error, list []model.XdGoods, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdGoods{})
	err, treeMap := getAllChildGoodsCategoryBySlug(info.Category)
	var allSlug []string
	for _, category := range treeMap {
		allSlug = append(allSlug, category.Slug)
	}
	fmt.Printf("%s", allSlug)

	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}

	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	var order string

	if len(info.OrderKey) == 0 {
		info.OrderKey = "updated_at"
	}

	if info.Desc {
		order = info.OrderKey + " DESC"
	} else {
		order = info.OrderKey + " ASC"
	}
	if info.Category == "0" || info.Category == "" {
		if err = db.Where("status = ?", "published").Preload("SizeRelation").Preload("SizeValRelation").Preload("AuthorRelation").Preload("CategoryRelation").Order(order).Limit(limit).Offset(offset).Find(&list).Error; err != nil {
			return response.ErrorGoodsListGet, list, 0
		}
	} else {
		if err = db.Where("status = ?", "published").Preload("SizeRelation").Preload("SizeValRelation").Preload("AuthorRelation").Preload("CategoryRelation").Order(order).Limit(limit).Offset(offset).Where("category IN ?", allSlug).Find(&list).Error; err != nil {
			return response.ErrorGoodsListGet, list, 0
		}
	}
	return
}

func GetAdminGoodsList(info request.SearchGoodsParams, userId string) (err error, goods []model.XdGoods, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdGoods{})

	var order string

	if len(info.OrderKey) <= 0 {
		info.OrderKey = "updated_at"
	}

	if info.Desc {
		order = info.OrderKey + " DESC"
	} else {
		order = info.OrderKey + " ASC"
	}

	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}

	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}

	if info.Category != "" {
		db = db.Where("category = ?", info.Category)
	}

	if info.Author != "" {
		db = db.Where("author = ?", info.Author)
	}

	if info.CommentStatus != "" {
		db = db.Where("comment_status = ?", info.CommentStatus)
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorGoodsListGet, nil, 0
	}
	if err = db.Limit(limit).Offset(offset).Order(order).Preload("GoodsImg").Preload("AuthorRelation").Preload("SizeRelation").Preload("SizeValRelation").Preload("CategoryRelation").Find(&goods).Error; err != nil {
		return response.ErrorGoodsListGet, nil, 0
	}
	return
}

func GetAuthorGoodsCount(uuid uuid.UUID) (err error, total int64) {
	if err = global.XD_DB.Model(model.XdGoods{}).Where("author = ?", uuid).Count(&total).Error; err != nil {
		return response.ErrorCommentCreate, 0
	}
	return
}

func GetGoodsList(info request.SearchGoodsParams, userId string) (err error, goods []model.XdGoods, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdGoods{})

	db = db.Where("status = 'published'")

	var order string
	if len(info.OrderKey) != 0 {
		if info.Desc {
			order = info.OrderKey + " DESC"
		} else {
			order = info.OrderKey + " ASC"
		}
		db.Order(order)
	}

	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}

	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	if info.Type != "" {
		if info.Type == "focus" {
			var follows []model.XdFollow
			global.XD_DB.Where("form_uid = ?", userId).Find(&follows)
			var fans []string
			for _, follow := range follows {
				fans = append(fans, follow.ToUid.String())
			}
			fans = append(fans, info.Author)
			db = db.Where("author IN (?)", fans)
		} else {
			db = db.Where("type = ?", info.Type)
		}
	}

	if info.Category != "" {
		db = db.Where("category = ?", info.Category)
	}

	if info.Author != "" && info.Type != "focus" {
		db = db.Where("author = ?", info.Author)
	}

	if info.CommentStatus != "" {
		db = db.Where("comment_status = ?", info.CommentStatus)
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorGoodsListGet, nil, 0
	}
	if err = db.Limit(limit).Offset(offset).Preload("AuthorRelation").Preload("AuthorRelation.Authority").Preload("CategoryRelation").Preload("GoodsImg").Find(&goods).Error; err != nil {
		return response.ErrorGoodsListGet, nil, 0
	}
	return
}

func GetUserGoodsList(info request.SearchGoodsParams, userId uuid.UUID) (err error, goods []model.XdGoods, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdGoods{})

	db = db.Where("author = ?", userId)

	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}

	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	if info.Type != "" {
		db = db.Where("type LIKE ?", "%"+info.Type+"%")
	}

	if info.Category != "" {
		db = db.Where("category = ?", info.Category)
	}

	if info.Status != "" {
		db = db.Where("status LIKE ?", "%"+info.Status+"%")
	}

	if info.CommentStatus != "" {
		db = db.Where("comment_status LIKE ?", "%"+info.CommentStatus+"%")
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorGoodsListGet, nil, 0
	}
	if err = db.Limit(limit).Offset(offset).Preload("AuthorRelation").Preload("CategoryRelation").Preload("GoodsImg").Preload("CommentRelation").Find(&goods).Error; err != nil {
		return response.ErrorGoodsListGet, nil, 0
	}
	return
}

func GetGoodsByGoodsIds(goodsIds []string) (err error, list []model.XdGoods) {
	if err := global.XD_DB.Where("status = ?", "published").Where("goods_id IN ?", goodsIds).Preload("AuthorRelation").Preload("CategoryRelation").Find(&list).Error; err != nil {
		return response.ErrorGoodsListGet, nil
	}
	return err, list
}

func GetGoodsByGoodsId(goodsIds string) (err error, goods model.XdGoods) {
	if err := global.XD_DB.Where("goods_id = ?", goodsIds).Preload("SizeRelation").Preload("SizeValRelation").Preload("CategoryRelation").Preload("AuthorRelation").Preload("AuthorRelation.Authority").Preload("GoodsImg").First(&goods).Error; err != nil {
		return response.ErrorGoodsListGet, goods
	}
	return err, goods
}

func CreateGoods(p request.CreateGoodsParams) (err error) {
	var goods model.XdGoods
	goods = model.XdGoods{
		GoodsId:       global.XD_SNOWFLAKE.Generate().String(),
		Title:         p.Title,
		Content:       p.Content,
		Category:      p.Category,
		Author:        p.Author,
		Cover:         p.Cover,
		Type:          p.Type,
		View:          p.View,
		CommentStatus: p.CommentStatus,
		Status:        p.Status,
	}
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&goods).Error; err != nil {
			return response.ErrorGoodsCreate
		}
		for _, img := range p.GoodsImg {
			if err = tx.Model(&model.XdUploadFile{}).Where("uuid = ?", img.Uuid).Where("url = ?", img.Url).Updates(map[string]interface{}{
				"type":     "goods",
				"goods_id": goods.GoodsId,
			}).Error; err != nil {
				return response.ErrorGoodsCreate
			}
		}
		for _, t := range p.SizeRelation {
			v, _ := json.Marshal(t.Value)
			newGoodsSize := model.XdGoodsSize{
				Key:     t.Key,
				Value:   string(v),
				GoodsId: goods.GoodsId,
			}
			var goodsSize model.XdGoodsSize
			if errors.Is(tx.Where("value = ?", v).Where("goods_id = ?", goods.GoodsId).First(&goodsSize).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
				err = tx.Create(&newGoodsSize).Error
			}
		}
		for _, t := range p.SizeValRelation {
			println(1)
			v, _ := json.Marshal(t.Value)
			newGoodsSizeVal := model.XdGoodsSizeVal{
				Stock:      t.Stock,
				Value:      string(v),
				SaleAmount: t.SaleAmount,
				GoodsId:    goods.GoodsId,
			}
			var goodsSizeVal model.XdGoodsSizeVal
			if errors.Is(tx.Where("value = ?", v).Where("goods_id = ?", goods.GoodsId).First(&goodsSizeVal).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
				err = tx.Create(&newGoodsSizeVal).Error
			}
		}
		return nil
	}); err != nil {
		return response.ErrorGoodsCreate
	}
	return err
}

func UpdateGoods(p request.CreateGoodsParams) (err error) {
	var goods model.XdGoods
	goods = model.XdGoods{
		GoodsId:       p.GoodsId,
		Title:         p.Title,
		Content:       p.Content,
		Category:      p.Category,
		Author:        p.Author,
		Cover:         p.Cover,
		Type:          p.Type,
		View:          p.View,
		CommentStatus: p.CommentStatus,
		Status:        p.Status,
	}
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Updates(&goods).Error; err != nil {
			return response.ErrorGoodsUpdate
		}
		for _, img := range p.GoodsImg {
			if err = tx.Model(&model.XdUploadFile{}).Where("uuid = ?", img.Uuid).Where("url = ?", img.Url).Updates(map[string]interface{}{
				"type":     "goods",
				"goods_id": goods.GoodsId,
			}).Error; err != nil {
				return response.ErrorGoodsUpdate
			}
		}
		tx.Where("goods_id = ?", goods.GoodsId).Delete(&model.XdGoodsSize{})
		for _, t := range p.SizeRelation {
			v, _ := json.Marshal(t.Value)
			newGoodsSize := model.XdGoodsSize{
				Key:     t.Key,
				Value:   string(v),
				GoodsId: goods.GoodsId,
			}
			var goodsSize model.XdGoodsSize
			if errors.Is(tx.Where("value = ?", v).Where("goods_id = ?", goods.GoodsId).First(&goodsSize).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
				err = tx.Create(&newGoodsSize).Error
			}
		}
		tx.Where("goods_id = ?", goods.GoodsId).Delete(&model.XdGoodsSizeVal{})
		for _, t := range p.SizeValRelation {
			v, _ := json.Marshal(t.Value)
			newGoodsSizeVal := model.XdGoodsSizeVal{
				Stock:      t.Stock,
				Value:      string(v),
				SaleAmount: t.SaleAmount,
				GoodsId:    goods.GoodsId,
			}
			var goodsSizeVal model.XdGoodsSizeVal
			if errors.Is(tx.Where("value = ?", v).Where("goods_id = ?", goods.GoodsId).First(&goodsSizeVal).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
				err = tx.Create(&newGoodsSizeVal).Error
			}
		}
		return nil
	}); err != nil {
		return response.ErrorGoodsUpdate
	}
	return err
}

func UpdateGoodsStatus(p request.CreateGoodsParams) (err error) {
	if err = global.XD_DB.Model(&model.XdGoods{}).Where("goods_id = ?", p.GoodsId).Update("status", p.Status).Error; err != nil {
		return response.ErrorGoodsUpdate
	}
	return err
}

func DeleteGoods(p model.XdGoods) (err error) {
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(&p).Error; err != nil {
			return response.ErrorGoodsDelete
		}
		return nil
	}); err != nil {
		return response.ErrorGoodsDelete
	}
	return
}
