/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:28
 * @FilePath: /pinkmoe_server/controller/cms/user.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package cms

import (
	"server/controller"
	"server/global"
	adminLogic "server/logic/admin"
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
	if err := util.SendEmailCaptcha(p.EmailCode, p.EmailType); err != nil {
		zap.L().Error("util.SendEmailCaptcha err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (user *User) CaptchaGet(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.XD_CONFIG.BasicConfig.Captcha.ImgHeight, global.XD_CONFIG.BasicConfig.Captcha.ImgWidth, global.XD_CONFIG.BasicConfig.Captcha.KeyLong, 0.8, 60)
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

func (user *User) PwdForget(c *gin.Context) {
	var p request.PwdForget
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.PwdForget with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.PwdForget(p); err != nil {
		zap.L().Error("adminLogic.PwdForget err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

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

func (user *User) UserInfoGet(c *gin.Context) {
	var p request.SearchUserParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchUserParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, list := adminLogic.GetUserInfo(p); err != nil {
		zap.L().Error("adminLogic.GetPostList err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(list, c)
	}
}

func (user *User) AuthUserAvatarUpdate(c *gin.Context) {
	var p request.ReqUser
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqUser with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err := adminLogic.UpdateUserAvatar(p.Avatar, uuid); err != nil {
		zap.L().Error("adminLogic.UpdateUserAvatar err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(user, c)
	}
}

func (user *User) AuthUserBuyVip(c *gin.Context) {
	var p request.XdAuthorityParams
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.XdAuthorityParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err := adminLogic.UserBuyVip(p, uuid); err != nil {
		zap.L().Error("adminLogic.UserBuyVip err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(user, c)
	}
}

func (user *User) AuthUserBuyShop(c *gin.Context) {
	var p request.UserShop
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.UserShop with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err := adminLogic.UserBuyShop(p, uuid); err != nil {
		zap.L().Error("adminLogic.UserBuyShop err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(user, c)
	}
}

func (user *User) AuthUserCheckInStatus(c *gin.Context) {
	uuid, _ := util.GetCurrentUserID(c)
	if err, status := adminLogic.UserCheckInStatusGet(uuid); err != nil {
		zap.L().Error("adminLogic.UserCheckInStatusGet err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(status, c)
	}
}

func (user *User) AuthUserCheckIn(c *gin.Context) {
	uuid, _ := util.GetCurrentUserID(c)
	if err, credit := adminLogic.UpdateUserCheckIn(uuid, c.ClientIP()); err != nil {
		zap.L().Error("adminLogic.UpdateUserCheckIn err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(credit, c)
	}
}

func (user *User) AuthUserDetailUpdate(c *gin.Context) {
	var p request.ReqUser
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqUser with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err := adminLogic.UpdateUserDetail(p.NickName, p.Desc, uuid); err != nil {
		zap.L().Error("adminLogic.UpdateUserDetail err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(user, c)
	}
}

func (user *User) AuthUserEmailUpdate(c *gin.Context) {
	var p request.EmailUpdate
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.EmailUpdate with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err := adminLogic.UpdateUserEmail(p.EmailCode, p.EmailCaptcha, uuid); err != nil {
		zap.L().Error("adminLogic.UpdateUserEmail err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(user, c)
	}
}

func (user *User) AuthUserPwdUpdate(c *gin.Context) {
	var p request.PwdUpdate
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.PwdUpdate with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err := adminLogic.UpdateUserPwd(p.Password, p.OldPassword, uuid); err != nil {
		zap.L().Error("adminLogic.UpdateUserPwd err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(user, c)
	}
}
