/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-10-04 22:34:39
 * @FilePath: /pinkmoe/pinkmoe_server/main.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */

package main

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"os"
	"server/dao/mysql"
	"server/global"
	"server/initialize"
	"server/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("请指定config文件！")
		return
	}

	// 加载mysql配置
	global.XD_VIPER = initialize.ViperInit(os.Args[1])

	// 使用config文件配置
	if global.XD_DB == nil {
		// 初始化日志
		global.XD_LOG = initialize.ZapInit()
		defer zap.L().Sync()
	}

	// 初始化Sql连接
	global.XD_DB = initialize.DbInit()

	// 使用数据库配置
	if global.XD_DB != nil {
		// 初始化日志
		global.XD_LOG = initialize.ZapInit()
		defer zap.L().Sync()
	}

	if global.XD_DB != nil {
		initialize.RegisterTables(global.XD_DB) // 初始化表
		initialize.InitMysqlData()              //初始化数据
		// 程序结束前关闭数据库链接
		db, _ := global.XD_DB.DB()
		defer db.Close()
	}

	// 初始化数据库config配置项
	_, site := mysql.GetSettingItem(model.XdSetting{Slug: "system_site"})
	json.Unmarshal([]byte(site.Value), &global.XD_CONFIG.BasicConfig)

	// 初始化Redis连接
	global.XD_REDIS = initialize.RedisInit()
	defer global.XD_REDIS.Close()

	// 初始化雪花算法
	global.XD_SNOWFLAKE = initialize.SnowflakeInit(global.XD_CONFIG.StartTime, global.XD_CONFIG.MachineId)

	initialize.RunWindowsServer()
}
