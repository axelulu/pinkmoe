/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 07:06:19
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:55:11
 * @FilePath: /pinkmoe_server/model/xd_authority_params.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package model

type XdAuthorityParams struct {
	AuthorityId         string `json:"authorityId" gorm:"comment:角色ID"`          // 角色ID
	AuthorityParamId    string `json:"authorityParamId" gorm:"comment:角色扩展ID"`   // 角色扩展ID
	AuthorityParamKey   string `json:"authorityParamKey" gorm:"comment:角色扩展名"`   // 角色扩展名
	AuthorityParamValue int    `json:"authorityParamValue" gorm:"comment:角色扩展值"` // 角色扩展值
	AuthorityParamDay   int    `json:"authorityParamDay" gorm:"comment:角色扩展值2"`  // 角色扩展值2
}
