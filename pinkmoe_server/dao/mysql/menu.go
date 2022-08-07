/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:05
 * @FilePath: /pinkmoe_server/dao/mysql/menu.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/response"
	"strconv"

	"gorm.io/gorm"
)

func getMenuTreeMap(authorityId string) (err error, treeMap map[string][]model.XdMenu) {
	var allMenus []model.XdMenu
	treeMap = make(map[string][]model.XdMenu)
	err = global.XD_DB.Where("authority_id = ?", authorityId).Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[strconv.Itoa(int(v.ParentId))] = append(treeMap[strconv.Itoa(int(v.ParentId))], v)
	}
	return err, treeMap
}

func GetMenuTree(authorityId string) (err error, menus []model.XdMenu) {
	err, menuTree := getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	if err != nil {
		return response.ErrorMenuGet, nil
	}
	return err, menus
}

func getChildrenList(menu *model.XdMenu, treeMap map[string][]model.XdMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func GetInfoList() (err error, list interface{}, total int64) {
	var menuList []model.XdBaseMenu
	err, treeMap := getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	if err != nil {
		return response.ErrorMenuGet, nil, 0
	}
	return err, menuList, total
}

func getBaseMenuTreeMap() (err error, treeMap map[string][]model.XdBaseMenu) {
	var allMenus []model.XdBaseMenu
	treeMap = make(map[string][]model.XdBaseMenu)
	err = global.XD_DB.Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[strconv.Itoa(int(v.ParentId))] = append(treeMap[strconv.Itoa(int(v.ParentId))], v)
	}
	if err != nil {
		return response.ErrorMenuGet, nil
	}
	return err, treeMap
}

func getBaseChildrenList(menu *model.XdBaseMenu, treeMap map[string][]model.XdBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	if err != nil {
		return response.ErrorMenuGet
	}
	return err
}

func CreateMenu(menu *model.XdBaseMenu) (err error) {
	if !errors.Is(global.XD_DB.Where("name = ?", menu.Name).First(&model.XdBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return response.ErrorMenuName
	}
	if err = global.XD_DB.Create(&menu).Error; err != nil {
		return response.ErrorMenuCreate
	}
	return
}

func UpdateMenu(p *model.XdBaseMenu) (err error) {
	var oldMenu model.XdBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["parent_id"] = p.ParentId
	upDateMap["path"] = p.Path
	upDateMap["name"] = p.Name
	upDateMap["hidden"] = p.Hidden
	if p.Component == "LAYOUT" && p.ParentId != 0 {
		upDateMap["component"] = "PRENT_LAYOUT"
	} else if p.Component == "PRENT_LAYOUT" && p.ParentId == 0 {
		upDateMap["component"] = "LAYOUT"
	} else {
		upDateMap["component"] = p.Component
	}
	upDateMap["title"] = p.Title
	upDateMap["icon"] = p.Icon
	upDateMap["sort"] = p.Sort

	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", p.ID).Find(&oldMenu)
		if oldMenu.Name != p.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", p.ID, p.Name).First(&model.XdBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				return response.ErrorMenuName
			}
		}
		txErr := db.Updates(upDateMap).Error
		if txErr != nil {
			return response.ErrorMenuUpdate
		}
		return nil
	}); err != nil {
		return response.ErrorMenuUpdate
	}
	return err
}

func DeleteMenu(id int) (err error) {
	if err = global.XD_DB.Where("parent_id = ?", id).First(&model.XdBaseMenu{}).Error; err != nil {
		var menu model.XdBaseMenu
		db := global.XD_DB.Preload("XdAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		if len(menu.XdAuthoritys) > 0 {
			if err = global.XD_DB.Model(&menu).Association("XdAuthoritys").Delete(&menu.XdAuthoritys); err != nil {
				return response.ErrorMenuDel
			}
		} else {
			if err = db.Error; err != nil {
				return response.ErrorMenuDel
			}
		}
	} else {
		return response.ErrorMenuDelChild
	}
	return err
}
