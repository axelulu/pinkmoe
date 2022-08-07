/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:25
 * @FilePath: /pinkmoe_server/controller/cms/setting.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package cms

import (
	"server/controller"
	adminLogic "server/logic/admin"
	"server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	controller.Register(&Setting{})
}

type Setting struct{}

func (setting *Setting) SiteSettingGet(c *gin.Context) {
	if err, item := adminLogic.GetSiteSetting(); err != nil {
		zap.L().Error("adminLogic.GetSiteSetting err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(item, c)
	}
}
