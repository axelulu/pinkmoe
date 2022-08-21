/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:39
 * @FilePath: /pinkmoe_server/dao/mysql/authority.go
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

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func GetAuthorityList(info request.PageInfo, isVip bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdAuthority{})

	if isVip {
		db = db.Where("vip_start = 1")
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorAuthorityListGet, nil, 0
	}
	var authority []model.XdAuthority
	if err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Preload("AuthorityParams").Find(&authority).Error; err != nil {
		return response.ErrorAuthorityListGet, nil, 0
	}
	return err, authority, total
}

func GetAuthorityItem(authorityId string) (err error, authority model.XdAuthority) {
	db := global.XD_DB.Model(&model.XdAuthority{})
	if err = db.Where("authority_id = ?", authorityId).First(&authority).Error; err != nil {
		return response.ErrorAuthorityListGet, authority
	}
	return err, authority
}

func GetMenuAuthority(info request.GetAuthorityId) (err error, menus []model.XdMenu) {
	if err = global.XD_DB.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error; err != nil {
		return response.ErrorAuthorityMenuGet, nil
	}
	return err, menus
}

func GetAuthorityParams(uuid uuid.UUID, slug string) (err error, authorityParams model.XdAuthorityParams) {
	var user model.XdUser
	if err = global.XD_DB.Where("uuid = ? ", uuid).First(&user).Error; err != nil {
		return response.ErrorAuthorityParamsGet, authorityParams
	}
	if err = global.XD_DB.Where("authority_id = ? ", user.AuthorityId).Where("authority_param_id = ? ", slug).First(&authorityParams).Error; err != nil {
		return response.ErrorAuthorityParamsGet, authorityParams
	}
	return err, authorityParams
}

func CreateAuthority(auth model.XdAuthority) (err error, authority model.XdAuthority) {
	var authorityBox model.XdAuthority
	if !errors.Is(global.XD_DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return response.ErrorAuthoritySame, auth
	}
	if err = global.XD_DB.Create(&auth).Error; err != nil {
		return response.ErrorAuthorityCreate, auth
	}
	return err, auth
}

func UpdateMenuAuthority(auth model.XdAuthority) (err error) {
	var s model.XdAuthority
	global.XD_DB.Preload("XdBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	if err = global.XD_DB.Model(&s).Association("XdBaseMenus").Replace(&auth.XdBaseMenus); err != nil {
		return response.ErrorAuthorityMenuUpdate
	}
	return err
}

func UpdateAuthority(u model.XdAuthority, AuthorityId string) (err error) {
	if err = global.XD_DB.Where("authority_id = ?", AuthorityId).Select("*").Updates(&u).Error; err != nil {
		return response.ErrorAuthorityUpdate
	}
	return err
}

func CreateAuthorityParams(u *model.XdAuthorityParams) (err error) {
	if err = global.XD_DB.Create(&u).Error; err != nil {
		return response.ErrorAuthorityUpdate
	}
	return err
}

func DeleteAuthorityParams(AuthorityId string) (err error) {
	if err = global.XD_DB.Where("authority_id = ?", AuthorityId).Delete(&model.XdAuthorityParams{}).Error; err != nil {
		return response.ErrorAuthorityUpdate
	}
	return err
}

func DeleteAuthority(auth model.XdAuthority) (err error) {
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if !errors.Is(tx.Where("authority_id = ?", auth.AuthorityId).First(&model.XdUser{}).Error, gorm.ErrRecordNotFound) {
			return response.ErrorAuthorityDelByUser
		}
		db := tx.Preload("XdBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
		if err != nil {
			return response.ErrorAuthorityDel
		}
		if len(auth.XdBaseMenus) > 0 {
			err = tx.Model(auth).Association("XdBaseMenus").Delete(auth.XdBaseMenus)
			if err != nil {
				return response.ErrorAuthorityDel
			}
		} else {
			err = db.Error
			if err != nil {
				return response.ErrorAuthorityDel
			}
		}
		if err = tx.Delete(&[]model.XdUseAuthority{}, "xd_authority_authority_id = ?", auth.AuthorityId).Error; err != nil {
			return response.ErrorAuthorityDel
		}
		ClearCasbin(0, auth.AuthorityId)
		if err = db.Unscoped().Delete(auth).Error; err != nil {
			return response.ErrorCasbinUpdate
		}
		return nil
	}); err != nil {
		return response.ErrorAuthorityDel
	}
	return err
}

func DefaultMenu() []model.XdBaseMenu {
	return []model.XdBaseMenu{{
		XD_MODEL:  global.XD_MODEL{ID: 1},
		ParentId:  0,
		Path:      "/dashboard",
		Name:      "dashboard",
		Component: "LAYOUT",
		Sort:      0,
		Meta: model.Meta{
			Title: "Dashboard",
			Icon:  "DashboardOutlined",
		},
	}}
}

func DefaultCasbin() []request.CasbinInfo {
	return []request.CasbinInfo{
		{Path: "/api/v1/Admin/Menu/Menu", Method: "GET"},
		{Path: "/api/v1/Admin/User/User", Method: "POST"},
		{Path: "/api/v1/Admin/User/Logout", Method: "POST"},
		{Path: "/api/v1/Cms/User/Logout", Method: "POST"},
	}
}
