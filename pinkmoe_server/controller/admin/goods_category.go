/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:49
 * @FilePath: /pinkmoe_server/controller/admin/goodsCategory.go
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
	controller.Register(&GoodsCategory{})
}

type GoodsCategory struct{}

func (goodsCategory *GoodsCategory) AuthGoodsCategoryListGet(c *gin.Context) {
	var p request.SearchGoodsCategoryParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchGoodsCategoryParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, list, total := adminLogic.GetGoodsCategoryList(p); err != nil {
		zap.L().Error("adminLogic.GetGoodsCategoryList err", zap.Error(err))
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

func (goodsCategory *GoodsCategory) AuthGoodsCategoryCreate(c *gin.Context) {
	var p model.XdGoodsCategory
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdGoodsCategory with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.CreateGoodsCategory(p); err != nil {
		zap.L().Error("adminLogic.CreateGoodsCategory err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (goodsCategory *GoodsCategory) AuthGoodsCategoryUpdate(c *gin.Context) {
	var p model.XdGoodsCategory
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdGoodsCategory with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.UpdateGoodsCategory(p); err != nil {
		zap.L().Error("adminLogic.UpdateGoodsCategory err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (goodsCategory *GoodsCategory) AuthGoodsCategoryDelete(c *gin.Context) {
	var p model.XdGoodsCategory
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdGoodsCategory with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.DeleteGoodsCategory(p); err != nil {
		zap.L().Error("adminLogic.DeleteGoodsCategory err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}
