/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:37
 * @FilePath: /pinkmoe_server/dao/mysql/api.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package mysql

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"

	"gorm.io/gorm"
)

func GetAllApi() (err error, apis []model.XdApi) {
	if err = global.XD_DB.Find(&apis).Error; err != nil {
		return response.ErrorApiGet, nil
	}
	return
}

func GetAPIInfoList(p request.SearchApiParams) (err error, list interface{}, total int64) {
	limit := p.PageSize
	offset := p.PageSize * (p.Page - 1)
	db := global.XD_DB.Model(&model.XdApi{})
	var apiList []model.XdApi
	if p.Path != "" {
		db = db.Where("path LIKE ?", "%"+p.Path+"%")
	}

	if p.Method != "" {
		db = db.Where("method = ?", p.Method)
	}

	if p.ApiGroup != "" {
		db = db.Where("api_group LIKE ?", "%"+p.ApiGroup+"%")
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorApiCountGet, nil, 0
	} else {
		db = db.Limit(limit).Offset(offset)
		if p.OrderKey != "" {
			var OrderStr string
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[p.OrderKey] {
				if p.Desc {
					OrderStr = p.OrderKey + " desc"
				} else {
					OrderStr = p.OrderKey
				}
			}
			if err = db.Order(OrderStr).Find(&apiList).Error; err != nil {
				return response.ErrorApiGet, nil, 0
			}
		} else {
			if err = db.Order("api_group").Find(&apiList).Error; err != nil {
				return response.ErrorApiGet, nil, 0
			}
		}
	}
	return err, apiList, total
}

func GetApiAuthority(p request.GetAuthorityId) (pathMaps []request.CasbinInfo) {
	e := Casbin()
	list := e.GetFilteredPolicy(0, p.AuthorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

func CreateApi(p *model.XdApi) (err error) {
	if !errors.Is(global.XD_DB.Where("path = ? AND method = ?", p.Path, p.Method).First(&model.XdApi{}).Error, gorm.ErrRecordNotFound) {
		return response.ErrorApiBoth
	}
	return global.XD_DB.Create(&p).Error
}

func UpdateApi(api *model.XdApi) (err error) {
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		var oldA model.XdApi
		if err = tx.Where("id = ?", api.ID).First(&oldA).Error; err != nil {
			return response.ErrorApiUpdate
		} else {
			if oldA.Path != api.Path || oldA.Method != api.Method {
				if errors.Is(tx.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.XdApi{}).Error, gorm.ErrRecordNotFound) {
					return response.ErrorApiBoth
				}
			}
			if err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method); err != nil {
				return response.ErrorApiUpdate
			} else {
				if err = tx.Save(&api).Error; err != nil {
					return response.ErrorApiUpdate
				}
			}
		}
		return nil
	}); err != nil {
		return response.ErrorApiUpdate
	}
	return
}

func DeleteApi(api *model.XdApi) (err error) {
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if err = global.XD_DB.Delete(&api).Error; err != nil {
			return response.ErrorApiDel
		}
		if success := ClearCasbin(1, api.Path, api.Method); !success {
			return response.ErrorApiDel
		}
		return nil
	}); err != nil {
		return response.ErrorApiDel
	}
	return err
}
