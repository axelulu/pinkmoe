/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:23
 * @FilePath: /pinkmoe_server/model/response/home.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package response

import "server/model"

type ShopConfig struct {
	Popular   []string  `json:"popular"`
	Recommend []string  `json:"recommend"`
	Cms       []ShopCms `json:"cms"`
}

type ShopCms struct {
	Category string `json:"category"`
	Style    int    `json:"style"`
}

type ShopSettings struct {
	ParsingUrl string `json:"parsingUrl"`
}

type ShopContent struct {
	Slug  string          `json:"slug"`  // 菜单标识
	Name  string          `json:"name"`  // 菜单名称
	Icon  string          `json:"icon"`  // 菜单图标
	Sort  int             `json:"sort"`  // 排序标记
	Style int             `json:"style"` // 分类模块显示类型
	Goods []model.XdGoods `json:"goods"` // 文章
}

type Shop struct {
	Popular      []model.XdGoods         `json:"popular"`
	ShopCategory []model.XdGoodsCategory `json:"shopCategory"`
	Content      []ShopContent           `json:"content"`
}
