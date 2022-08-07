/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:00:23
 * @FilePath: /pinkmoe_server/global/global.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package global

import (
	"server/config"

	sf "github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	XD_CONFIG    *config.AppConfig
	XD_VIPER     *viper.Viper
	XD_LOG       *zap.Logger
	XD_REDIS     *redis.Client
	XD_SNOWFLAKE *sf.Node
	XD_DB        *gorm.DB
)
