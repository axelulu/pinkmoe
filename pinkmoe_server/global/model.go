/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:27
 * @FilePath: /pinkmoe_server/global/model.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package global

type XD_MODEL struct {
	ID        uint    `gorm:"primarykey"` // 主键ID
	CreatedAt XdTime  // 创建时间
	UpdatedAt XdTime  // 更新时间
	DeletedAt *XdTime `gorm:"index" json:"-"` // 删除时间
}
