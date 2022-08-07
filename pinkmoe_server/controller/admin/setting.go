/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:03
 * @FilePath: /pinkmoe_server/controller/admin/setting.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package admin

import (
	"server/controller"
	adminLogic "server/logic/admin"
	"server/model"
	"server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	controller.Register(&Setting{})
}

type Setting struct{}

func (setting *Setting) AuthSettingItemGet(c *gin.Context) {
	var p model.XdSetting
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdSetting with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, item := adminLogic.GetSettingItem(p); err != nil {
		zap.L().Error("adminLogic.GetSettingItem err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(item, c)
	}
}

func (setting *Setting) AuthSettingCreate(c *gin.Context) {
	var p model.XdSetting
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdSetting with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.CreateSetting(p); err != nil {
		zap.L().Error("adminLogic.CreateSetting err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (setting *Setting) AuthSettingUpdate(c *gin.Context) {
	var p model.XdSetting
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdSetting with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.UpdateSetting(p); err != nil {
		zap.L().Error("adminLogic.UpdateSetting err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}
