/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:47
 * @FilePath: /pinkmoe_server/controller/admin/authority.go
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
	controller.Register(&Authority{})
}

type Authority struct{}

func (authority *Authority) AuthAuthorityListGet(c *gin.Context) {
	var p request.PageInfo
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.PageInfo with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, list, total := adminLogic.GetAdminAuthorityList(p); err != nil {
		zap.L().Error("adminLogic.GetAdminAuthorityList err", zap.Error(err))
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

func (authority *Authority) AuthAuthorityMenuGet(c *gin.Context) {
	var p request.GetAuthorityId
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.GetAuthorityId with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, menus := adminLogic.GetMenuAuthority(p); err != nil {
		zap.L().Error("adminLogic.GetMenuAuthority err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(menus, c)
	}
}

func (authority *Authority) AuthAuthorityCreate(c *gin.Context) {
	var p model.XdAuthority
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdAuthority with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.CreateAuthority(p); err != nil {
		zap.L().Error("adminLogic.CreateAuthority err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (authority *Authority) AuthAuthorityMenuUpdate(c *gin.Context) {
	var p request.UpdateMenuAuthorityInfo
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.UpdateMenuAuthorityInfo with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.UpdateMenuAuthority(p); err != nil {
		zap.L().Error("adminLogic.UpdateMenuAuthority err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (authority *Authority) AuthAuthorityUpdate(c *gin.Context) {
	var p model.XdAuthority
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdAuthority with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.UpdateAuthority(p); err != nil {
		zap.L().Error("adminLogic.UpdateAuthority err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (authority *Authority) AuthAuthorityDelete(c *gin.Context) {
	var p model.XdAuthority
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdAuthority with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.DeleteAuthority(p); err != nil {
		zap.L().Error("adminLogic.DeleteAuthority err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}
