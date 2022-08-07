/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:45
 * @FilePath: /pinkmoe_server/model/request/post.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

import (
	"server/global"
	"server/model"

	uuid "github.com/satori/go.uuid"
)

// SearchPostParams api分页条件查询及排序结构体
type SearchPostParams struct {
	PageInfo
	PostId        string `json:"postId" form:"postId"`               // 用户UUID
	Title         string `json:"title" form:"title"`                 // 文章标题
	Author        string `json:"author" form:"author"`               // 文章作者
	UserId        string `json:"userId" form:"userId"`               // 文章作者
	Type          string `json:"type" form:"type"`                   // 文章类型
	Category      string `json:"category" form:"category"`           // 文章分类
	Status        string `json:"status" form:"status"`               // 文章状态
	CommentStatus string `json:"commentStatus" form:"commentStatus"` // 评论状态
	Content       string `json:"content" form:"content"`             // 文章内容
	OrderKey      string `json:"orderKey" form:"orderKey"`           // 排序
	Desc          bool   `json:"desc" form:"desc"`                   // 排序方式:升序false(默认)|降序true
}

type ReqDownloadBuy struct {
	PostId     string `json:"postId" form:"postId"`         // 文章ID
	DownloadId int    `json:"downloadId" form:"downloadId"` // 下载ID
}

type ReqMusicBuy struct {
	PostId  string `json:"postId" form:"postId"`   // 文章ID
	MusicId int    `json:"musicId" form:"musicId"` // 音乐ID
}

type ReqVideoBuy struct {
	PostId  string `json:"postId" form:"postId"`   // 文章ID
	VideoId int    `json:"videoId" form:"videoId"` // 视频ID
}

type ReqPost struct {
	CreatedAt     global.XdTime  // 创建时间
	UpdatedAt     global.XdTime  // 更新时间
	DeletedAt     *global.XdTime `gorm:"index" json:"-"`                     // 删除时间
	PostId        string         `json:"postId" form:"postId"`               // 用户UUID
	Title         string         `json:"title" form:"title"`                 // 文章标题
	Exerpt        string         `json:"exerpt" form:"exerpt"`               // 文章摘要
	Content       string         `json:"content" form:"content"`             // 文章内容
	Category      string         `json:"category" form:"category"`           // 文章分类
	Author        string         `json:"author" form:"author"`               // 文章作者
	Cover         string         `json:"cover" form:"cover"`                 // 文章封面
	Type          string         `json:"type" form:"type"`                   // 文章类型
	View          int            `json:"view" form:"view"`                   // 文章查看
	Video         string         `json:"video" form:"video"`                 // 文章视频
	Music         string         `json:"music" form:"music"`                 // 文章音乐
	Download      string         `json:"download" form:"download"`           // 文章下载
	From          string         `json:"from" form:"from"`                   // 文章来源
	CommentStatus string         `json:"commentStatus" form:"commentStatus"` // 评论状态
	Status        string         `json:"status" form:"status"`               // 文章状态
}

// CreatePostParams api分页条件查询及排序结构体
type CreatePostParams struct {
	PostId           string                 `json:"postId" form:"postId" gorm:"comment:文章ID"`                              // 用户UUID
	Title            string                 `json:"title" form:"title" gorm:"comment:文章标题"`                                // 文章标题
	Exerpt           string                 `json:"exerpt" form:"exerpt" gorm:"comment:文章摘要"`                              // 文章摘要
	Content          string                 `json:"content" form:"content" gorm:"comment:文章内容"`                            // 文章内容
	Category         string                 `json:"category" form:"category" gorm:"comment:文章分类;default:null"`             // 文章分类
	Author           uuid.UUID              `json:"author" form:"author" gorm:"comment:文章作者"`                              // 文章作者
	Cover            string                 `json:"cover" form:"cover" gorm:"comment:文章封面"`                                // 文章封面
	Type             string                 `json:"type" form:"type" gorm:"comment:文章类型"`                                  // 文章类型
	From             string                 `json:"from" form:"from"`                                                      // 文章来源
	CommentStatus    string                 `json:"commentStatus" form:"commentStatus" gorm:"comment:评论状态;default:'open'"` // 评论状态
	Status           string                 `json:"status" form:"status" gorm:"comment:文章状态"`                              // 文章状态
	Topic            []string               `json:"topic" form:"topic"`
	PostImg          []model.XdUploadFile   `json:"postImg" form:"postImg"`
	DownloadRelation []model.XdPostDownload `json:"downloadRelation" form:"downloadRelation"`
	MusicRelation    []model.XdPostMusic    `json:"musicRelation" form:"musicRelation"`
	VideoRelation    []model.XdPostVideo    `json:"videoRelation" form:"videoRelation"`
}
