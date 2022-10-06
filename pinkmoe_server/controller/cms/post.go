/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:24
 * @FilePath: /pinkmoe_server/controller/cms/post.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package cms

import (
	"server/controller"
	adminLogic "server/logic/admin"
	"server/model/request"
	"server/model/response"
	"server/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	controller.Register(&Post{})
}

type Post struct{}

func (post *Post) PostListGet(c *gin.Context) {
	var p request.SearchPostParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchPostParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, list, total := adminLogic.GetPostList(p, p.UserId); err != nil {
		zap.L().Error("adminLogic.GetPostList err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(response.PageResult{
			List:      list,
			PageCount: util.GetPageCount(total, p.PageSize),
			Page:      p.Page,
			PageSize:  p.PageSize,
		}, c)
	}
}

func (post *Post) AuthUserPostListGet(c *gin.Context) {
	var p request.SearchPostParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchPostParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err, list, total := adminLogic.GetUserPostList(p, userId); err != nil {
		zap.L().Error("adminLogic.GetPostList err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(response.PageResult{
			List:      list,
			PageCount: util.GetPageCount(total, p.PageSize),
			Page:      p.Page,
			PageSize:  p.PageSize,
		}, c)
	}
}

func (post *Post) CategoryPostListGet(c *gin.Context) {
	var p request.SearchPostParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchPostParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, list, total := adminLogic.GetCategoryPostList(p); err != nil {
		zap.L().Error("adminLogic.GetCategoryPostList err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(response.PageResult{
			List:      list,
			PageCount: util.GetPageCount(total, p.PageSize),
			Page:      p.Page,
			PageSize:  p.PageSize,
		}, c)
	}
}

func (post *Post) SearchPostListGet(c *gin.Context) {
	var p request.SearchPostParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchPostParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, list, total := adminLogic.GetSearchPostList(p); err != nil {
		zap.L().Error("adminLogic.GetCategoryPostList err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(response.PageResult{
			List:      list,
			PageCount: util.GetPageCount(total, p.PageSize),
			Page:      p.Page,
			PageSize:  p.PageSize,
		}, c)
	}
}

func (post *Post) PostItemGet(c *gin.Context) {
	var p request.SearchPostParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.SearchPostParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err, list := adminLogic.GetPostByPostId(p, userId); err != nil {
		zap.L().Error("adminLogic.GetPostByPostId err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(list, c)
	}
}

func (post *Post) AuthPostItemDownloadGet(c *gin.Context) {
	var p request.ReqPost
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqPost with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err, list := adminLogic.GetPostDownloadByPostId(p, uuid); err != nil {
		zap.L().Error("adminLogic.GetPostDownloadByPostId err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(list, c)
	}
}

func (post *Post) AuthPostItemDownloadBuy(c *gin.Context) {
	var p request.ReqDownloadBuy
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqDownloadBuy with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err := adminLogic.BuyPostDownloadBy(p, uuid, c.ClientIP()); err != nil {
		zap.L().Error("adminLogic.BuyPostDownloadBy err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("购买成功", c)
	}
}

func (post *Post) AuthPostItemVideoGet(c *gin.Context) {
	var p request.ReqPost
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqPost with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err, list := adminLogic.GetPostVideoByPostId(p, uuid); err != nil {
		zap.L().Error("adminLogic.GetPostVideoByPostId err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(list, c)
	}
}

func (post *Post) AuthPostItemVideoBuy(c *gin.Context) {
	var p request.ReqVideoBuy
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqVideoBuy with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err := adminLogic.BuyPostVideoBy(p, uuid, c.ClientIP()); err != nil {
		zap.L().Error("adminLogic.BuyPostVideoBy err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("购买成功", c)
	}
}

func (post *Post) AuthPostItemMusicGet(c *gin.Context) {
	var p request.ReqPost
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqPost with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err, list := adminLogic.GetPostMusicByPostId(p, uuid); err != nil {
		zap.L().Error("adminLogic.GetPostMusicByPostId err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(list, c)
	}
}

func (post *Post) AuthPostItemMusicBuy(c *gin.Context) {
	var p request.ReqMusicBuy
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ReqMusicBuy with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	uuid, _ := util.GetCurrentUserID(c)
	if err := adminLogic.BuyPostMusicBy(p, uuid, c.ClientIP()); err != nil {
		zap.L().Error("adminLogic.BuyPostMusicBy err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("购买成功", c)
	}
}

func (post *Post) TopicPostGet(c *gin.Context) {
	var p request.SearchTopicParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志s
		zap.L().Error("request.SearchTopicParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err, list, total := adminLogic.GetTopicPost(p); err != nil {
		zap.L().Error("adminLogic.GetTopicPost err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(response.PageResult{
			List:      list,
			PageCount: util.GetPageCount(total, p.PageSize),
			Page:      p.Page,
			PageSize:  p.PageSize,
		}, c)
	}
}

func (post *Post) AuthCollectPostGet(c *gin.Context) {
	var p request.PageInfo
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志s
		zap.L().Error("request.SearchTopicParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err, list, total := adminLogic.GetCollectPost(p, userId); err != nil {
		zap.L().Error("adminLogic.GetCollectPost err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(response.PageResult{
			List:      list,
			PageCount: util.GetPageCount(total, p.PageSize),
			Page:      p.Page,
			PageSize:  p.PageSize,
		}, c)
	}
}

func (post *Post) PostViewUpdate(c *gin.Context) {
	var p request.CreatePostParams
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.CreatePostParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if err := adminLogic.PostViewUpdate(p); err != nil {
		zap.L().Error("adminLogic.PostViewUpdate err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (post *Post) AuthPostCollect(c *gin.Context) {
	var p request.CreatePostParams
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.CreatePostParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err := adminLogic.PostCollect(p, userId); err != nil {
		zap.L().Error("adminLogic.PostCollect err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (post *Post) AuthPostUnCollect(c *gin.Context) {
	var p request.CreatePostParams
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.CreatePostParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	userId, _ := util.GetCurrentUserID(c)
	if err := adminLogic.PostUnCollect(p, userId); err != nil {
		zap.L().Error("adminLogic.PostUnCollect err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

func (post *Post) AuthPostCreate(c *gin.Context) {
	var p request.CreatePostParams
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("request.CreatePostParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	if p.Type == "active" {
		p.Status = "published"
	} else {
		if p.Status != "pending" && p.Status != "draft" {
			p.Status = "pending"
		}
	}
	if err := adminLogic.CreatePost(p); err != nil {
		zap.L().Error("adminLogic.CreatePost err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}
