/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:01
 * @FilePath: /pinkmoe_server/initialize/initDB/menu.go
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

var BaseMenu = new(menu)

type menu struct{}

func (m *menu) TableName() string {
	return "xd_base_menus"
}

func (m *menu) Initialize() error {
	entities := []model.XdBaseMenu{
		{Hidden: false, ParentId: 0, Path: "/dashboard", Name: "dashboard", Component: "LAYOUT", Sort: 0, Meta: model.Meta{Title: "仪表盘", Icon: "DashboardOutlined"}},
		{Hidden: false, ParentId: 1, Path: "console", Name: "console", Component: "/dashboard/console/console", Sort: 0, Meta: model.Meta{Title: "控制台", Icon: ""}},
		{Hidden: false, ParentId: 1, Path: "workplace", Name: "workplace", Component: "/dashboard/workplace/workplace", Sort: 1, Meta: model.Meta{Title: "工作区", Icon: ""}},
		{Hidden: false, ParentId: 0, Path: "/system", Name: "system", Component: "LAYOUT", Sort: 1, Meta: model.Meta{Title: "系统设置", Icon: "OptionsSharp"}},
		{Hidden: false, ParentId: 4, Path: "/settings", Name: "settings", Component: "PARENT_LAYOUT", Sort: 0, Meta: model.Meta{Title: "系统设置", Icon: ""}},
		{Hidden: false, ParentId: 5, Path: "site", Name: "site", Component: "/system/settings/site/site", Sort: 1, Meta: model.Meta{Title: "站点设置", Icon: ""}},
		{Hidden: false, ParentId: 5, Path: "upload", Name: "upload", Component: "/system/settings/upload/upload", Sort: 2, Meta: model.Meta{Title: "上传设置", Icon: ""}},
		{Hidden: false, ParentId: 5, Path: "email", Name: "email", Component: "/system/settings/email/email", Sort: 3, Meta: model.Meta{Title: "邮箱设置", Icon: ""}},
		{Hidden: false, ParentId: 4, Path: "user", Name: "user", Component: "/system/user/user", Sort: 1, Meta: model.Meta{Title: "用户管理", Icon: ""}},
		{Hidden: false, ParentId: 4, Path: "role", Name: "role", Component: "/system/role/role", Sort: 2, Meta: model.Meta{Title: "角色管理", Icon: ""}},
		{Hidden: false, ParentId: 4, Path: "api", Name: "api", Component: "/system/api/api", Sort: 3, Meta: model.Meta{Title: "接口管理", Icon: ""}},
		{Hidden: false, ParentId: 4, Path: "menu", Name: "menu", Component: "/system/menu/menu", Sort: 4, Meta: model.Meta{Title: "菜单管理", Icon: ""}},
		{Hidden: false, ParentId: 0, Path: "/monitor", Name: "monitor", Component: "LAYOUT", Sort: 1, Meta: model.Meta{Title: "系统监控", Icon: "DesktopOutline"}},
		{Hidden: false, ParentId: 13, Path: "userOnline", Name: "userOnline", Component: "/monitor/userOnline/userOnline", Sort: 0, Meta: model.Meta{Title: "在线用户", Icon: ""}},
		{Hidden: false, ParentId: 13, Path: "loginLog", Name: "loginLog", Component: "/monitor/loginLog/loginLog", Sort: 1, Meta: model.Meta{Title: "登陆日志", Icon: ""}},
		{Hidden: false, ParentId: 13, Path: "operation", Name: "operation", Component: "/monitor/operation/operation", Sort: 2, Meta: model.Meta{Title: "操作日志", Icon: ""}},
		{Hidden: false, ParentId: 0, Path: "/content", Name: "content", Component: "LAYOUT", Sort: 2, Meta: model.Meta{Title: "文章内容", Icon: "ReaderOutline"}},
		{Hidden: false, ParentId: 17, Path: "post", Name: "post", Component: "/content/post/post", Sort: 0, Meta: model.Meta{Title: "文章管理", Icon: ""}},
		{Hidden: false, ParentId: 17, Path: "category", Name: "category", Component: "/content/category/category", Sort: 1, Meta: model.Meta{Title: "分类管理", Icon: ""}},
		{Hidden: false, ParentId: 17, Path: "topic", Name: "topic", Component: "/content/topic/topic", Sort: 2, Meta: model.Meta{Title: "话题管理", Icon: ""}},
		{Hidden: false, ParentId: 17, Path: "comment", Name: "comment", Component: "/content/comment/comment", Sort: 3, Meta: model.Meta{Title: "评论管理", Icon: ""}},
		{Hidden: false, ParentId: 0, Path: "/userCenter", Name: "userCenter", Component: "LAYOUT", Sort: 3, Meta: model.Meta{Title: "用户中心", Icon: "People"}},
		{Hidden: false, ParentId: 22, Path: "key", Name: "key", Component: "/userCenter/key/key", Sort: 1, Meta: model.Meta{Title: "卡密管理", Icon: ""}},
		{Hidden: false, ParentId: 22, Path: "report", Name: "report", Component: "/userCenter/report/report", Sort: 2, Meta: model.Meta{Title: "反馈管理", Icon: ""}},
		{Hidden: false, ParentId: 0, Path: "/shop", Name: "shop", Component: "LAYOUT", Sort: 4, Meta: model.Meta{Title: "商城管理", Icon: "BagHandleOutline"}},
		{Hidden: false, ParentId: 0, Path: "/forum", Name: "forum", Component: "LAYOUT", Sort: 5, Meta: model.Meta{Title: "论坛管理", Icon: "ChatbubblesOutline"}},
		{Hidden: false, ParentId: 0, Path: "/android", Name: "android", Component: "LAYOUT", Sort: 6, Meta: model.Meta{Title: "app管理", Icon: "LogoAndroid"}},
		{Hidden: false, ParentId: 0, Path: "/setting", Name: "setting", Component: "LAYOUT", Sort: 7, Meta: model.Meta{Title: "网站设置", Icon: "CogSharp"}},
		{Hidden: false, ParentId: 28, Path: "basic", Name: "basic", Component: "/setting/basic/basic", Sort: 0, Meta: model.Meta{Title: "基本设置", Icon: ""}},
		{Hidden: false, ParentId: 28, Path: "reward", Name: "reward", Component: "/setting/reward/reward", Sort: 1, Meta: model.Meta{Title: "用户奖励", Icon: ""}},
		{Hidden: false, ParentId: 28, Path: "level", Name: "level", Component: "/setting/level/level", Sort: 2, Meta: model.Meta{Title: "用户等级", Icon: ""}},
		{Hidden: false, ParentId: 28, Path: "userShop", Name: "userShop", Component: "/setting/shop/shop", Sort: 2, Meta: model.Meta{Title: "用户商城", Icon: ""}},
		{Hidden: false, ParentId: 28, Path: "search", Name: "search", Component: "/setting/search/search", Sort: 3, Meta: model.Meta{Title: "网站搜索", Icon: ""}},
		{Hidden: false, ParentId: 28, Path: "homeMod", Name: "homeMod", Component: "/setting/homeMod/homeMod", Sort: 4, Meta: model.Meta{Title: "首页模块", Icon: ""}},
		{Hidden: false, ParentId: 28, Path: "footer", Name: "footer", Component: "/setting/footer/footer", Sort: 5, Meta: model.Meta{Title: "底部内容", Icon: ""}},
		{Hidden: false, ParentId: 0, Path: "/log", Name: "log", Component: "LAYOUT", Sort: 8, Meta: model.Meta{Title: "登陆管理", Icon: "LogInOutline"}},
		{Hidden: false, ParentId: 0, Path: "/style", Name: "style", Component: "LAYOUT", Sort: 9, Meta: model.Meta{Title: "主题外观", Icon: "ColorPaletteOutline"}},
		{Hidden: false, ParentId: 0, Path: "/extension", Name: "extension", Component: "LAYOUT", Sort: 10, Meta: model.Meta{Title: "扩展插件", Icon: "ExtensionPuzzleOutline"}},
		{Hidden: false, ParentId: 38, Path: "plugin", Name: "plugin", Component: "/extension/plugin/plugin", Sort: 0, Meta: model.Meta{Title: "我的插件", Icon: ""}},
		{Hidden: false, ParentId: 38, Path: "pluginShop", Name: "pluginShop", Component: "/extension/pluginShop/pluginShop", Sort: 1, Meta: model.Meta{Title: "插件商店", Icon: ""}},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return errors.Wrap(err, m.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (m *menu) CheckDataExist() bool {
	if errors.Is(global.XD_DB.Where("path = ?", "/system").First(&model.XdBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
