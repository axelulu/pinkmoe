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
	uuid "github.com/satori/go.uuid"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"strconv"
	"strings"
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
	if err = db.Preload("FormUidRelation").Preload("FormUidRelation.Authority").Preload("ToUidRelation").Find(&comment).Error; err != nil {
		return response.ErrorCommentListGet, nil, 0
	}
	// 生成评论列表
	treeMap := make(map[string][]model.XdComment)
	treeComm := make([]model.XdComment, 0)
	treeOffsetComm := make([]model.XdComment, 0)
	for _, v := range comment {
		treeMap[strconv.Itoa(int(v.ParentId))] = append(treeMap[strconv.Itoa(int(v.ParentId))], v)
	}
	commentList := treeMap["0"]
	for i := 0; i < len(commentList); i++ {
		if commentList[i].ParentId == 0 {
			treeComm = append(treeComm, commentList[i])
			err = getCommentOffsetLimitList(&commentList[i], &treeComm, treeMap)
		}
	}

	// 评论分页处理
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	offset2 := len(treeComm) % limit
	if info.Page*info.PageSize-limit >= len(treeComm) {
		return err, nil, total
	}
	if offset2 == 0 || info.Page != len(treeComm)/limit+1 {
		treeOffsetComm = treeComm[offset : offset+limit]
	} else {
		treeOffsetComm = treeComm[offset : offset+offset2]
	}

	// 替换smile表情
	var commentSmile []model.XdCommentSmile
	if err = global.XD_DB.Model(&model.XdCommentSmile{}).Find(&commentSmile).Error; err != nil {
		return response.ErrorCommentListGet, nil, 0
	}
	for i, _ := range treeOffsetComm {
		for _, smile := range commentSmile {
			if strings.Contains(treeOffsetComm[i].Content, "["+smile.Name+"]") {
				treeOffsetComm[i].Content = strings.Replace(treeOffsetComm[i].Content, "["+smile.Name+"]", "<img src='"+smile.Url+"' class='w-12 h-12 mx-1' />", -1)
			}
		}
	}

	// 获取最终分页后的评论树
	treeMapOffset := make(map[string][]model.XdComment)
	for _, v := range treeOffsetComm {
		treeMapOffset[strconv.Itoa(int(v.ParentId))] = append(treeMapOffset[strconv.Itoa(int(v.ParentId))], v)
	}
	commentOffsetLimitList := treeMapOffset["0"]
	for i := 0; i < len(commentOffsetLimitList); i++ {
		if commentOffsetLimitList[i].ParentId == 0 {
			treeComm = append(treeComm, commentOffsetLimitList[i])
			err = getCommentOffsetLimitTree(&commentOffsetLimitList[i], treeMapOffset)
		}
	}
	return err, commentOffsetLimitList, int64(len(treeComm))
}

func getCommentOffsetLimitTree(menu *model.XdComment, treeMap map[string][]model.XdComment) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getCommentOffsetLimitTree(&menu.Children[i], treeMap)
	}
	if err != nil {
		return response.ErrorCommentListGet
	}
	return err
}

func getCommentOffsetLimitList(menu *model.XdComment, treeComm *[]model.XdComment, treeMap map[string][]model.XdComment) (err error) {
	*treeComm = append(*treeComm, treeMap[strconv.Itoa(int(menu.ID))]...)
	for i := 0; i < len(*treeComm); i++ {
		if (*treeComm)[i].ParentId == menu.ID {
			err = getCommentOffsetLimitList(&(*treeComm)[i], treeComm, treeMap)
		}
	}
	if err != nil {
		return response.ErrorCommentListGet
	}
	return err
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
