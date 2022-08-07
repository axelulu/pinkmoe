/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:37
 * @FilePath: /pinkmoe_server/initialize/viper.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package initialize

import (
	"fmt"
	"server/global"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func ViperInit(filename string) *viper.Viper {
	// 方法一
	v := viper.New()
	v.SetConfigFile(filename) //从指定文件中读取配置文件
	err := v.ReadInConfig()   //读取配置文件
	if err != nil {
		panic(fmt.Errorf("v.ReadInConfig() err: %v \n", err))
	}
	// 把读取到到配置信息反序列化到Conf变量中
	if err := v.Unmarshal(&global.XD_CONFIG); err != nil {
		fmt.Printf("v.Unmarshal failed, err: %v \n", err)
	}
	// 监控配置文件
	v.WatchConfig()
	// 当配置文件修改了
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件修改了！")
		if err := v.Unmarshal(&global.XD_CONFIG); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v \n", err)
		}
	})
	return v
}
