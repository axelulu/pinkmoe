/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:08
 * @FilePath: /pinkmoe_server/controller/admin/user.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package admin

import (
	"encoding/json"
	"server/controller"
	"server/dao/mysql"
	"server/global"
	adminLogic "server/logic/admin"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/util"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

func init() {
	controller.Register(&User{})
}

type User struct{}

func (user *User) EmailCaptcha(c *gin.Context) {
	var p request.EmailCaptcha
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.EmailCaptcha with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	_, email := mysql.GetSettingItem(model.XdSetting{Slug: "system_email"})
	var emailConfig response.EmailConfig
	err := json.Unmarshal([]byte(email.Value), &emailConfig)
	if err != nil {
		// 记录日志
		zap.L().Error("request.EmailCaptcha with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := util.SendEmailCaptcha(p.EmailCode, p.EmailType, emailConfig); err != nil {
		zap.L().Error("util.SendEmailCaptcha err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (user *User) CaptchaGet(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.XD_CONFIG.BasicConfig.Captcha.ImgHeight, global.XD_CONFIG.BasicConfig.Captcha.ImgWidth, global.XD_CONFIG.BasicConfig.Captcha.KeyLong, 0.9, 80)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	if id, b64s, err := captcha.Generate(); err != nil {
		zap.L().Error("base64Captcha.NewCaptcha err", zap.Error(err))
		response.FailWithMessage(response.CodeValidateCode.Msg(), c)
		return
	} else {
		response.OkWithData(request.Captcha{
			CaptchaId:  id,
			CaptchaImg: b64s,
		}, c)
	}
}

func (user *User) Login(c *gin.Context) {
	var p request.Login
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.Login with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if base64Captcha.DefaultMemStore.Verify(p.CaptchaId, p.CaptchaCode, true) {
		if token, err := adminLogic.Login(p, c.GetHeader("User-Agent"), c.ClientIP()); err != nil {
			zap.L().Error("adminLogic.Login err", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			response.OkWithData(token, c)
		}
	} else {
		response.FailWithMessage(response.ErrorValidateCode.Error(), c)
	}
}

func (user *User) Reg(c *gin.Context) {
	var p request.Reg
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.Reg with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if user, err := adminLogic.Reg(p); err != nil {
		zap.L().Error("adminLogic.Reg err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(user, c)
	}
}

func (user *User) PwdForget(c *gin.Context) {}

func (user *User) AuthLogout(c *gin.Context) {
	userId, err := util.GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("ErrorUserNotExist err", zap.Error(err))
		response.FailWithMessage(response.ErrorUserNotExist.Error(), c)
		return
	}
	userToken := util.GetCurrentUserToken(c)
	if err := adminLogic.Logout(userId, userToken); err != nil {
		zap.L().Error("adminLogic.Logout err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (user *User) AuthUserGet(c *gin.Context) {
	if token, err := adminLogic.UserInfo(c); err != nil {
		zap.L().Error("adminLogic.UserInfo err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(token, c)
	}
}

func (user *User) AuthUserListGet(c *gin.Context) {
	var p request.SearchUserParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchUserParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, list, total := adminLogic.GetUserInfoList(p); err != nil {
		zap.L().Error("adminLogic.GetUserInfoList err", zap.Error(err))
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

func (user *User) AuthUserCreate(c *gin.Context) {
	var p model.XdUser
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("model.XdUser with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if user, err := adminLogic.CreateUser(p); err != nil {
		zap.L().Error("adminLogic.CreateUser err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(user, c)
	}
}

func (user *User) AuthUserUpdate(c *gin.Context) {
	var p request.ReqUser
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqUser with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.UpdateUser(p); err != nil {
		zap.L().Error("adminLogic.UpdateUser err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(user, c)
	}
}

func (user *User) AuthUserDelete(c *gin.Context) {
	var p request.GetByUUId
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.GetByUUId with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, err := util.GetCurrentUserID(c)
	if err != nil {
		// 记录日志
		zap.L().Error("ErrorUserNotExist err", zap.Error(err))
		response.FailWithMessage(response.ErrorUserNotExist.Error(), c)
		return
	}
	if err := adminLogic.DeleteUser(userId, p); err != nil {
		zap.L().Error("adminLogic.DeleteUser err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}
