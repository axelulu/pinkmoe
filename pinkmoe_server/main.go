/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:51:57
 * @FilePath: /pinkmoe_server/main.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */

package main

import (
	"fmt"
	"os"
	"server/global"
	"server/initialize"

	"go.uber.org/zap"
)

// @title 这里写标题
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path
// GO Web开发比较通用的脚手架模版
func main() {
	if len(os.Args) < 2 {
		fmt.Print("请指定config文件！")
		return
	}

	// 1. 加载配置
	global.XD_VIPER = initialize.ViperInit(os.Args[1])

	// 2. 初始化日志
	global.XD_LOG = initialize.ZapInit()
	defer zap.L().Sync()

	// 3. 初始化Sql连接
	global.XD_DB = initialize.DbInit()
	if global.XD_DB != nil {
		initialize.RegisterTables(global.XD_DB) // 初始化表
		initialize.InitMysqlData()              //初始化数据
		// 程序结束前关闭数据库链接
		db, _ := global.XD_DB.DB()
		defer db.Close()
	}

	// 4. 初始化Redis连接
	global.XD_REDIS = initialize.RedisInit()
	defer global.XD_REDIS.Close()

	// 5. 初始化雪花算法
	global.XD_SNOWFLAKE = initialize.SnowflakeInit(global.XD_CONFIG.StartTime, global.XD_CONFIG.MachineId)

	initialize.RunWindowsServer()
}
