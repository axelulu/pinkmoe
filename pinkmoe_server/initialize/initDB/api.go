/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:42
 * @FilePath: /pinkmoe_server/initialize/initDB/api.go
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

var Api = new(api)

type api struct{}

func (a *api) TableName() string {
	return "xd_apis"
}

func (a *api) Initialize() error {
	entities := []model.XdApi{
		{ApiGroup: "仪表盘", Method: "GET", Path: "/api/v1/Admin/Dashboard/Console", Description: "获取控制台信息"},

		{ApiGroup: "系统用户", Method: "GET", Path: "/api/v1/Admin/User/User", Description: "获取当前用户"},
		{ApiGroup: "系统用户", Method: "GET", Path: "/api/v1/Admin/User/UserList", Description: "获取用户列表"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/api/v1/Admin/User/Logout", Description: "退出登陆"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/api/v1/Admin/User/UserCreate", Description: "创建用户"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/api/v1/Admin/User/UserUpdate", Description: "更新用户"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/api/v1/Admin/User/UserDelete", Description: "删除用户"},

		{ApiGroup: "api", Method: "GET", Path: "/api/v1/Admin/Api/AllApi", Description: "获取所有api"},
		{ApiGroup: "api", Method: "GET", Path: "/api/v1/Admin/Api/ApiList", Description: "获取api列表"},
		{ApiGroup: "api", Method: "GET", Path: "/api/v1/Admin/Api/ApiAuthority", Description: "获取角色api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/v1/Admin/Api/ApiAuthorityUpdate", Description: "更新角色api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/v1/Admin/Api/ApiCreate", Description: "创建api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/v1/Admin/Api/ApiUpdate", Description: "更新api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/v1/Admin/Api/ApiDelete", Description: "删除api"},

		{ApiGroup: "角色", Method: "GET", Path: "/api/v1/Admin/Authority/AuthorityList", Description: "获取角色列表"},
		{ApiGroup: "角色", Method: "GET", Path: "/api/v1/Admin/Authority/AuthorityMenu", Description: "获取角色列表"},
		{ApiGroup: "角色", Method: "POST", Path: "/api/v1/Admin/Authority/AuthorityCreate", Description: "创建角色"},
		{ApiGroup: "角色", Method: "POST", Path: "/api/v1/Admin/Authority/AuthorityUpdate", Description: "更新角色"},
		{ApiGroup: "角色", Method: "POST", Path: "/api/v1/Admin/Authority/AuthorityMenuUpdate", Description: "更新角色菜单"},
		{ApiGroup: "角色", Method: "POST", Path: "/api/v1/Admin/Authority/AuthorityDelete", Description: "删除角色"},

		{ApiGroup: "菜单", Method: "GET", Path: "/api/v1/Admin/Menu/Menu", Description: "获取当前用户菜单"},
		{ApiGroup: "菜单", Method: "GET", Path: "/api/v1/Admin/Menu/AllMenu", Description: "获取所有菜单"},
		{ApiGroup: "菜单", Method: "GET", Path: "/api/v1/Admin/Menu/MenuList", Description: "获取菜单树"},
		{ApiGroup: "菜单", Method: "POST", Path: "/api/v1/Admin/Menu/MenuCreate", Description: "创建菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/api/v1/Admin/Menu/MenuUpdate", Description: "更新菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/api/v1/Admin/Menu/MenuDelete", Description: "删除菜单"},

		{ApiGroup: "操作日志", Method: "GET", Path: "/api/v1/Admin/Operation/OperationList", Description: "获取操作日志"},
		{ApiGroup: "操作日志", Method: "POST", Path: "/api/v1/Admin/Operation/OperationCreate", Description: "创建操作日志"},
		{ApiGroup: "操作日志", Method: "POST", Path: "/api/v1/Admin/Operation/OperationUpdate", Description: "更新操作日志"},
		{ApiGroup: "操作日志", Method: "POST", Path: "/api/v1/Admin/Operation/OperationDelete", Description: "删除操作日志"},

		{ApiGroup: "在线用户", Method: "GET", Path: "/api/v1/Admin/UserOnline/UserOnlineList", Description: "获取在线用户"},
		{ApiGroup: "在线用户", Method: "POST", Path: "/api/v1/Admin/UserOnline/UserOnlineCreate", Description: "创建在线用户"},
		{ApiGroup: "在线用户", Method: "POST", Path: "/api/v1/Admin/UserOnline/UserOnlineUpdate", Description: "更新在线用户"},
		{ApiGroup: "在线用户", Method: "POST", Path: "/api/v1/Admin/UserOnline/UserOnlineDelete", Description: "删除在线用户"},

		{ApiGroup: "登陆日志", Method: "GET", Path: "/api/v1/Admin/LoginLog/LoginLogList", Description: "获取登陆日志"},
		{ApiGroup: "登陆日志", Method: "POST", Path: "/api/v1/Admin/LoginLog/LoginLogCreate", Description: "创建登陆日志"},
		{ApiGroup: "登陆日志", Method: "POST", Path: "/api/v1/Admin/LoginLog/LoginLogUpdate", Description: "更新登陆日志"},
		{ApiGroup: "登陆日志", Method: "POST", Path: "/api/v1/Admin/LoginLog/LoginLogDelete", Description: "删除登陆日志"},

		{ApiGroup: "文章", Method: "GET", Path: "/api/v1/Admin/Post/PostList", Description: "获取文章列表"},
		{ApiGroup: "文章", Method: "POST", Path: "/api/v1/Admin/Post/PostCreate", Description: "创建文章"},
		{ApiGroup: "文章", Method: "POST", Path: "/api/v1/Admin/Post/PostUpdate", Description: "更新文章"},
		{ApiGroup: "文章", Method: "POST", Path: "/api/v1/Admin/Post/PostUpdateStatus", Description: "更新文章状态"},
		{ApiGroup: "文章", Method: "POST", Path: "/api/v1/Admin/Post/PostDelete", Description: "删除文章"},

		{ApiGroup: "分类", Method: "GET", Path: "/api/v1/Admin/Category/CategoryList", Description: "获取分类列表"},
		{ApiGroup: "分类", Method: "POST", Path: "/api/v1/Admin/Category/CategoryCreate", Description: "创建分类"},
		{ApiGroup: "分类", Method: "POST", Path: "/api/v1/Admin/Category/CategoryUpdate", Description: "更新分类"},
		{ApiGroup: "分类", Method: "POST", Path: "/api/v1/Admin/Category/CategoryDelete", Description: "删除分类"},

		{ApiGroup: "话题", Method: "GET", Path: "/api/v1/Admin/Topic/TopicList", Description: "获取话题列表"},
		{ApiGroup: "话题", Method: "POST", Path: "/api/v1/Admin/Topic/TopicCreate", Description: "创建话题"},
		{ApiGroup: "话题", Method: "POST", Path: "/api/v1/Admin/Topic/TopicUpdate", Description: "更新话题"},
		{ApiGroup: "话题", Method: "POST", Path: "/api/v1/Admin/Topic/TopicDelete", Description: "删除话题"},

		{ApiGroup: "卡密", Method: "GET", Path: "/api/v1/Admin/Key/KeyList", Description: "获取卡密列表"},
		{ApiGroup: "卡密", Method: "POST", Path: "/api/v1/Admin/Key/KeyCreate", Description: "创建卡密"},
		{ApiGroup: "卡密", Method: "POST", Path: "/api/v1/Admin/Key/KeyDelete", Description: "删除卡密"},

		{ApiGroup: "评论", Method: "GET", Path: "/api/v1/Admin/Comment/CommentList", Description: "获取评论列表"},
		{ApiGroup: "评论", Method: "POST", Path: "/api/v1/Admin/Comment/CommentCreate", Description: "创建评论"},
		{ApiGroup: "评论", Method: "POST", Path: "/api/v1/Admin/Comment/CommentUpdate", Description: "更新评论"},
		{ApiGroup: "评论", Method: "POST", Path: "/api/v1/Admin/Comment/CommentDelete", Description: "删除评论"},

		{ApiGroup: "设置", Method: "GET", Path: "/api/v1/Admin/Setting/SettingItem", Description: "获取菜单"},
		{ApiGroup: "设置", Method: "POST", Path: "/api/v1/Admin/Setting/SettingCreate", Description: "创建菜单"},
		{ApiGroup: "设置", Method: "POST", Path: "/api/v1/Admin/Setting/SettingUpdate", Description: "更新菜单"},

		{ApiGroup: "文件", Method: "POST", Path: "/api/v1/Admin/Upload/FileUpload", Description: "文件上传"},
		{ApiGroup: "文件", Method: "POST", Path: "/api/v1/Admin/Upload/ImgFileDelete", Description: "图片文件删除"},

		// cms权限
		{ApiGroup: "CMS用户", Method: "POST", Path: "/api/v1/Cms/User/Logout", Description: "用户登出"},
		{ApiGroup: "CMS用户", Method: "GET", Path: "/api/v1/Cms/User/User", Description: "CMS用户信息获取"},
		{ApiGroup: "CMS用户", Method: "POST", Path: "/api/v1/Cms/User/UserEmailUpdate", Description: "CMS用户邮箱修改"},
		{ApiGroup: "CMS用户", Method: "POST", Path: "/api/v1/Cms/User/UserDetailUpdate", Description: "CMS用户信息修改"},
		{ApiGroup: "CMS用户", Method: "POST", Path: "/api/v1/Cms/User/UserAvatarUpdate", Description: "CMS用户头像修改"},
		{ApiGroup: "CMS用户", Method: "POST", Path: "/api/v1/Cms/User/UserPwdUpdate", Description: "CMS用户密码修改"},
		{ApiGroup: "CMS用户", Method: "POST", Path: "/api/v1/Cms/User/UserCheckIn", Description: "CMS用户签到"},
		{ApiGroup: "CMS用户", Method: "POST", Path: "/api/v1/Cms/User/UserCheckInStatus", Description: "CMS用户签到状态"},
		{ApiGroup: "CMS用户", Method: "POST", Path: "/api/v1/Cms/User/UserBuyVip", Description: "CMS用户购买会员"},
		{ApiGroup: "CMS用户", Method: "POST", Path: "/api/v1/Cms/User/UserBuyShop", Description: "CMS用户商城购买"},
		{ApiGroup: "CMS用户", Method: "GET", Path: "/api/v1/Cms/Authority/AuthorityList", Description: "CMS用户角色列表"},

		{ApiGroup: "CMS用户关注", Method: "POST", Path: "/api/v1/Cms/Follow/CreateFollow", Description: "CMS用户关注"},
		{ApiGroup: "CMS用户关注", Method: "POST", Path: "/api/v1/Cms/Follow/DeleteFollow", Description: "CMS用户取消关注"},
		{ApiGroup: "CMS用户关注", Method: "GET", Path: "/api/v1/Cms/Follow/FollowStatus", Description: "CMS用户关注状态"},

		{ApiGroup: "CMS文件", Method: "GET", Path: "/api/v1/Cms/Upload/FileList", Description: "文件列表获取"},
		{ApiGroup: "CMS文件", Method: "POST", Path: "/api/v1/Cms/Upload/FileUpload", Description: "文件上传"},
		{ApiGroup: "CMS文件", Method: "POST", Path: "/api/v1/Cms/Upload/ImgFileDelete", Description: "图片文件删除"},

		{ApiGroup: "CMS文章", Method: "GET", Path: "/api/v1/Cms/Post/PostItemDownload", Description: "CMS文章下载获取"},
		{ApiGroup: "CMS文章", Method: "POST", Path: "/api/v1/Cms/Post/PostItemDownloadBuy", Description: "CMS文章下载购买"},
		{ApiGroup: "CMS文章", Method: "GET", Path: "/api/v1/Cms/Post/PostItemVideo", Description: "CMS文章视频获取"},
		{ApiGroup: "CMS文章", Method: "POST", Path: "/api/v1/Cms/Post/PostItemVideoBuy", Description: "CMS文章视频购买"},
		{ApiGroup: "CMS文章", Method: "GET", Path: "/api/v1/Cms/Post/PostItemMusic", Description: "CMS文章音乐获取"},
		{ApiGroup: "CMS文章", Method: "POST", Path: "/api/v1/Cms/Post/PostItemMusicBuy", Description: "CMS文章音乐购买"},
		{ApiGroup: "CMS文章", Method: "POST", Path: "/api/v1/Cms/Post/PostCreate", Description: "创建文章"},
		{ApiGroup: "CMS文章", Method: "GET", Path: "/api/v1/Cms/Post/UserPostList", Description: "用户文章列表"},
		{ApiGroup: "CMS文章", Method: "POST", Path: "/api/v1/Cms/Post/PostCollect", Description: "用户文章收藏"},
		{ApiGroup: "CMS文章", Method: "POST", Path: "/api/v1/Cms/Post/PostUnCollect", Description: "用户文章取消收藏"},
		{ApiGroup: "CMS文章", Method: "GET", Path: "/api/v1/Cms/Post/CollectPost", Description: "用户收藏文章"},

		{ApiGroup: "CMS私聊", Method: "GET", Path: "/api/v1/Cms/Chat/ChatList", Description: "私聊列表"},
		{ApiGroup: "CMS私聊", Method: "GET", Path: "/api/v1/Cms/Chat/ChatRelationList", Description: "私聊关系列表"},
		{ApiGroup: "CMS私聊", Method: "POST", Path: "/api/v1/Cms/Chat/CreateChat", Description: "创建私聊"},
		{ApiGroup: "CMS私聊", Method: "POST", Path: "/api/v1/Cms/Chat/CreateChatRelation", Description: "创建私聊关系"},
		{ApiGroup: "CMS私聊", Method: "POST", Path: "/api/v1/Cms/Chat/DeleteChatRelation", Description: "删除私聊关系"},

		{ApiGroup: "CMS评论创建", Method: "POST", Path: "/api/v1/Cms/Comment/CommentCreate", Description: "创建评论"},

		{ApiGroup: "CMS通知", Method: "GET", Path: "/api/v1/Cms/Notification/NotificationList", Description: "获取通知"},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, a.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (a *api) CheckDataExist() bool {
	if errors.Is(global.XD_DB.Where("path = ? AND method = ?", "/api/v1/Admin/Upload/FileUpload", "POST").First(&model.XdApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
