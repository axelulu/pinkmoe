/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:11
 * @FilePath: /pinkmoe_server/initialize/mysql.go
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

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbInit() *gorm.DB {
	m := global.XD_CONFIG.MySqlConfig
	if m.Dbname == "" {
		os.Exit(0)
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config, // DSN data source name
		DefaultStringSize:         191,                                                                                            // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                                                                                          // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
