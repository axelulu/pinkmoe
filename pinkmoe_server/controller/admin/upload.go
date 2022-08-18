/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:06
 * @FilePath: /pinkmoe_server/controller/admin/upload.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package admin

import (
	"server/controller"
	adminLogic "server/logic/admin"
	"server/model/request"
	"server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	controller.Register(&Upload{})
}

type Upload struct{}

func (upload *Upload) FileUpload(c *gin.Context) {
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	postId := c.Request.FormValue("post_id")
	uuid := c.Request.FormValue("uuid")
	uploadType := c.Request.FormValue("type")
	if err != nil {
		// 记录日志
		zap.L().Error("file with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, p := adminLogic.UploadFile(header, noSave, postId, "", uuid, uploadType); err != nil {
		zap.L().Error("adminLogic.UploadFile err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(response.ExaFileResponse{
			File: p,
		}, c)
	}
}

func (upload *Upload) AuthImgFileDelete(c *gin.Context) {
	var p request.DeleteImgFileParams
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.DeleteImgFileParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.DeleteImgFile(p); err != nil {
		zap.L().Error("adminLogic.DeleteImgFile err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}
