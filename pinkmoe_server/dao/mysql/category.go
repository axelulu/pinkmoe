/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:43
 * @FilePath: /pinkmoe_server/dao/mysql/category.go
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

func GetCategoryList(p request.SearchCategoryParams) (err error, list interface{}, total int64) {
	var menuList []model.XdCategory
	err, treeMap := getCategoryTreeMap(p)
	menuList = treeMap["0"]
	db := global.XD_DB.Model(&model.XdCategory{})
	for i := 0; i < len(menuList); i++ {
		err = getCategoryList(&menuList[i], treeMap)
	}
	if err != nil {
		return response.ErrorMenuGet, nil, 0
	}
	if err = db.Count(&total).Error; err != nil {
		return response.ErrorCommentListGet, nil, 0
	}
	return err, menuList, total
}

func GetCategoryById(categorySlug string) (err error, category model.XdCategory) {
	if err := global.XD_DB.Where("slug = ?", categorySlug).First(&category).Error; err != nil {
		return response.ErrorPostListGet, category
	}
	return err, category
}

func GetCategoryTreeById(categorySlug string) (err error, categoriesList []model.XdCategory) {
	var category model.XdCategory
	var categories []model.XdCategory
	if err := global.XD_DB.Where("slug = ?", categorySlug).First(&category).Error; err != nil {
		return response.ErrorPostListGet, categoriesList
	}
	categoriesList = append(categoriesList, category)
	err = getCategoryTree(category, &categories)
	categoriesList = append(categoriesList, categories...)
	util.Reverse(categoriesList)
	return
}

func getCategoryTree(category model.XdCategory, categories *[]model.XdCategory) (err error) {
	if category.ParentId != 0 {
		var categoryItem model.XdCategory
		if err := global.XD_DB.Where("id = ?", category.ParentId).First(&categoryItem).Error; err != nil {
			return response.ErrorPostListGet
		}
		*categories = append(*categories, categoryItem)
		err = getCategoryTree(categoryItem, categories)
	}
	return
}

func getCategoryTreeMap(p request.SearchCategoryParams) (err error, treeMap map[string][]model.XdCategory) {
	var allMenus []model.XdCategory
	treeMap = make(map[string][]model.XdCategory)
	db := global.XD_DB.Model(&model.XdCategory{})
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

func getAllChildCategoryBySlug(slug string) (err error, treeMap []model.XdCategory) {
	var allMenus []model.XdCategory
	db := global.XD_DB.Model(&model.XdCategory{})
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
	getAllChildCategory(allMenus, parentsId, &treeMap)
	return err, treeMap
}

func getAllChildCategory(allMenus []model.XdCategory, parentsId uint, treeMap *[]model.XdCategory) {
	for _, v := range allMenus {
		if v.ParentId == parentsId {
			*treeMap = append(*treeMap, v)
			for _, menu := range allMenus {
				if menu.ParentId == v.ID {
					getAllChildCategory(allMenus, v.ID, treeMap)
					break
				}
			}
		}
	}
}

func getCategoryList(menu *model.XdCategory, treeMap map[string][]model.XdCategory) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getCategoryList(&menu.Children[i], treeMap)
	}
	if err != nil {
		return response.ErrorMenuGet
	}
	return err
}

func CreateCategory(p model.XdCategory) (err error) {
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorCategoryCreate
	}
	return err
}

func UpdateCategory(p model.XdCategory) (err error) {
	var oldMenu model.XdCategory
	upDateMap := make(map[string]interface{})
	upDateMap["parent_id"] = p.ParentId
	upDateMap["name"] = p.Name
	upDateMap["slug"] = p.Slug
	upDateMap["icon"] = p.Icon
	upDateMap["sort"] = p.Sort

	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", p.ID).Find(&oldMenu)
		if oldMenu.Name != p.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", p.ID, p.Name).First(&model.XdCategory{}).Error, gorm.ErrRecordNotFound) {
				return response.ErrorMenuName
			}
		}
		txErr := db.Updates(upDateMap).Error
		if txErr != nil {
			return response.ErrorCategoryUpdate
		}
		return nil
	}); err != nil {
		return response.ErrorCategoryUpdate
	}
	return err
}

func DeleteCategory(p model.XdCategory) (err error) {
	if !errors.Is(global.XD_DB.Where("parent_id = ?", p.ID).First(&model.XdCategory{}).Error, gorm.ErrRecordNotFound) {
		return response.ErrorCategoryDelBychild
	}
	if !errors.Is(global.XD_DB.Where("category = ?", p.Slug).First(&model.XdPost{}).Error, gorm.ErrRecordNotFound) {
		return response.ErrorCategoryDelByPost
	}
	if err = global.XD_DB.Delete(&p).Error; err != nil {
		return response.ErrorCategoryDel
	}
	return
}
