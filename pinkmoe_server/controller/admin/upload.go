/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:06
 * @FilePath: /pinkmoe_server/controller/admin/upload.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
