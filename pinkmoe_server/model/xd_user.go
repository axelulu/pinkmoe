/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:59
 * @FilePath: /pinkmoe_server/model/xd_user.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package model

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
	Desc        string         `json:"desc" form:"desc" gorm:"comment:用户描述"`                                                    // 用户描述
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
