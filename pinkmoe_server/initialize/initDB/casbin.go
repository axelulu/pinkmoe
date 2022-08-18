/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:54
 * @FilePath: /pinkmoe_server/initialize/initDB/casbin.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package initDB

import (
	"server/global"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

func (c *casbin) TableName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (c *casbin) Initialize() error {
	entities := []adapter.CasbinRule{
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Dashboard/Console", V2: "GET"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/User/User", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/User/UserList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/User/Logout", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/User/UserCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/User/UserUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/User/UserDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Api/AllApi", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Api/ApiList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Api/ApiAuthority", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Api/ApiAuthorityUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Api/ApiCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Api/ApiUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Api/ApiDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Authority/AuthorityList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Authority/AuthorityMenu", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Authority/AuthorityCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Authority/AuthorityMenuUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Authority/AuthorityUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Authority/AuthorityDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Menu/Menu", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Menu/AllMenu", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Menu/MenuList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Menu/MenuCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Menu/MenuUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Menu/MenuDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Operation/OperationList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Operation/OperationCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Operation/OperationUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Operation/OperationDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/UserOnline/UserOnlineList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/UserOnline/UserOnlineCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/UserOnline/UserOnlineUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/UserOnline/UserOnlineDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/LoginLog/LoginLogList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/LoginLog/LoginLogCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/LoginLog/LoginLogUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/LoginLog/LoginLogDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Post/PostList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Post/PostCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Post/PostUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Post/PostUpdateStatus", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Post/PostDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Category/CategoryList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Category/CategoryCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Category/CategoryUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Category/CategoryDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Goods/GoodsList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Goods/GoodsCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Goods/GoodsUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Goods/GoodsUpdateStatus", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Goods/GoodsDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/GoodsCategory/GoodsCategoryList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/GoodsCategory/GoodsCategoryCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/GoodsCategory/GoodsCategoryUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/GoodsCategory/GoodsCategoryDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Topic/TopicList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Topic/TopicCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Topic/TopicUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Topic/TopicDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Key/KeyList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Key/KeyCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Key/KeyDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Comment/CommentList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Comment/CommentCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Comment/CommentUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Comment/CommentDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Setting/SettingItem", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Setting/SettingCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Setting/SettingUpdate", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Upload/FileUpload", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Upload/ImgFileDelete", V2: "POST"},

		// cms
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/User/Logout", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/User/UserPwdUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/User/UserEmailUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/User/UserDetailUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/User/UserAvatarUpdate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/User/User", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/User/UserCheckIn", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/User/UserCheckInStatus", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/User/UserBuyVip", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/User/UserBuyShop", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Authority/AuthorityList", V2: "GET"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Follow/CreateFollow", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Follow/FollowStatus", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Follow/DeleteFollow", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Upload/FileList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Upload/FileUpload", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Upload/ImgFileDelete", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Post/PostItemDownload", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Post/PostItemDownloadBuy", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Post/PostItemVideo", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Post/PostItemVideoBuy", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Post/PostItemMusic", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Post/PostItemMusicBuy", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Post/PostCreate", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Post/UserPostList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Post/PostCollect", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Post/PostUnCollect", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Post/CollectPost", V2: "GET"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Chat/ChatList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Chat/ChatRelationList", V2: "GET"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Chat/CreateChat", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Chat/CreateChatRelation", V2: "POST"},
		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Chat/DeleteChatRelation", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Comment/CommentCreate", V2: "POST"},

		{Ptype: "p", V0: "9527", V1: "/api/v1/Cms/Notification/NotificationList", V2: "GET"},

		// 2333
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/User/Logout", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/User/UserPwdUpdate", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/User/UserEmailUpdate", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/User/UserDetailUpdate", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/User/UserAvatarUpdate", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/User/User", V2: "GET"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/User/UserCheckIn", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/User/UserCheckInStatus", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/User/UserBuyVip", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/User/UserBuyShop", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Authority/AuthorityList", V2: "GET"},

		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Follow/CreateFollow", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Follow/FollowStatus", V2: "GET"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Follow/DeleteFollow", V2: "POST"},

		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Upload/FileList", V2: "GET"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Upload/FileUpload", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Upload/ImgFileDelete", V2: "POST"},

		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Post/PostItemDownload", V2: "GET"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Post/PostItemDownloadBuy", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Post/PostItemVideo", V2: "GET"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Post/PostItemVideoBuy", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Post/PostItemMusic", V2: "GET"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Post/PostItemMusicBuy", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Post/PostCreate", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Post/UserPostList", V2: "GET"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Post/PostCollect", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Post/PostUnCollect", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Post/CollectPost", V2: "GET"},

		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Chat/ChatList", V2: "GET"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Chat/ChatRelationList", V2: "GET"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Chat/CreateChat", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Chat/CreateChatRelation", V2: "POST"},
		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Chat/DeleteChatRelation", V2: "POST"},

		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Comment/CommentCreate", V2: "POST"},

		{Ptype: "p", V0: "2333", V1: "/api/v1/Cms/Notification/NotificationList", V2: "GET"},
	}
	if err := global.XD_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, c.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (c *casbin) CheckDataExist() bool {
	if errors.Is(global.XD_DB.Where(adapter.CasbinRule{Ptype: "p", V0: "9527", V1: "/api/v1/Admin/Setting/SettingUpdate", V2: "POST"}).First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
