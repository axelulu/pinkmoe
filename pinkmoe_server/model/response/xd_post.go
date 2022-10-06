/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:44
 * @FilePath: /pinkmoe_server/model/response/xd_post.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package response

import (
	"server/global"
	"server/model"

	uuid "github.com/satori/go.uuid"
)

type XdPost struct {
	CreatedAt        global.XdTime          // 创建时间
	UpdatedAt        global.XdTime          // 更新时间
	DeletedAt        *global.XdTime         `gorm:"index" json:"-"`                                                       // 删除时间
	PostId           string                 `json:"postId" form:"postId" gorm:"not null;unique;primary_key;comment:文章ID"` // 用户UUID
	Title            string                 `json:"title" form:"title" gorm:"comment:文章标题"`                               // 文章标题
	Exerpt           string                 `json:"exerpt" form:"exerpt" gorm:"comment:文章摘要"`                             // 文章摘要
	Content          string                 `json:"content" form:"content" gorm:"type:text;comment:文章内容"`                 // 文章内容
	Category         string                 `json:"category" form:"category" gorm:"comment:文章分类"`                         // 文章分类
	Author           uuid.UUID              `json:"author" form:"author" gorm:"comment:文章作者"`                             // 文章作者
	Cover            string                 `json:"cover" form:"cover" gorm:"comment:文章封面"`                               // 文章封面
	Type             string                 `json:"type" form:"type" gorm:"comment:文章类型"`                                 // 文章类型
	View             int                    `json:"view" form:"view" gorm:"comment:文章查看"`                                 // 文章查看
	From             string                 `json:"from" form:"from" gorm:"comment:文章来源"`                                 // 文章来源
	CommentStatus    string                 `json:"commentStatus" form:"commentStatus" gorm:"comment:评论状态"`               // 评论状态
	Status           string                 `json:"status" form:"status" gorm:"comment:文章状态"`                             // 文章状态
	AuthorRelation   XdUser                 `gorm:"foreignKey:Author;references:UUID;"`
	CategoryRelation XdCategory             `gorm:"foreignKey:Category;references:Slug"`
	TopicRelations   []model.XdTopic        `json:"topic" gorm:"many2many:xd_topic_relation"`
	PostImg          []model.XdUploadFile   `json:"postImg" gorm:"foreignKey:PostId;references:PostId;"`
	DownloadRelation []model.XdPostDownload `json:"downloadRelation" gorm:"foreignKey:PostId;references:PostId;"`
	MusicRelation    []model.XdPostMusic    `json:"musicRelation" gorm:"foreignKey:PostId;references:PostId;"`
	VideoRelation    []model.XdPostVideo    `json:"videoRelation" gorm:"foreignKey:PostId;references:PostId;"`
	CommentRelation  []model.XdComment      `json:"commentRelation" gorm:"foreignKey:PostId;references:PostId;"`
	StarRelation     []model.XdPostStar     `json:"starRelation" gorm:"foreignKey:PostId;references:PostId;"`
	CollectRelation  []model.XdPostCollect  `json:"collectRelation" gorm:"foreignKey:PostId;references:PostId;"`
}
