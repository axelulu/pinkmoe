/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:59
 * @FilePath: /pinkmoe_server/initialize/initDB/comment_smile.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package initDB

import (
	"server/global"
	"server/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var CommentSmile = new(commentSmile)

type commentSmile struct{}

func (u *commentSmile) TableName() string {
	return "xd_posts"
}

func (u *commentSmile) Initialize() error {
	entities := []model.XdCommentSmile{
		{
			Name: "bcr_no",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/no.webp",
			Type: "smile",
		},
		{
			Name: "bcr_ok",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/ok.webp",
			Type: "smile",
		},
		{
			Name: "bcr_一切都会好起来的",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/一切都会好起来的.webp",
			Type: "smile",
		},
		{
			Name: "bcr_了解",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/了解.webp",
			Type: "smile",
		},
		{
			Name: "bcr_交给我吧",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/交给我吧.webp",
			Type: "smile",
		},
		{
			Name: "bcr_做得好",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/做得好.webp",
			Type: "smile",
		},
		{
			Name: "bcr_加油",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/加油.webp",
			Type: "smile",
		},
		{
			Name: "bcr_参战",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/参战.webp",
			Type: "smile",
		},
		{
			Name: "bcr_吸溜",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/吸溜.webp",
			Type: "smile",
		},
		{
			Name: "bcr_呜呼呼",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/呜呼呼.webp",
			Type: "smile",
		},
		{
			Name: "bcr_哇哇",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/哇哇.webp",
			Type: "smile",
		},
		{
			Name: "bcr_哈罗",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/哈罗.webp",
			Type: "smile",
		},
		{
			Name: "bcr_喵喵喵",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/喵喵喵.webp",
			Type: "smile",
		},
		{
			Name: "bcr_在那里",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/在那里.webp",
			Type: "smile",
		},
		{
			Name: "bcr_已经举报了",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/已经举报了.webp",
			Type: "smile",
		},
		{
			Name: "bcr_干杯",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/干杯.webp",
			Type: "smile",
		},
		{
			Name: "bcr_快回复吧",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/快回复吧.webp",
			Type: "smile",
		},
		{
			Name: "bcr_抱歉啦",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/抱歉啦.webp",
			Type: "smile",
		},
		{
			Name: "bcr_接下来拜托了",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/接下来拜托了.webp",
			Type: "smile",
		},
		{
			Name: "bcr_早上好",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/早上好.webp",
			Type: "smile",
		},
		{
			Name: "bcr_晚上好",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/晚上好.webp",
			Type: "smile",
		},
		{
			Name: "bcr_晚安zzz",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/晚安zzz.webp",
			Type: "smile",
		},
		{
			Name: "bcr_来吧",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/来吧.webp",
			Type: "smile",
		},
		{
			Name: "bcr_气呼呼",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/气呼呼.webp",
			Type: "smile",
		},
		{
			Name: "bcr_真是难办啊",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/真是难办啊.webp",
			Type: "smile",
		},
		{
			Name: "bcr_糟糕",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/糟糕.webp",
			Type: "smile",
		},
		{
			Name: "bcr_给我等一下",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/给我等一下.webp",
			Type: "smile",
		},
		{
			Name: "bcr_给我记住了",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/给我记住了.webp",
			Type: "smile",
		},
		{
			Name: "bcr_要上了哦",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/要上了哦.webp",
			Type: "smile",
		},
		{
			Name: "bcr_谢谢",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/谢谢.webp",
			Type: "smile",
		},
		{
			Name: "bcr_跟随妾身而来吧",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/跟随妾身而来吧.webp",
			Type: "smile",
		},
		{
			Name: "bcr_颤颤抖抖",
			Url:  "https://cdn.inn-studio.com/themes/zero/images/bcr/颤颤抖抖.webp",
			Type: "smile",
		},
		{
			Name: "滑稽",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1f0wwjnedhoj200u00ugld.jpg",
			Type: "smile",
		},
		{
			Name: "感谢分享",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5ir4rznj202s02eglh.jpg",
			Type: "smile",
		},
		{
			Name: "脸红",
			Url:  "https://sinaimg.inn-studio.com/large/686ee05djw1eu8ijxc3p7g201c01c3yd.gif",
			Type: "smile",
		},
		{
			Name: "杯具",
			Url:  "https://sinaimg.inn-studio.com/large/686ee05djw1eu8ikpw34jg201e01emx1.gif",
			Type: "smile",
		},
		{
			Name: "亚历山大",
			Url:  "https://sinaimg.inn-studio.com/large/686ee05djw1eu8iliwosmg201e01e74h.gif",
			Type: "smile",
		},
		{
			Name: "想要",
			Url:  "https://sinaimg.inn-studio.com/large/686ee05djw1eu8ilzci2jg202s02sglo.gif",
			Type: "smile",
		},
		{
			Name: "吃惊",
			Url:  "https://sinaimg.inn-studio.com/large/686ee05djw1eu8j1vay4ej204h049jrb.jpg",
			Type: "smile",
		},
		{
			Name: "好样的",
			Url:  "https://sinaimg.inn-studio.com/large/686ee05djw1eu8iomh5cbg203g03cdgx.gif",
			Type: "smile",
		},
		{
			Name: "感谢分享2",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1f44hq6g2ftj202s02swef.jpg",
			Type: "smile",
		},
		{
			Name: "生气",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1eu8fzbvagqj202s02sa9z.jpg",
			Type: "smile",
		},
		{
			Name: "卖萌",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1eu8g1m0y49j202s02sjrb.jpg",
			Type: "smile",
		},
		{
			Name: "OK！",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1eu8fza5ep2j202s02s3yg.jpg",
			Type: "smile",
		},
		{
			Name: "不行",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1eu8fz69u1qj202s02st8n.jpg",
			Type: "smile",
		},
		{
			Name: "叹气",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5jxgin9j202s02e743.jpg",
			Type: "smile",
		},
		{
			Name: "棒！",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1eu8fz1nm1rj202s02smx3.jpg",
			Type: "smile",
		},
		{
			Name: "偷笑",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1f124864q4qj202s02e3ye.jpg",
			Type: "smile",
		},
		{
			Name: "噫",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1f0z83hc4ndj202s02ea9x.jpg",
			Type: "smile",
		},
		{
			Name: "good",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1f0z83fr4l7j202s02emx2.jpg",
			Type: "smile",
		},
		{
			Name: "不明真相",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5fo2790j202s02eweb.jpg",
			Type: "smile",
		},
		{
			Name: "太棒了",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5lerczej202s02edfn.jpg",
			Type: "smile",
		},
		{
			Name: "不知所措",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5qzpg9aj202s02eq2u.jpg",
			Type: "smile",
		},
		{
			Name: "盯..",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5pi74hgj202s02et8i.jpg",
			Type: "smile",
		},
		{
			Name: "所以呢",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1f0z838j5d4j202s02emx1.jpg",
			Type: "smile",
		},
		{
			Name: "残念",
			Url:  "https://sinaimg.inn-studio.com/large/c524f7d4jw1fbk5hkglbsj202s02e3yf.jpg",
			Type: "smile",
		},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (u *commentSmile) CheckDataExist() bool {
	if errors.Is(global.XD_DB.Where("name = ?", "残念").First(&model.XdCommentSmile{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
