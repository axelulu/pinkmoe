/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:07
 * @FilePath: /pinkmoe_server/initialize/gorm.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package initialize

import (
	"os"
	"server/global"
	"server/initialize/initDB"
	"server/model"

	adapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitMysqlData() error {
	return model.MysqlDataInitialize(
		initDB.Api,
		initDB.BaseMenu,
		initDB.Authority,
		initDB.AuthorityParams,
		initDB.AuthoritiesMenus,
		initDB.User,
		initDB.Category,
		initDB.Post,
		initDB.CommentSmile,
		initDB.UserAuthority,
		initDB.Setting,
		initDB.Casbin,
		initDB.ViewAuthorityMenuMysql,
	)
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		model.XdApi{},
		model.XdUser{},
		model.XdBaseMenu{},
		model.XdJwtBlacklist{},
		model.XdAuthority{},
		model.XdAuthorityParams{},
		model.XdPlugin{},
		model.XdSetting{},
		model.XdLoginLog{},
		model.XdTask{},
		model.XdUserOnline{},
		model.XdFollow{},
		model.XdOperationLog{},
		adapter.CasbinRule{},

		// 业务模块表
		model.XdPost{},
		model.XdPostDownload{},
		model.XdPostMusic{},
		model.XdPostVideo{},
		model.XdPostCollect{},
		model.XdPostStar{},
		model.XdGoods{},
		model.XdGoodsCategory{},
		model.XdGoodsSize{},
		model.XdGoodsSizeVal{},
		model.XdGoodsStar{},
		model.XdChat{},
		model.XdChatRelation{},
		model.XdAddress{},
		model.XdOrder{},
		model.XdKey{},
		model.XdUploadFile{},
		model.XdCategory{},
		model.XdComment{},
		model.XdCommentSmile{},
		model.XdTopic{},
		model.XdCheckIn{},
		model.XdNotification{},
	)
	if err != nil {
		global.XD_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.XD_LOG.Info("register table success")
}
