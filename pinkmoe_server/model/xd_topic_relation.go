/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:44
 * @FilePath: /pinkmoe_server/model/xd_topic_relation.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

type XdTopicRelation struct {
	XdPostId  string `gorm:"column:xd_post_post_id"`
	XdTopicId uint   `gorm:"column:xd_topic_id"`
}

func (s *XdTopicRelation) TableName() string {
	return "xd_topic_relation"
}
