/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:04
 * @FilePath: /pinkmoe_server/model/xd_address.go
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

type XdAddress struct {
	CreatedAt global.XdTime  // 创建时间
	UpdatedAt global.XdTime  // 更新时间
	DeletedAt *global.XdTime `gorm:"index" json:"-"`                                            // 删除时间
	AddressId string         `json:"addressId" gorm:"not null;unique;primary_key;comment:订单ID"` // 订单ID
	Address   int64          `json:"address" gorm:"comment:地址信息"`                               // 地址信息
	Uuid      uuid.UUID      `json:"uuid" gorm:"comment:用户ID"`                                  // 用户ID
}
