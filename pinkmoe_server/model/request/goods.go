/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:45
 * @FilePath: /pinkmoe_server/model/request/post.go
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

	uuid "github.com/satori/go.uuid"
)

// SearchGoodsParams api分页条件查询及排序结构体
type SearchGoodsParams struct {
	PageInfo
	GoodsId       string `json:"goodsId" form:"goodsId" gorm:"not null;unique;primary_key;comment:商品ID"` // 商品ID
	Title         string `json:"title" form:"title" gorm:"comment:商品标题"`                                 // 商品标题
	Content       string `json:"content" form:"content" gorm:"type:text;comment:商品内容"`                   // 商品内容
	Category      string `json:"category" form:"category" gorm:"comment:商品分类"`                           // 商品分类
	Author        string `json:"author" form:"author" gorm:"comment:商品作者"`                               // 商品作者
	Cover         string `json:"cover" form:"cover" gorm:"comment:商品封面"`                                 // 商品封面
	Type          string `json:"type" form:"type" gorm:"comment:商品类型"`                                   // 商品类型
	View          int    `json:"view" form:"view" gorm:"comment:商品查看"`                                   // 商品查看
	CommentStatus string `json:"commentStatus" form:"commentStatus" gorm:"default:true;comment:评论状态"`    // 评论状态
	Status        string `json:"status" form:"status" gorm:"comment:商品状态"`                               // 商品状态
	OrderKey      string `json:"orderKey" form:"orderKey" gorm:"default:updated_at;"`                    // 排序
	Desc          bool   `json:"desc" form:"desc"`                                                       // 排序方式:升序false(默认)|降序true
}

type ReqGoods struct {
	CreatedAt     global.XdTime  // 创建时间
	UpdatedAt     global.XdTime  // 更新时间
	DeletedAt     *global.XdTime `gorm:"index" json:"-"`                                                         // 删除时间
	GoodsId       string         `json:"goodsId" form:"goodsId" gorm:"not null;unique;primary_key;comment:商品ID"` // 商品ID
	Title         string         `json:"title" form:"title" gorm:"comment:商品标题"`                                 // 商品标题
	Content       string         `json:"content" form:"content" gorm:"type:text;comment:商品内容"`                   // 商品内容
	Category      string         `json:"category" form:"category" gorm:"comment:商品分类"`                           // 商品分类
	Author        uuid.UUID      `json:"author" form:"author" gorm:"comment:商品作者"`                               // 商品作者
	Cover         string         `json:"cover" form:"cover" gorm:"comment:商品封面"`                                 // 商品封面
	Type          string         `json:"type" form:"type" gorm:"comment:商品类型"`                                   // 商品类型
	View          int            `json:"view" form:"view" gorm:"comment:商品查看"`                                   // 商品查看
	CommentStatus string         `json:"commentStatus" form:"commentStatus" gorm:"default:true;comment:评论状态"`    // 评论状态
	Status        string         `json:"status" form:"status" gorm:"comment:商品状态"`                               // 商品状态
}

type XdGoodsSize struct {
	CreatedAt global.XdTime  // 创建时间
	UpdatedAt global.XdTime  // 更新时间
	DeletedAt *global.XdTime `gorm:"index" json:"-"`                                                         // 删除时间
	Key       string         `json:"key" form:"key" gorm:"comment:规格名称"`                                     // 规格名称
	Value     []string       `json:"value" form:"value" gorm:"comment:规格配置"`                                 // 规格配置
	GoodsId   string         `json:"goodsId" form:"goodsId" gorm:"not null;unique;primary_key;comment:商品ID"` // 商品ID
}

type XdGoodsSizeVal struct {
	CreatedAt  global.XdTime  // 创建时间
	UpdatedAt  global.XdTime  // 更新时间
	DeletedAt  *global.XdTime `gorm:"index" json:"-"`                                                         // 删除时间
	Stock      int            `json:"stock" form:"stock" gorm:"comment:商品库存"`                                 // 商品库存
	Value      []string       `json:"value" form:"value" gorm:"comment:商品值"`                                  // 商品值
	SaleAmount int            `json:"saleAmount" form:"saleAmount" gorm:"comment:商品零售价"`                      // 商品零售价
	GoodsId    string         `json:"goodsId" form:"goodsId" gorm:"not null;unique;primary_key;comment:商品ID"` // 商品ID
}

// CreateGoodsParams api分页条件查询及排序结构体
type CreateGoodsParams struct {
	GoodsId         string               `json:"goodsId" form:"goodsId" gorm:"not null;unique;primary_key;comment:商品ID"` // 商品ID
	Title           string               `json:"title" form:"title" gorm:"comment:商品标题"`                                 // 商品标题
	Content         string               `json:"content" form:"content" gorm:"type:text;comment:商品内容"`                   // 商品内容
	Category        string               `json:"category" form:"category" gorm:"comment:商品分类"`                           // 商品分类
	Author          uuid.UUID            `json:"author" form:"author" gorm:"comment:商品作者"`                               // 商品作者
	Cover           string               `json:"cover" form:"cover" gorm:"comment:商品封面"`                                 // 商品封面
	Type            string               `json:"type" form:"type" gorm:"comment:商品类型"`                                   // 商品类型
	View            int                  `json:"view" form:"view" gorm:"comment:商品查看"`                                   // 商品查看
	CommentStatus   string               `json:"commentStatus" form:"commentStatus" gorm:"default:true;comment:评论状态"`    // 评论状态
	Status          string               `json:"status" form:"status" gorm:"comment:商品状态"`                               // 商品状态
	SizeRelation    []XdGoodsSize        `json:"sizeRelation" gorm:"foreignKey:GoodsId;references:GoodsId;"`
	SizeValRelation []XdGoodsSizeVal     `json:"sizeValRelation" gorm:"foreignKey:GoodsId;references:GoodsId;"`
	GoodsImg        []model.XdUploadFile `json:"goodsImg" form:"goodsImg"`
}
