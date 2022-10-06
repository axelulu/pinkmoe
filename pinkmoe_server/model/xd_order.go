/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:10
 * @FilePath: /pinkmoe_server/model/xd_order.go
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

type XdOrder struct {
	global.XD_MODEL
	OrderId         string        `json:"orderId" gorm:"comment:订单ID"`                                          // 订单ID
	OrderAddress    string        `json:"orderAddress" gorm:"comment:订单地址;default:null"`                        // 订单地址
	Price           int           `json:"price" gorm:"comment:订单价格"`                                            // 订单价格
	MoneyType       string        `json:"moneyType" gorm:"comment:货币类型"`                                        // 订单价格
	PaymentState    string        `json:"paymentState" gorm:"comment:支付状态"`                                     // 支付状态
	ShipmentState   string        `json:"shipmentState" gorm:"comment:运输状态"`                                    // 运输状态
	PaidAt          global.XdTime `json:"paidAt" gorm:"comment:支付时间"`                                           // 支付时间
	UserIp          string        `json:"userIp" gorm:"comment:用户ip"`                                           // 用户ip
	PostId          string        `json:"postId" gorm:"comment:文章ID;default:null"`                              // 文章ID
	Type            string        `json:"type" gorm:"comment:订单类型;default:'shop'"`                              // 订单类型
	Uuid            uuid.UUID     `json:"uuid" gorm:"comment:用户ID"`                                             // 用户ID
	GoodsId         uint          `json:"goodsId" gorm:"comment:商品ID"`                                          // 商品ID
	AddressRelation XdAddress     `json:"addressRelation" gorm:"foreignKey:OrderAddress;references:AddressId;"` // 订单地址
}
