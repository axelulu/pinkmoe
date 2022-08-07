/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:19
 * @FilePath: /pinkmoe_server/controller/cms/follow.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package cms

import (
	"server/controller"
	adminLogic "server/logic/admin"
	"server/model/request"
	"server/model/response"
	"server/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	controller.Register(&Follow{})
}

type Follow struct{}

func (follow *Follow) FollowListGet(c *gin.Context) {
	var p request.SearchFollowParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchFollowParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, list, total := adminLogic.GetFollowList(p); err != nil {
		zap.L().Error("adminLogic.GetFollowList err", zap.Error(err))
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

func (follow *Follow) AuthFollowStatusGet(c *gin.Context) {
	var p request.ReqFollow
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqFollow with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err, isFollow := adminLogic.GetFollowStatus(p, userId); err != nil {
		zap.L().Error("adminLogic.GetFollowStatus err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(isFollow, c)
	}
}

func (follow *Follow) AuthCreateFollow(c *gin.Context) {
	var p request.ReqFollow
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqFollow with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err := adminLogic.CreateFollow(p, userId); err != nil {
		zap.L().Error("adminLogic.CreateFollow err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (follow *Follow) AuthDeleteFollow(c *gin.Context) {
	var p request.ReqFollow
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqFollow with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err := adminLogic.DeleteFollow(p, userId); err != nil {
		zap.L().Error("adminLogic.DeleteFollow err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}
