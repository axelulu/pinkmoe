/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:42
 * @FilePath: /pinkmoe_server/model/response/xd_post_video.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package response

import (
	"server/global"
)

type XdPostVideo struct {
	ID        uint          `gorm:"primarykey"` // 主键ID
	CreatedAt global.XdTime // 创建时间
	UpdatedAt global.XdTime // 更新时间
	PostId    string        `json:"postId" form:"postId" gorm:"not null;comment:文章ID"` // 文章ID
	Price     int64         `json:"price" form:"price" gorm:"comment:链接价格"`            // 链接价格
	PriceType string        `json:"priceType" form:"priceType" gorm:"comment:链接价格类型"`  // 链接价格类型
	Name      string        `json:"name" form:"name" gorm:"comment:视频名称"`              // 视频名称
	Url       string        `json:"url" form:"url" gorm:"comment:视频链接"`                // 视频链接
	Subtitles string        `json:"subtitles" form:"subtitles" gorm:"comment:视频字幕"`    // 视频字幕
	Buy       int           `json:"buy" form:"buy" gorm:"default:1;comment:已购买"`       // 已购买
}
