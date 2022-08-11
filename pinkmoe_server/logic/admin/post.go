/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:41
 * @FilePath: /pinkmoe_server/logic/admin/post.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package adminLogic

import (
	"encoding/json"
	"server/dao/mysql"
	"server/model"
	"server/model/request"
	"server/model/response"

	uuid "github.com/satori/go.uuid"
)

func GetAdminPostList(p request.SearchPostParams, userId string) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetAdminPostList(p, userId)
	return
}

func GetPostList(p request.SearchPostParams, userId string) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetPostList(p, userId)
	return
}

func GetUserPostList(p request.SearchPostParams, userId uuid.UUID) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetUserPostList(p, userId)
	return
}

func GetCategoryPostList(p request.SearchPostParams) (err error, list map[string]interface{}, total int64) {
	list = make(map[string]interface{})
	err, list["post"], total = mysql.GetCategoryPostList(p)
	err, list["category"] = mysql.GetCategoryTreeById(p.Category)
	return
}

func GetSearchPostList(p request.SearchPostParams) (err error, list []model.XdPost, total int64) {
	err, list, total = mysql.GetCategoryPostList(p)
	return
}

func GetTopicPost(p request.SearchTopicParams) (err error, list model.XdTopic, total int64) {
	err, list, total = mysql.GetTopicPostList(p)
	return
}

func GetCollectPost(p request.PageInfo, userId uuid.UUID) (err error, list []model.XdPostCollect, total int64) {
	err, list, total = mysql.GetCollectPost(p, userId)
	return
}

func GetPostDownloadByPostId(p request.ReqPost, uuid uuid.UUID) (err error, list []response.XdPostDownload) {
	err, list = mysql.GetPostDownloadByPostId(p.PostId, uuid)
	return
}

func BuyPostDownloadBy(p request.ReqDownloadBuy, uuid uuid.UUID, ip string) (err error) {
	err, _, price, priceSlug := mysql.BuyPostDownload(p, uuid, ip)
	if err == nil {
		err, authorityParam := mysql.GetAuthorityParams(uuid, "downloadDiscounts")
		if err != nil {
			return err
		}
		if priceSlug == "cash" {
			CreateNotification(uuid, "", p.PostId, 0, -(int(price * authorityParam.AuthorityParamValue / 100)), 0, "buyDownload", "")
		} else {
			CreateNotification(uuid, "", p.PostId, -(int(price * authorityParam.AuthorityParamValue / 100)), 0, 0, "buyDownload", "")
		}
	}
	return
}

func GetPostVideoByPostId(p request.ReqPost, uuid uuid.UUID) (err error, list []response.XdPostVideo) {
	err, list = mysql.GetPostVideoByPostId(p.PostId, uuid)
	return
}

func BuyPostVideoBy(p request.ReqVideoBuy, uuid uuid.UUID, ip string) (err error) {
	err, _, price, priceSlug := mysql.BuyPostVideo(p, uuid, ip)
	if err == nil {
		err, authorityParam := mysql.GetAuthorityParams(uuid, "downloadDiscounts")
		if err != nil {
			return err
		}
		if priceSlug == "cash" {
			CreateNotification(uuid, "", p.PostId, 0, -(int(price * authorityParam.AuthorityParamValue / 100)), 0, "buyVideo", "")
		} else {
			CreateNotification(uuid, "", p.PostId, -(int(price * authorityParam.AuthorityParamValue / 100)), 0, 0, "buyVideo", "")
		}
	}
	return
}

func GetPostMusicByPostId(p request.ReqPost, uuid uuid.UUID) (err error, list []response.XdPostMusic) {
	err, list = mysql.GetPostMusicByPostId(p.PostId, uuid)
	return
}

func BuyPostMusicBy(p request.ReqMusicBuy, uuid uuid.UUID, ip string) (err error) {
	err, _, price, priceSlug := mysql.BuyPostMusic(p, uuid, ip)
	if err == nil {
		err, authorityParam := mysql.GetAuthorityParams(uuid, "downloadDiscounts")
		if err != nil {
			return err
		}
		if priceSlug == "cash" {
			CreateNotification(uuid, "", p.PostId, 0, -(int(price * authorityParam.AuthorityParamValue / 100)), 0, "buyMusic", "")
		} else {
			CreateNotification(uuid, "", p.PostId, -(int(price * authorityParam.AuthorityParamValue / 100)), 0, 0, "buyMusic", "")
		}
	}
	return
}

func GetPostByPostId(p request.SearchPostParams, userId uuid.UUID) (err error, list map[string]interface{}) {
	list = make(map[string]interface{})
	err, post := mysql.GetPostByPostId(p.PostId)
	err, list["followCount"] = mysql.GetFollowCount(post.Author)
	err, list["followStatus"] = mysql.GetFollowStatus(post.Author, userId)
	err, list["fansCount"] = mysql.GetFansCount(post.Author)
	err, list["postCount"] = mysql.GetAuthorPostCount(post.Author)
	err, list["commentCount"] = mysql.GetAuthorCommentCount(post.Author)
	list["post"] = post
	err, list["authorPosts"], _ = mysql.GetPostList(request.SearchPostParams{
		Author: post.Author.String(),
		PageInfo: request.PageInfo{
			PageSize: 12,
			Page:     1,
		},
	}, "0")
	err, list["users"], _ = mysql.GetUserInfoList(request.SearchUserParams{
		PageInfo: request.PageInfo{
			Page:     1,
			PageSize: 9,
		},
		OrderKey: "credit",
		Desc:     false,
	})
	err, list["comments"], _ = mysql.GetCommentList(request.SearchCommentParams{
		PageInfo: request.PageInfo{
			Page:     1,
			PageSize: 5,
		},
		OrderKey: "",
		Desc:     false,
	})
	return
}

func GetBbsSilder() (err error, list map[string]interface{}) {
	list = make(map[string]interface{})
	err, list["users"], _ = mysql.GetUserInfoList(request.SearchUserParams{
		PageInfo: request.PageInfo{
			Page:     1,
			PageSize: 9,
		},
		OrderKey: "credit",
		Desc:     false,
	})
	err, list["comments"], _ = mysql.GetCommentList(request.SearchCommentParams{
		PageInfo: request.PageInfo{
			Page:     1,
			PageSize: 5,
		},
		OrderKey: "",
		Desc:     false,
	})
	return
}

func CreatePost(p request.CreatePostParams) (err error) {
	err = mysql.CreatePost(p)
	return
}

func PostViewUpdate(p request.CreatePostParams) (err error) {
	err = mysql.PostViewUpdate(p)
	return
}

func PostCollect(p request.CreatePostParams, userId uuid.UUID) (err error) {
	err = mysql.PostCollect(p, userId)
	return
}

func PostUnCollect(p request.CreatePostParams, userId uuid.UUID) (err error) {
	err = mysql.PostUnCollect(p, userId)
	return
}

func UpdatePost(p request.CreatePostParams) (err error) {
	err = mysql.UpdatePost(p)
	return
}

func UpdatePostStatus(p request.CreatePostParams) (err error) {
	if err = mysql.UpdatePostStatus(p); err != nil {
		return err
	}
	// 添加10积分并通知
	if p.Status == "published" {
		// 初始化积分通知
		err, userReward := mysql.GetSettingItem(model.XdSetting{
			Slug: "user_reward",
		})
		var reward response.ResUserReward
		if err = json.Unmarshal([]byte(userReward.Value), &reward); err != nil {
			return response.ErrorPostUpdate
		}
		if reward.PublishPostType == "cash" {
			err = mysql.UpdateUserCash(reward.PublishPostCredit, p.Author)
			if err == nil {
				CreateNotification(p.Author, "", p.PostId, 0, reward.PublishPostCredit, 0, "publishPost", "")
			}
		} else {
			err = mysql.UpdateUserCredit(reward.PublishPostCredit, p.Author)
			if err == nil {
				CreateNotification(p.Author, "", p.PostId, reward.PublishPostCredit, 0, 0, "publishPost", "")
			}
		}
	}
	return
}

func DeletePost(p model.XdPost) (err error) {
	err = mysql.DeletePost(p)
	return
}
