/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:48
 * @FilePath: /pinkmoe_server/dao/mysql/comment.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func GetCommentList(info request.SearchCommentParams) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&response.XdComment{})
	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	if info.Type != "" {
		db = db.Where("type LIKE ?", "%"+info.Type+"%")
	}

	if info.Status != "" {
		db = db.Where("status LIKE ?", "%"+info.Status+"%")
	}

	if info.PostId != "" {
		db = db.Where("post_id = ?", info.PostId)
	}

	if info.FormUid != "" {
		db = db.Where("form_uid = ?", info.FormUid)
	}

	if len(info.OrderKey) > 0 {
		if info.Desc {
			db = db.Order(info.OrderKey + " desc")
		} else {
			db = db.Order(info.OrderKey + " asc")
		}
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorCommentListGet, nil, 0
	}
	var comment []response.XdComment
	if err = db.Limit(limit).Offset(offset).Preload("PostRelation").Preload("FormUidRelation").Preload("FormUidRelation.Authority").Preload("ToUidRelation").Preload("PostRelation").Find(&comment).Error; err != nil {
		return response.ErrorCommentListGet, nil, 0
	}
	return err, comment, total
}

func GetCommentTreeList(info request.SearchCommentParams) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdComment{})
	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	if info.PostId != "" {
		db = db.Where("post_id = ?", info.PostId)
	}

	if info.Type != "" {
		db = db.Where("type LIKE ?", "%"+info.Type+"%")
	}

	if info.Status != "" {
		db = db.Where("status LIKE ?", "%"+info.Status+"%")
	}

	if len(info.OrderKey) > 0 {
		if info.Desc {
			db = db.Order(info.OrderKey + " desc")
		} else {
			db = db.Order(info.OrderKey + " asc")
		}
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorCommentListGet, nil, 0
	}
	var comment []model.XdComment
	if err = db.Limit(limit).Offset(offset).Preload("FormUidRelation").Preload("FormUidRelation.Authority").Preload("ToUidRelation").Find(&comment).Error; err != nil {
		return response.ErrorCommentListGet, nil, 0
	}
	// 替换smile表情
	var commentSmile []model.XdCommentSmile
	if err = global.XD_DB.Model(&model.XdCommentSmile{}).Find(&commentSmile).Error; err != nil {
		return response.ErrorCommentListGet, nil, 0
	}
	for i, _ := range comment {
		for _, smile := range commentSmile {
			if strings.Contains(comment[i].Content, "["+smile.Name+"]") {
				comment[i].Content = strings.Replace(comment[i].Content, "["+smile.Name+"]", "<img src='"+smile.Url+"' class='w-12 h-12 mx-1' />", -1)
			}
		}
	}
	// 生成评论树
	treeMap := make(map[string][]model.XdComment)
	for _, v := range comment {
		treeMap[strconv.Itoa(int(v.ParentId))] = append(treeMap[strconv.Itoa(int(v.ParentId))], v)
	}
	commentList := treeMap["0"]
	for i := 0; i < len(commentList); i++ {
		err = getCommentTree(&commentList[i], treeMap)
	}
	return err, commentList, total
}

func getCommentTree(menu *model.XdComment, treeMap map[string][]model.XdComment) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getCommentTree(&menu.Children[i], treeMap)
	}
	if err != nil {
		return response.ErrorCommentListGet
	}
	return err
}

func GetAuthorCommentCount(uuid uuid.UUID) (err error, total int64) {
	if err = global.XD_DB.Model(model.XdComment{}).Where("form_uid = ?", uuid).Count(&total).Error; err != nil {
		return response.ErrorCommentCreate, 0
	}
	return
}

func CreateComment(p model.XdComment) (err error) {
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorCommentCreate
	}
	return err
}

func UpdateComment(u model.XdComment) (err error) {
	if err = global.XD_DB.Where("id = ?", u.ID).First(&model.XdComment{}).Updates(&model.XdComment{
		PostId:          u.PostId,
		ParentId:        u.ParentId,
		FormUid:         u.FormUid,
		ToUid:           u.ToUid,
		Content:         u.Content,
		Type:            u.Type,
		Status:          u.Status,
		Children:        u.Children,
		FormUidRelation: u.FormUidRelation,
	}).Error; err != nil {
		return response.ErrorCommentUpdate
	}
	return err
}

func DeleteComment(p model.XdComment) (err error) {
	if err = global.XD_DB.Delete(&p).Error; err != nil {
		return response.ErrorCommentDelete
	}
	return
}
