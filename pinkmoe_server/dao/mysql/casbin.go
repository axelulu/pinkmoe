/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:41
 * @FilePath: /pinkmoe_server/dao/mysql/casbin.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) (err error) {
	ClearCasbin(0, authorityId)
	var rules [][]string
	for _, v := range casbinInfos {
		cm := model.CasbinModel{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	e := Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return response.ErrorCasbinUpdate
	}
	return nil
}

func UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) (err error) {
	if err = global.XD_DB.Table("casbin_rule").Model(&model.CasbinModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error; err != nil {
		return response.ErrorCasbinUpdate
	}
	return
}

func GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

func ClearCasbin(v int, p ...string) (success bool) {
	e := Casbin()
	success, _ = e.RemoveFilteredPolicy(v, p...)
	return
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.XD_DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.XD_CONFIG.Casbin.ModelPath, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
