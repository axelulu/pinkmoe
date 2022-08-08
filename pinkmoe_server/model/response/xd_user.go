/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:58
 * @FilePath: /pinkmoe_server/model/response/xd_user.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package response

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdUser struct {
	CreatedAt   global.XdTime  // 创建时间
	UpdatedAt   global.XdTime  // 更新时间
	DeletedAt   *global.XdTime `gorm:"index" json:"-"`                                                                          // 删除时间
	UUID        uuid.UUID      `json:"uuid" form:"uuid" gorm:"not null;unique;primary_key;comment:用户UUID"`                      // 用户UUID
	Username    string         `json:"userName" form:"userName" gorm:"comment:用户登录名"`                                           // 用户登录名
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                                // 用户登录密码
	NickName    string         `json:"nickName" form:"nickName" gorm:"default:系统用户;comment:用户昵称"`                               // 用户昵称
	Avatar      string         `json:"avatar" form:"avatar" gorm:"default:/uploads/file/default/avatar.png;comment:用户头像"`       // 用户头像
	HeaderImg   string         `json:"headerImg" form:"headerImg" gorm:"default:/uploads/file/default/avatar.png;comment:用户背景"` // 用户头像
	Sex         string         `json:"sex" form:"sex" gorm:"comment:用户性别;default:'1'"`                                          // 用户性别
	Cash        int            `json:"cash" form:"cash" gorm:"comment:用户现金;default:0"`                                          // 用户现金
	Credit      int            `json:"credit" form:"credit" gorm:"comment:用户积分;default:0"`                                      // 用户积分
	Exp         int            `json:"exp" form:"exp" gorm:"comment:用户经验;default:0"`                                            // 用户经验
	Phone       string         `json:"phone" form:"phone" gorm:"comment:用户手机号"`                                                 // 用户手机号
	Email       string         `json:"email" form:"email" gorm:"comment:用户邮箱"`                                                  // 用户角色ID
	AuthorityId string         `json:"authorityId" form:"authorityId" gorm:"default:2333;comment:用户角色ID;size:90"`               // 用户角色ID
	Authority   XdAuthority    `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []XdAuthority  `json:"authorities" gorm:"many2many:xd_user_authority;"`
}
