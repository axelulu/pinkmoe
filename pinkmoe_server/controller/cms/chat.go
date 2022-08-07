/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 10:47:19
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:17
 * @FilePath: /pinkmoe_server/controller/cms/chat.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package cms

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
	controller.Register(&Chat{})
}

type Chat struct{}

func (chat *Chat) AuthChatListGet(c *gin.Context) {
	var p request.SearchChatParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchChatParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err, list, total := adminLogic.GetChatList(p, userId); err != nil {
		zap.L().Error("adminLogic.GetChatList err", zap.Error(err))
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

func (chat *Chat) AuthChatRelationListGet(c *gin.Context) {
	var p request.SearchChatRelationParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchChatRelationParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err, list, total := adminLogic.GetChatRelationList(p, userId); err != nil {
		zap.L().Error("adminLogic.GetChatRelationList err", zap.Error(err))
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

func (chat *Chat) AuthCreateChat(c *gin.Context) {
	var p model.XdChat
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdChat with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err := adminLogic.CreateChat(p, userId); err != nil {
		zap.L().Error("adminLogic.CreateChat err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (chat *Chat) AuthCreateChatRelation(c *gin.Context) {
	var p model.XdChatRelation
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdChatRelation with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err := adminLogic.CreateChatRelation(p, userId); err != nil {
		zap.L().Error("adminLogic.CreateChatRelation err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (chat *Chat) AuthDeleteChatRelation(c *gin.Context) {
	var p model.XdChatRelation
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdChatRelation with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err := adminLogic.DeleteChatRelation(p, userId); err != nil {
		zap.L().Error("adminLogic.DeleteChatRelation err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}
