/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:56
 * @FilePath: /pinkmoe_server/controller/admin/loginLog.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package admin

import (
	"server/controller"
	adminLogic "server/logic/admin"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	controller.Register(&LoginLog{})
}

type LoginLog struct{}

func (operation *LoginLog) AuthLoginLogListGet(c *gin.Context) {
	var p request.SearchLoginLogParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchLoginLogParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, list, total := adminLogic.GetLoginLogList(p); err != nil {
		zap.L().Error("adminLogic.GetLoginLogList err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(response.PageResult{
			List:      list,
			PageCount: util.GetPageCount(total, p.PageSize),
			Page:      p.Page,
			PageSize:  p.PageSize,
		}, c)
	}
}

func (operation *LoginLog) AuthLoginLogCreate(c *gin.Context) {
	var p model.XdLoginLog
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdLoginLog with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.CreateLoginLog(p); err != nil {
		zap.L().Error("adminLogic.CreateLoginLog err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (operation *LoginLog) AuthLoginLogUpdate(c *gin.Context) {
	var p model.XdLoginLog
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdLoginLog with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.UpdateLoginLog(p); err != nil {
		zap.L().Error("adminLogic.UpdateLoginLog err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (operation *LoginLog) AuthLoginLogDelete(c *gin.Context) {
	var p model.XdLoginLog
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdLoginLog with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.DeleteLoginLog(p); err != nil {
		zap.L().Error("adminLogic.DeleteLoginLog err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}
