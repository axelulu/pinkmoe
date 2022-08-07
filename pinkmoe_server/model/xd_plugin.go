/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:56:14
 * @FilePath: /pinkmoe_server/model/xd_plugin.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package model

import "server/global"

type XdPlugin struct {
	global.XD_MODEL
	StoreId       int    `json:"storeId" gorm:"comment:插件在商城中的id"`  // 插件在商城中的id
	PName         string `json:"pName" gorm:"comment:插件名称英文"`       // 插件名称英文
	PTitle        string `json:"pTitle" gorm:"comment:插件名称"`        // 插件名称
	PDescription  string `json:"pDescription" gorm:"comment:插件介绍"`  // 插件介绍
	PAuth         string `json:"pAuth" gorm:"comment:作者"`           // 作者
	IsInstall     int    `json:"isInstall" gorm:"comment:是否安装"`     // 是否安装
	Status        int    `json:"status" gorm:"comment:状态"`          // 状态
	Version       string `json:"version" gorm:"comment:当前版本"`       // 当前版本
	Price         uint   `json:"price" gorm:"comment:价格"`           // 价格
	DownloadTimes uint   `json:"downloadTimes" gorm:"comment:下载次数"` // 下载次数
	InstallPath   string `json:"installPath" gorm:"comment:安装路径"`   // 安装路径
}
