/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:01
 * @FilePath: /pinkmoe_server/model/initdb.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package model

import "github.com/gookit/color"

const (
	Mysql           = "mysql"
	Pgsql           = "pgsql"
	InitSuccess     = "\n[%v] --> 初始数据成功!\n"
	AuthorityMenu   = "\n[%v] --> %v 视图已存在!\n"
	InitDataExist   = "\n[%v] --> %v 表的初始数据已存在!\n"
	InitDataFailed  = "\n[%v] --> %v 表初始数据失败! \nerr: %+v\n"
	InitDataSuccess = "\n[%v] --> %v 表初始数据成功!\n"
)

type InitData interface {
	TableName() string
	Initialize() error
	CheckDataExist() bool
}

// MysqlDataInitialize Mysql 初始化接口使用封装
// Author [SliverHorn](https://github.com/SliverHorn)
func MysqlDataInitialize(inits ...InitData) error {
	var entity XdMenu
	for i := 0; i < len(inits); i++ {
		if inits[i].TableName() == entity.TableName() {
			if k := inits[i].CheckDataExist(); k {
				color.Info.Printf(AuthorityMenu, Mysql, inits[i].TableName())
				continue
			}
		} else {
			if inits[i].CheckDataExist() {
				color.Info.Printf(InitDataExist, Mysql, inits[i].TableName())
				continue
			}
		}

		if err := inits[i].Initialize(); err != nil {
			color.Info.Printf(InitDataFailed, Mysql, err)
			return err
		} else {
			color.Info.Printf(InitDataSuccess, Mysql, inits[i].TableName())
		}
	}
	color.Info.Printf(InitSuccess, Mysql)
	return nil
}
