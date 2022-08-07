/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:56
 * @FilePath: /pinkmoe_server/model/response/xd_user_authority.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package response

import uuid "github.com/satori/go.uuid"

type XdUseAuthority struct {
	XdUserId               uuid.UUID `gorm:"column:xd_user_uuid"`
	XdAuthorityAuthorityId string    `gorm:"column:xd_authority_authority_id"`
}

func (s *XdUseAuthority) TableName() string {
	return "xd_user_authority"
}
