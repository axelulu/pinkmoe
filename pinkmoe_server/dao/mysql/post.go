/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:16
 * @FilePath: /pinkmoe_server/dao/mysql/post.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"encoding/json"
	"errors"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func GetTopicPostList(info request.SearchTopicParams) (err error, list model.XdTopic, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdTopic{})
	if info.Desc {
		if err = db.Where("value = ?", info.Value).Preload("PostRelations", func(db *gorm.DB) *gorm.DB {
			err = db.Count(&total).Error
			return db.Where("status = ?", "published").Order(info.OrderKey + " DESC").Preload("CategoryRelation").Limit(limit).Offset(offset)
		}).First(&list).Error; err != nil {
			return response.ErrorPostListGet, list, 0
		}
	} else {
		if err = db.Where("value = ?", info.Value).Preload("PostRelations", func(db *gorm.DB) *gorm.DB {
			err = db.Count(&total).Error
			return db.Where("status = ?", "published").Order(info.OrderKey + " ASC").Preload("CategoryRelation").Limit(limit).Offset(offset)
		}).First(&list).Error; err != nil {
			return response.ErrorPostListGet, list, 0
		}
	}
	return
}

func GetCategoryPostList(info request.SearchPostParams) (err error, list []model.XdPost, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdPost{})
	err, treeMap := getAllChildCategoryBySlug(info.Category)
	var allSlug []string
	for _, category := range treeMap {
		allSlug = append(allSlug, category.Slug)
	}

	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}

	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	var order string

	if info.Desc {
		order = info.OrderKey + " DESC"
	} else {
		order = info.OrderKey + " ASC"
	}
	if info.Category == "0" || info.Category == "" {
		if err = db.Where("status = ?", "published").Preload("AuthorRelation").Preload("CategoryRelation").Order(order).Limit(limit).Offset(offset).Find(&list).Error; err != nil {
			return response.ErrorPostListGet, list, 0
		}
	} else {
		if err = db.Where("status = ?", "published").Preload("AuthorRelation").Preload("CategoryRelation").Order(order).Limit(limit).Offset(offset).Where("category IN ?", allSlug).Find(&list).Error; err != nil {
			return response.ErrorPostListGet, list, 0
		}
	}
	return
}

func GetCollectPost(p request.PageInfo, userId uuid.UUID) (err error, post []model.XdPostCollect, total int64) {
	limit := p.PageSize
	offset := p.PageSize * (p.Page - 1)
	if err = global.XD_DB.Where("uuid = ?", userId).Preload("PostIdRelation").Preload("PostIdRelation.StarRelation").Preload("PostIdRelation.DownloadRelation").Limit(limit).Offset(offset).Find(&post).Error; err != nil {
		return response.ErrorPostCollect, nil, 0
	}
	return
}

func GetAdminPostList(info request.SearchPostParams, userId string) (err error, post []model.XdPost, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdPost{})

	var order string

	if len(info.OrderKey) <= 0 {
		info.OrderKey = "updated_at"
	}

	if info.Desc {
		order = info.OrderKey + " DESC"
	} else {
		order = info.OrderKey + " ASC"
	}

	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}

	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}

	if info.Category != "" {
		db = db.Where("category = ?", info.Category)
	}

	if info.Author != "" {
		db = db.Where("author = ?", info.Author)
	}

	if info.CommentStatus != "" {
		db = db.Where("comment_status = ?", info.CommentStatus)
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorPostListGet, nil, 0
	}
	if err = db.Limit(limit).Offset(offset).Order(order).Preload("CollectRelation").Preload("AuthorRelation").Preload("TopicRelations").Preload("CategoryRelation").Preload("PostImg").Preload("DownloadRelation").Preload("MusicRelation").Preload("VideoRelation").Preload("FileRelation").Find(&post).Error; err != nil {
		return response.ErrorPostListGet, nil, 0
	}
	return
}

func GetPostList(info request.SearchPostParams, userId string) (err error, post []model.XdPost, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdPost{})

	db = db.Where("status = 'published'")

	var order string

	if info.Desc {
		order = info.OrderKey + " DESC"
	} else {
		order = info.OrderKey + " ASC"
	}

	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}

	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	if info.Type != "" {
		if info.Type == "focus" {
			var follows []model.XdFollow
			global.XD_DB.Where("form_uid = ?", userId).Find(&follows)
			var fans []string
			for _, follow := range follows {
				fans = append(fans, follow.ToUid.String())
			}
			fans = append(fans, info.Author)
			db = db.Where("author IN (?)", fans)
		} else {
			db = db.Where("type = ?", info.Type)
		}
	}

	if info.Category != "" {
		db = db.Where("category = ?", info.Category)
	}

	if info.Author != "" && info.Type != "focus" {
		db = db.Where("author = ?", info.Author)
	}

	if info.CommentStatus != "" {
		db = db.Where("comment_status = ?", info.CommentStatus)
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorPostListGet, nil, 0
	}
	if err = db.Limit(limit).Offset(offset).Order(order).Preload("CollectRelation", func(db *gorm.DB) *gorm.DB {
		return db.Where("uuid = ?", userId)
	}).Preload("AuthorRelation").Preload("AuthorRelation.Authority").Preload("TopicRelations").Preload("CategoryRelation").Preload("PostImg").Preload("MusicRelation", func(db *gorm.DB) *gorm.DB {
		return db.Select("post_id", "price", "price_type", "name")
	}).Preload("VideoRelation", func(db *gorm.DB) *gorm.DB {
		return db.Select("post_id", "price", "price_type", "name")
	}).Preload("DownloadRelation", func(db *gorm.DB) *gorm.DB {
		return db.Select("post_id", "price", "price_type", "name")
	}).Preload("FileRelation").Find(&post).Error; err != nil {
		return response.ErrorPostListGet, nil, 0
	}
	return
}

func GetUserPostList(info request.SearchPostParams, userId uuid.UUID) (err error, post []response.XdPost, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&response.XdPost{})

	db = db.Where("author = ?", userId)

	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}

	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}

	if info.Type != "" {
		db = db.Where("type LIKE ?", "%"+info.Type+"%")
	}

	if info.Category != "" {
		db = db.Where("category = ?", info.Category)
	}

	if info.Status != "" {
		db = db.Where("status LIKE ?", "%"+info.Status+"%")
	}

	if info.CommentStatus != "" {
		db = db.Where("comment_status LIKE ?", "%"+info.CommentStatus+"%")
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorPostListGet, nil, 0
	}
	if err = db.Limit(limit).Offset(offset).Preload("AuthorRelation").Preload("TopicRelations").Preload("CategoryRelation").Preload("PostImg").Preload("DownloadRelation").Preload("MusicRelation").Preload("VideoRelation").Preload("CommentRelation").Preload("StarRelation").Preload("CollectRelation").Find(&post).Error; err != nil {
		return response.ErrorPostListGet, nil, 0
	}
	return
}

func GetPostByPostIds(postIds []string) (err error, list []model.XdPost) {
	if err := global.XD_DB.Where("post_id IN ?", postIds).Preload("AuthorRelation").Preload("CategoryRelation").Find(&list).Error; err != nil {
		return response.ErrorPostListGet, nil
	}
	return err, list
}

func GetPostByPostId(postId string) (err error, posts response.XdPost) {
	if err := global.XD_DB.Where("post_id = ?", postId).Preload("CategoryRelation").Preload("TopicRelations").Preload("AuthorRelation").Preload("AuthorRelation.Authority").Preload("PostImg").Preload("MusicRelation", func(db *gorm.DB) *gorm.DB {
		return db.Select("post_id", "price", "price_type", "name")
	}).Preload("VideoRelation", func(db *gorm.DB) *gorm.DB {
		return db.Select("post_id", "price", "price_type", "name")
	}).Preload("DownloadRelation", func(db *gorm.DB) *gorm.DB {
		return db.Select("post_id", "price", "price_type", "name")
	}).First(&posts).Error; err != nil {
		return response.ErrorPostListGet, posts
	}
	return err, posts
}

// 下载
func BuyPostDownload(p request.ReqDownloadBuy, uuid uuid.UUID, ip string) (err error, credit int, price int, priceSlug string) {
	var download model.XdPostDownload
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Where("post_id = ?", p.PostId).Where("id = ?", p.DownloadId).First(&download).Error; err != nil {
			return response.ErrorPostDownloadBuy
		}
		var user model.XdUser
		if err = tx.Where("uuid = ?", uuid).First(&user).Error; err != nil {
			return response.ErrorPostDownloadBuy
		}
		if download.PriceType == "credit" {
			user.Credit = user.Credit - download.Price
			if user.Credit < 0 {
				return response.ErrorPostDownloadCreditBuy
			}
			priceSlug = "credit"
			credit = user.Credit
		} else if download.PriceType == "cash" {
			user.Cash = user.Cash - download.Price
			if user.Cash < 0 {
				return response.ErrorPostDownloadCashBuy
			}
			priceSlug = "cash"
			credit = user.Cash
		}
		if err = tx.Create(&model.XdOrder{
			XD_MODEL:     global.XD_MODEL{},
			OrderId:      global.XD_SNOWFLAKE.Generate().String(),
			Price:        download.Price,
			MoneyType:    download.PriceType,
			PaymentState: "paid",
			//PaidAt:       global.XdTime.UTC(),
			UserIp:  ip,
			PostId:  p.PostId,
			Type:    "download",
			Uuid:    uuid,
			GoodsId: download.ID,
		}).Error; err != nil {
			return response.ErrorPostDownloadBuy
		}
		if err = tx.Model(&model.XdUser{}).Where("uuid = ?", uuid).Update(priceSlug, credit).Error; err != nil {
			return response.ErrorPostDownloadBuy
		}
		return nil
	}); err != nil {
		return response.ErrorPostDownloadBuy, credit, download.Price, priceSlug
	}
	return err, credit, download.Price, priceSlug
}

func GetPostDownloadByPostId(postId string, uuid uuid.UUID) (err error, postDownload []response.XdPostDownload) {
	if err := global.XD_DB.Where("post_id = ?", postId).Find(&postDownload).Error; err != nil {
		return response.ErrorPostListGet, postDownload
	}
	for i := 0; i < len(postDownload); i++ {
		postDownload[i].Buy = 1
		var order model.XdOrder
		if errors.Is(global.XD_DB.Where(map[string]interface{}{
			"post_id":       postId,
			"type":          "download",
			"uuid":          uuid,
			"payment_state": "paid",
			"goods_id":      postDownload[i].ID,
		}).First(&order).Error, gorm.ErrRecordNotFound) && postDownload[i].Price != 0 {
			postDownload[i].Buy = 0
			postDownload[i].Url = ""
			postDownload[i].ExtractPwd = ""
			postDownload[i].UnpackPwd = ""
		}
	}
	return err, postDownload
}

// 音乐
func BuyPostMusic(p request.ReqMusicBuy, uuid uuid.UUID, ip string) (err error, credit int, price int, priceSlug string) {
	var download model.XdPostMusic
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Where("post_id = ?", p.PostId).Where("id = ?", p.MusicId).First(&download).Error; err != nil {
			return response.ErrorPostMusicBuy
		}
		var user model.XdUser
		if err = tx.Where("uuid = ?", uuid).First(&user).Error; err != nil {
			return response.ErrorPostMusicBuy
		}
		if download.PriceType == "credit" {
			user.Credit = user.Credit - download.Price
			if user.Credit < 0 {
				return response.ErrorPostMusicCreditBuy
			}
			priceSlug = "credit"
			credit = user.Credit
		} else if download.PriceType == "cash" {
			user.Cash = user.Cash - download.Price
			if user.Cash < 0 {
				return response.ErrorPostMusicCashBuy
			}
			priceSlug = "cash"
			credit = user.Cash
		}
		if err = tx.Create(&model.XdOrder{
			XD_MODEL:     global.XD_MODEL{},
			OrderId:      global.XD_SNOWFLAKE.Generate().String(),
			Price:        download.Price,
			MoneyType:    download.PriceType,
			PaymentState: "paid",
			//PaidAt:       global.XdTime.UTC(),
			UserIp:  ip,
			PostId:  p.PostId,
			Type:    "music",
			Uuid:    uuid,
			GoodsId: download.ID,
		}).Error; err != nil {
			return response.ErrorPostMusicBuy
		}
		if err = tx.Model(&model.XdUser{}).Where("uuid = ?", uuid).Update(priceSlug, credit).Error; err != nil {
			return response.ErrorPostMusicBuy
		}
		return nil
	}); err != nil {
		return response.ErrorPostMusicBuy, credit, download.Price, priceSlug
	}
	return err, credit, download.Price, priceSlug
}

func GetPostMusicByPostId(postId string, uuid uuid.UUID) (err error, postMusic []response.XdPostMusic) {
	if err := global.XD_DB.Where("post_id = ?", postId).Find(&postMusic).Error; err != nil {
		return response.ErrorPostListGet, postMusic
	}
	for i := 0; i < len(postMusic); i++ {
		postMusic[i].Buy = 1
		var order model.XdOrder
		if errors.Is(global.XD_DB.Where(map[string]interface{}{
			"post_id":       postId,
			"type":          "music",
			"uuid":          uuid,
			"payment_state": "paid",
			"goods_id":      postMusic[i].ID,
		}).First(&order).Error, gorm.ErrRecordNotFound) && postMusic[i].Price != 0 {
			postMusic[i].Buy = 0
			postMusic[i].Url = ""
		}
	}
	return err, postMusic
}

// 视频
func BuyPostVideo(p request.ReqVideoBuy, uuid uuid.UUID, ip string) (err error, credit int, price int, priceSlug string) {
	var download model.XdPostVideo
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Where("post_id = ?", p.PostId).Where("id = ?", p.VideoId).First(&download).Error; err != nil {
			return response.ErrorPostVideoBuy
		}
		var user model.XdUser
		if err = tx.Where("uuid = ?", uuid).First(&user).Error; err != nil {
			return response.ErrorPostVideoBuy
		}
		if download.PriceType == "credit" {
			user.Credit = user.Credit - download.Price
			if user.Credit < 0 {
				return response.ErrorPostVideoCreditBuy
			}
			priceSlug = "credit"
			credit = user.Credit
		} else if download.PriceType == "cash" {
			user.Cash = user.Cash - download.Price
			if user.Cash < 0 {
				return response.ErrorPostVideoCashBuy
			}
			priceSlug = "cash"
			credit = user.Cash
		}
		if err = tx.Create(&model.XdOrder{
			XD_MODEL:     global.XD_MODEL{},
			OrderId:      global.XD_SNOWFLAKE.Generate().String(),
			Price:        download.Price,
			MoneyType:    download.PriceType,
			PaymentState: "paid",
			//PaidAt:       global.XdTime.UTC(),
			UserIp:  ip,
			PostId:  p.PostId,
			Type:    "video",
			Uuid:    uuid,
			GoodsId: download.ID,
		}).Error; err != nil {
			return response.ErrorPostVideoBuy
		}
		if err = tx.Model(&model.XdUser{}).Where("uuid = ?", uuid).Update(priceSlug, credit).Error; err != nil {
			return response.ErrorPostVideoBuy
		}
		return nil
	}); err != nil {
		return response.ErrorPostVideoBuy, credit, download.Price, priceSlug
	}
	return err, credit, download.Price, priceSlug
}

func GetPostVideoByPostId(postId string, uuid uuid.UUID) (err error, postVideo []response.XdPostVideo) {
	if err := global.XD_DB.Where("post_id = ?", postId).Find(&postVideo).Error; err != nil {
		return response.ErrorPostListGet, postVideo
	}
	err, settings := GetSettingItem(model.XdSetting{Slug: "website_basic"})
	var setting response.Settings
	json.Unmarshal([]byte(settings.Value), &setting)
	for i := 0; i < len(postVideo); i++ {
		postVideo[i].Buy = 1
		postVideo[i].Url = setting.ParsingUrl + postVideo[i].Url
		var order model.XdOrder
		if errors.Is(global.XD_DB.Where(map[string]interface{}{
			"post_id":       postId,
			"type":          "video",
			"uuid":          uuid,
			"payment_state": "paid",
			"goods_id":      postVideo[i].ID,
		}).First(&order).Error, gorm.ErrRecordNotFound) && postVideo[i].Price != 0 {
			postVideo[i].Buy = 0
			postVideo[i].Url = ""
		}
	}
	return err, postVideo
}

func GetAuthorPostCount(uuid uuid.UUID) (err error, total int64) {
	if err = global.XD_DB.Model(model.XdPost{}).Where("author = ?", uuid).Count(&total).Error; err != nil {
		return response.ErrorCommentCreate, 0
	}
	return
}

func CreatePost(p request.CreatePostParams) (err error) {
	var post model.XdPost
	if len(p.Category) > 0 {
		post = model.XdPost{
			PostId:           global.XD_SNOWFLAKE.Generate().String(),
			Title:            p.Title,
			Exerpt:           p.Exerpt,
			Content:          p.Content,
			Category:         p.Category,
			Author:           p.Author,
			Cover:            p.Cover,
			Type:             p.Type,
			VideoRelation:    p.VideoRelation,
			MusicRelation:    p.MusicRelation,
			DownloadRelation: p.DownloadRelation,
			CommentStatus:    p.CommentStatus,
			Status:           p.Status,
		}
	} else {
		post = model.XdPost{
			PostId:           global.XD_SNOWFLAKE.Generate().String(),
			Title:            p.Title,
			Exerpt:           p.Exerpt,
			Content:          p.Content,
			Author:           p.Author,
			Cover:            p.Cover,
			Type:             p.Type,
			VideoRelation:    p.VideoRelation,
			MusicRelation:    p.MusicRelation,
			DownloadRelation: p.DownloadRelation,
			CommentStatus:    p.CommentStatus,
			Status:           p.Status,
		}
	}
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&post).Error; err != nil {
			return response.ErrorPostCreate
		}
		for _, img := range p.PostImg {
			if err = tx.Model(&model.XdUploadFile{}).Where("uuid = ?", img.Uuid).Where("url = ?", img.Url).Updates(map[string]interface{}{
				"type":    "post",
				"post_id": post.PostId,
			}).Error; err != nil {
				return response.ErrorPostCreate
			}
		}
		for _, t := range p.Topic {
			newTopic := model.XdTopic{
				XD_MODEL: global.XD_MODEL{},
				Value:    t,
				Label:    t,
				Icon:     "https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg",
				Sort:     0,
			}
			var topic model.XdTopic
			if errors.Is(tx.Where("value = ?", t).First(&topic).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
				err = tx.Create(&newTopic).Error
				err = tx.Create(&model.XdTopicRelation{
					XdPostId:  post.PostId,
					XdTopicId: newTopic.ID,
				}).Error
			} else {
				err = tx.Create(&model.XdTopicRelation{
					XdPostId:  post.PostId,
					XdTopicId: topic.ID,
				}).Error
			}
		}
		return nil
	}); err != nil {
		return response.ErrorPostCreate
	}
	return err
}

func UpdatePost(p request.CreatePostParams) (err error) {
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(&[]model.XdTopicRelation{}, "xd_post_post_id = ?", p.PostId).Error; err != nil {
			return response.ErrorPostTopicDelete
		}
		for _, t := range p.Topic {
			newTopic := model.XdTopic{
				XD_MODEL: global.XD_MODEL{},
				Value:    t,
				Label:    t,
				Icon:     "https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg",
				Sort:     0,
			}
			var topic model.XdTopic
			if errors.Is(tx.Where("value = ?", t).First(&topic).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
				err = tx.Create(&newTopic).Error
				err = tx.Create(&model.XdTopicRelation{
					XdPostId:  p.PostId,
					XdTopicId: newTopic.ID,
				}).Error
			} else {
				err = tx.Create(&model.XdTopicRelation{
					XdPostId:  p.PostId,
					XdTopicId: topic.ID,
				}).Error
			}
		}
		if err = tx.Where("post_id = ?", p.PostId).Delete(&model.XdPostDownload{}).Error; err != nil {
			return response.ErrorPostUpdate
		}
		if len(p.DownloadRelation) > 0 {
			if err = tx.Create(&p.DownloadRelation).Error; err != nil {
				return response.ErrorPostUpdate
			}
		}
		if err = tx.Where("post_id = ?", p.PostId).Delete(&model.XdPostMusic{}).Error; err != nil {
			return response.ErrorPostUpdate
		}
		if len(p.MusicRelation) > 0 {
			if err = tx.Create(&p.MusicRelation).Error; err != nil {
				return response.ErrorPostUpdate
			}
		}
		if err = tx.Where("post_id = ?", p.PostId).Delete(&model.XdPostVideo{}).Error; err != nil {
			return response.ErrorPostUpdate
		}
		if len(p.VideoRelation) > 0 {
			if err = tx.Create(&p.VideoRelation).Error; err != nil {
				return response.ErrorPostUpdate
			}
		}
		if err = tx.Where("post_id = ?", p.PostId).Updates(&model.XdPost{
			PostId:        p.PostId,
			Title:         p.Title,
			Exerpt:        p.Exerpt,
			Content:       p.Content,
			Category:      p.Category,
			Author:        p.Author,
			Cover:         p.Cover,
			Type:          p.Type,
			From:          p.From,
			CommentStatus: p.CommentStatus,
			Status:        p.Status,
		}).Error; err != nil {
			return response.ErrorPostUpdate
		}
		return nil
	}); err != nil {
		return response.ErrorPostUpdate
	}
	return err
}

func UpdatePostStatus(p request.CreatePostParams) (err error) {
	if err = global.XD_DB.Model(&model.XdPost{}).Where("post_id = ?", p.PostId).Update("status", p.Status).Error; err != nil {
		return response.ErrorPostUpdate
	}
	return err
}

func PostViewUpdate(p request.CreatePostParams) (err error) {
	var post model.XdPost
	if err = global.XD_DB.Model(&model.XdPost{}).Where("post_id = ?", p.PostId).First(&post).Error; err != nil {
		return response.ErrorPostUpdate
	}
	if err = global.XD_DB.Model(&model.XdPost{}).Where("post_id = ?", p.PostId).UpdateColumns(model.XdPost{
		View: post.View + 1,
		//UpdatedAt: found.UpdatedAt,
	}).Update("view", post.View+1).Error; err != nil {
		return response.ErrorPostUpdate
	}
	return err
}

func PostCollect(p request.CreatePostParams, userId uuid.UUID) (err error) {
	var collect model.XdPostCollect
	if !errors.Is(global.XD_DB.Where("post_id = ?", p.PostId).First(&collect).Error, gorm.ErrRecordNotFound) {
		return response.ErrorPostCollected
	}
	if err = global.XD_DB.Create(&model.XdPostCollect{
		PostId: p.PostId,
		Uuid:   userId,
	}).Error; err != nil {
		return response.ErrorPostCollected
	}
	return err
}

func PostUnCollect(p request.CreatePostParams, userId uuid.UUID) (err error) {
	var collect model.XdPostCollect
	if errors.Is(global.XD_DB.Where("post_id = ?", p.PostId).First(&collect).Error, gorm.ErrRecordNotFound) {
		return response.ErrorPostUnCollected
	}
	if err = global.XD_DB.Delete(&[]model.XdPostCollect{}, "post_id = ?", p.PostId).Error; err != nil {
		return response.ErrorPostUnCollected
	}
	return err
}

func DeletePost(p model.XdPost) (err error) {
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(&[]model.XdTopicRelation{}, "xd_post_post_id = ?", p.PostId).Error; err != nil {
			return response.ErrorPostTopicDelete
		}
		if err = tx.Delete(&p).Error; err != nil {
			return response.ErrorPostDelete
		}
		return nil
	}); err != nil {
		return response.ErrorPostDelete
	}
	return
}
