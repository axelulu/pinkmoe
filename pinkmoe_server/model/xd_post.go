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

type XdPost struct {
	CreatedAt        global.XdTime    `gorm:"autoCreateTime:false"`                                                 // 创建时间
	UpdatedAt        global.XdTime    `gorm:"autoCreateTime:false"`                                                 // 更新时间
	DeletedAt        *global.XdTime   `gorm:"index" json:"-"`                                                       // 删除时间
	PostId           string           `json:"postId" form:"postId" gorm:"not null;unique;primary_key;comment:文章ID"` // 用户UUID
	Title            string           `json:"title" form:"title" gorm:"comment:文章标题"`                               // 文章标题
	Exerpt           string           `json:"exerpt" form:"exerpt" gorm:"comment:文章摘要"`                             // 文章摘要
	Content          string           `json:"content" form:"content" gorm:"type:text;comment:文章内容"`                 // 文章内容
	Category         string           `json:"category" form:"category" gorm:"comment:文章分类;default:null"`            // 文章分类
	Author           uuid.UUID        `json:"author" form:"author" gorm:"comment:文章作者"`                             // 文章作者
	Cover            string           `json:"cover" form:"cover" gorm:"comment:文章封面"`                               // 文章封面
	Type             string           `json:"type" form:"type" gorm:"comment:文章类型"`                                 // 文章类型
	View             int              `json:"view" form:"view" gorm:"comment:文章查看"`                                 // 文章查看
	From             string           `json:"from" form:"from" gorm:"comment:文章来源"`                                 // 文章来源
	CommentStatus    string           `json:"commentStatus" form:"commentStatus" gorm:"default:true;comment:评论状态"`  // 评论状态
	Status           string           `json:"status" form:"status" gorm:"comment:文章状态"`                             // 文章状态
	AuthorRelation   XdUser           `gorm:"foreignKey:Author;references:UUID;"`
	CategoryRelation XdCategory       `gorm:"foreignKey:Category;references:Slug"`
	TopicRelations   []XdTopic        `json:"topic" gorm:"many2many:xd_topic_relation"`
	PostImg          []XdUploadFile   `json:"postImg" gorm:"foreignKey:PostId;references:PostId;"`
	DownloadRelation []XdPostDownload `json:"downloadRelation" gorm:"foreignKey:PostId;references:PostId;"`
	MusicRelation    []XdPostMusic    `json:"musicRelation" gorm:"foreignKey:PostId;references:PostId;"`
	VideoRelation    []XdPostVideo    `json:"videoRelation" gorm:"foreignKey:PostId;references:PostId;"`
	StarRelation     []XdPostStar     `json:"starRelation" gorm:"foreignKey:PostId;references:PostId;"`
	CollectRelation  []XdPostCollect  `json:"collectRelation" gorm:"foreignKey:PostId;references:PostId;"`
	FileRelation     []XdUploadFile   `json:"fileRelation" gorm:"foreignKey:PostId;references:PostId;"`
}
