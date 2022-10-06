/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:54
 * @FilePath: /pinkmoe_server/model/request/user.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package request

import (
	"server/global"
	"server/model"
)

// Login 登录参数
type Login struct {
	EmailCode   string `json:"emailCode" binding:"required"`
	CaptchaCode string `json:"captchaCode" binding:"required"`
	CaptchaId   string `json:"captchaId" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

// Reg 注册参数
type Reg struct {
	Username    string `json:"username" binding:"required"`
	EmailCode   string `json:"emailCode" binding:"required"`
	CaptchaCode string `json:"captchaCode" binding:"required"`
	CaptchaType string `json:"captchaType" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RePassword  string `json:"rePassword" binding:"required"`
}

// PwdForget 注册参数
type PwdForget struct {
	EmailCode   string `json:"emailCode" binding:"required"`
	CaptchaCode string `json:"captchaCode" binding:"required"`
	CaptchaType string `json:"captchaType" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RePassword  string `json:"rePassword" binding:"required"`
}

// ReqUser 创建新用户
type ReqUser struct {
	global.XD_MODEL
	UUID        string              `json:"uuid" form:"uuid" gorm:"comment:用户UUID"`                                                  // 用户UUID
	Username    string              `json:"userName" form:"userName" gorm:"comment:用户登录名"`                                           // 用户登录名
	Password    string              `json:"password" form:"password" gorm:"comment:用户登录密码"`                                          // 用户登录密码
	NickName    string              `json:"nickName" form:"nickName" gorm:"default:系统用户;comment:用户昵称"`                               // 用户昵称
	Desc        string              `json:"desc" form:"desc" gorm:"comment:用户描述"`                                                    // 用户描述
	Avatar      string              `json:"avatar" form:"avatar" gorm:"default:/uploads/file/default/avatar.png;comment:用户头像"`       // 用户头像
	HeaderImg   string              `json:"headerImg" form:"headerImg" gorm:"default:/uploads/file/default/avatar.png;comment:用户背景"` // 用户头像
	Sex         string              `json:"sex" form:"sex" gorm:"comment:用户性别default:0"`                                             // 用户性别
	Cash        int                 `json:"cash" form:"cash" gorm:"comment:用户现金default:0"`                                           // 用户现金
	Credit      int                 `json:"credit" form:"credit" gorm:"comment:用户积分default:0"`                                       // 用户积分
	Exp         int                 `json:"exp" form:"exp" gorm:"comment:用户经验default:0"`                                             // 用户经验
	Phone       string              `json:"phone" form:"phone" gorm:"comment:用户手机号"`                                                 // 用户手机号
	Email       string              `json:"email" form:"email" gorm:"comment:用户邮箱"`                                                  // 用户邮箱
	AuthorityId string              `json:"authorityId" form:"authorityId" gorm:"default:2333;comment:用户角色ID;size:90"`               // 用户角色ID
	Authority   model.XdAuthority   `json:"authority" gorm:"comment:用户角色"`
	Authorities []model.XdAuthority `json:"authorities" gorm:"many2many:xd_user_authority;"`
}

// Captcha 注册参数
type Captcha struct {
	CaptchaId  string `json:"captchaId" binding:"required"`
	CaptchaImg string `json:"captchaImg" binding:"required"`
}

// EmailCaptcha 注册参数
type EmailCaptcha struct {
	EmailCode string `json:"emailCode" form:"emailCode" binding:"required"`
	EmailType string `json:"emailType" form:"emailType" binding:"required"`
}

// EmailUpdate 注册参数
type EmailUpdate struct {
	EmailCaptcha string `json:"emailCaptcha" form:"emailCaptcha" binding:"required"`
	EmailCode    string `json:"emailCode" form:"emailCode" binding:"required"`
}

// PwdUpdate 注册参数
type PwdUpdate struct {
	Password    string `json:"password" form:"password" binding:"required"`
	OldPassword string `json:"oldPassword" form:"oldPassword" binding:"required"`
}

// UserShop 用户商城参数
type UserShop struct {
	ShopType   string `json:"shopType" form:"shopType" binding:"required"`
	PriceType  string `json:"priceType" form:"priceType" binding:"required"`
	ShopCredit int    `json:"shopCredit" form:"shopCredit"`
	ShopValue  int    `json:"shopValue" form:"shopValue" binding:"required"`
	ShopNum    int    `json:"shopNum" form:"shopNum" binding:"required"`
	ShopKey    string `json:"shopKey" form:"shopKey"`
}

// SearchUserParams api分页条件查询及排序结构体
type SearchUserParams struct {
	PageInfo
	UUID        string `json:"uuid" form:"uuid" gorm:"comment:用户UUID"`                                    // 用户UUID
	Username    string `json:"userName" form:"userName" gorm:"comment:用户登录名"`                             // 用户登录名
	NickName    string `json:"nickName" form:"nickName" gorm:"default:系统用户;comment:用户昵称"`                 // 用户昵称
	Sex         string `json:"sex" form:"sex" gorm:"comment:用户性别default:0"`                               // 用户性别
	Phone       string `json:"phone" form:"phone" gorm:"comment:用户手机号"`                                   // 用户手机号
	Email       string `json:"email" form:"email" gorm:"comment:用户邮箱"`                                    // 用户邮箱
	AuthorityId string `json:"authorityId" form:"authorityId" gorm:"default:2333;comment:用户角色ID;size:90"` // 用户角色ID
	OrderKey    string `json:"orderKey"`                                                                  // 排序
	Desc        bool   `json:"desc"`                                                                      // 排序方式:升序false(默认)|降序true
}
