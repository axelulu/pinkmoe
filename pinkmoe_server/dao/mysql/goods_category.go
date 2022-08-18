/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:43
 * @FilePath: /pinkmoe_server/dao/mysql/goodsCategory.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/util"
	"strconv"

	"gorm.io/gorm"
)

func GetGoodsCategoryList(p request.SearchGoodsCategoryParams) (err error, list []model.XdGoodsCategory, total int64) {
	var menuList []model.XdGoodsCategory
	err, treeMap := getGoodsCategoryTreeMap(p)
	db := global.XD_DB.Model(&model.XdGoodsCategory{})
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getGoodsCategoryList(&menuList[i], treeMap)
	}
	if err != nil {
		return response.ErrorMenuGet, nil, 0
	}
	if err = db.Count(&total).Error; err != nil {
		return response.ErrorCommentListGet, nil, 0
	}
	return err, menuList, total
}

func GetGoodsCategoryById(goodsCategorySlug string) (err error, goodsCategory model.XdGoodsCategory) {
	if err := global.XD_DB.Where("slug = ?", goodsCategorySlug).First(&goodsCategory).Error; err != nil {
		return response.ErrorPostListGet, goodsCategory
	}
	return err, goodsCategory
}

func GetGoodsCategoryTreeById(goodsCategorySlug string) (err error, categoriesList []model.XdGoodsCategory) {
	var goodsCategory model.XdGoodsCategory
	var categories []model.XdGoodsCategory
	if err := global.XD_DB.Where("slug = ?", goodsCategorySlug).First(&goodsCategory).Error; err != nil {
		return response.ErrorPostListGet, categoriesList
	}
	categoriesList = append(categoriesList, goodsCategory)
	err = getGoodsCategoryTree(goodsCategory, &categories)
	categoriesList = append(categoriesList, categories...)
	util.Reverse(categoriesList)
	return
}

func getGoodsCategoryTree(goodsCategory model.XdGoodsCategory, categories *[]model.XdGoodsCategory) (err error) {
	if goodsCategory.ParentId != 0 {
		var goodsCategoryItem model.XdGoodsCategory
		if err := global.XD_DB.Where("id = ?", goodsCategory.ParentId).First(&goodsCategoryItem).Error; err != nil {
			return response.ErrorPostListGet
		}
		*categories = append(*categories, goodsCategoryItem)
		err = getGoodsCategoryTree(goodsCategoryItem, categories)
	}
	return
}

func getGoodsCategoryTreeMap(p request.SearchGoodsCategoryParams) (err error, treeMap map[string][]model.XdGoodsCategory) {
	var allMenus []model.XdGoodsCategory
	treeMap = make(map[string][]model.XdGoodsCategory)
	db := global.XD_DB.Model(&model.XdGoodsCategory{})
	if p.Name != "" {
		db = db.Where("name LIKE ?", "%"+p.Name+"%")
	}

	if p.Slug != "" {
		db = db.Where("slug LIKE ?", "%"+p.Slug+"%")
	}

	if err = db.Order("sort").Find(&allMenus).Error; err != nil {
		return response.ErrorMenuGet, nil
	}
	for _, v := range allMenus {
		treeMap[strconv.Itoa(int(v.ParentId))] = append(treeMap[strconv.Itoa(int(v.ParentId))], v)
	}
	return err, treeMap
}

func getAllChildGoodsCategoryBySlug(slug string) (err error, treeMap []model.XdGoodsCategory) {
	var allMenus []model.XdGoodsCategory
	db := global.XD_DB.Model(&model.XdGoodsCategory{})
	if err = db.Order("sort").Find(&allMenus).Error; err != nil {
		return response.ErrorMenuGet, nil
	}
	var parentsId uint
	for _, v := range allMenus {
		if v.Slug == slug {
			treeMap = append(treeMap, v)
			parentsId = v.ID
			break
		}
	}
	for _, v := range allMenus {
		if v.ParentId == parentsId {
			treeMap = append(treeMap, v)
		}
	}
	return err, treeMap
}

func getGoodsCategoryList(menu *model.XdGoodsCategory, treeMap map[string][]model.XdGoodsCategory) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getGoodsCategoryList(&menu.Children[i], treeMap)
	}
	if err != nil {
		return response.ErrorMenuGet
	}
	return err
}

func CreateGoodsCategory(p model.XdGoodsCategory) (err error) {
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorGoodsCategoryCreate
	}
	return err
}

func UpdateGoodsCategory(p model.XdGoodsCategory) (err error) {
	var oldMenu model.XdGoodsCategory
	upDateMap := make(map[string]interface{})
	upDateMap["parent_id"] = p.ParentId
	upDateMap["name"] = p.Name
	upDateMap["slug"] = p.Slug
	upDateMap["icon"] = p.Icon
	upDateMap["icon_type"] = p.IconType
	upDateMap["sort"] = p.Sort

	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", p.ID).Find(&oldMenu)
		if oldMenu.Name != p.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", p.ID, p.Name).First(&model.XdGoodsCategory{}).Error, gorm.ErrRecordNotFound) {
				return response.ErrorMenuName
			}
		}
		txErr := db.Updates(upDateMap).Error
		if txErr != nil {
			return response.ErrorGoodsCategoryUpdate
		}
		return nil
	}); err != nil {
		return response.ErrorGoodsCategoryUpdate
	}
	return err
}

func DeleteGoodsCategory(p model.XdGoodsCategory) (err error) {
	if !errors.Is(global.XD_DB.Where("parent_id = ?", p.ID).First(&model.XdGoodsCategory{}).Error, gorm.ErrRecordNotFound) {
		return response.ErrorGoodsCategoryDelBychild
	}
	if !errors.Is(global.XD_DB.Where("category = ?", p.Slug).First(&model.XdPost{}).Error, gorm.ErrRecordNotFound) {
		return response.ErrorGoodsCategoryDelByPost
	}
	if err = global.XD_DB.Delete(&p).Error; err != nil {
		return response.ErrorGoodsCategoryDel
	}
	return
}
