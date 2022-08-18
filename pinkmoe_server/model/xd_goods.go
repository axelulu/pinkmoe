/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:32
 * @FilePath: /pinkmoe_server/model/xd_post.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdGoods struct {
	CreatedAt        global.XdTime    // 创建时间
	UpdatedAt        global.XdTime    // 更新时间
	DeletedAt        *global.XdTime   `gorm:"index" json:"-"`                                                         // 删除时间
	GoodsId          string           `json:"goodsId" form:"goodsId" gorm:"not null;unique;primary_key;comment:商品ID"` // 商品ID
	Title            string           `json:"title" form:"title" gorm:"comment:商品标题"`                                 // 商品标题
	Content          string           `json:"content" form:"content" gorm:"type:text;comment:商品内容"`                   // 商品内容
	Category         string           `json:"category" form:"category" gorm:"comment:商品分类"`                           // 商品分类
	Author           uuid.UUID        `json:"author" form:"author" gorm:"comment:商品作者"`                               // 商品作者
	Cover            string           `json:"cover" form:"cover" gorm:"comment:商品封面"`                                 // 商品封面
	Type             string           `json:"type" form:"type" gorm:"comment:商品类型"`                                   // 商品类型
	View             int              `json:"view" form:"view" gorm:"comment:商品查看"`                                   // 商品查看
	CommentStatus    string           `json:"commentStatus" form:"commentStatus" gorm:"default:true;comment:评论状态"`    // 评论状态
	Status           string           `json:"status" form:"status" gorm:"comment:商品状态"`                               // 商品状态
	AuthorRelation   XdUser           `gorm:"foreignKey:Author;references:UUID;"`
	CategoryRelation XdGoodsCategory  `gorm:"foreignKey:Category;references:Slug"`
	SizeRelation     []XdGoodsSize    `json:"sizeRelation" gorm:"foreignKey:GoodsId;references:GoodsId;"`
	SizeValRelation  []XdGoodsSizeVal `json:"sizeValRelation" gorm:"foreignKey:GoodsId;references:GoodsId;"`
	GoodsImg         []XdUploadFile   `json:"goodsImg" gorm:"foreignKey:GoodsId;references:GoodsId;"`
}
