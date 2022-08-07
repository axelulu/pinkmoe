/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:53:03
 * @FilePath: /pinkmoe_server/model/request/api.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package request

// SearchApiParams api分页条件查询及排序结构体
type SearchApiParams struct {
	PageInfo
	Path     string `json:"path" form:"path" gorm:"comment:api路径"`               // api路径
	ApiGroup string `json:"apiGroup" form:"apiGroup" gorm:"comment:api组"`        // api组
	Method   string `json:"method" form:"method" gorm:"default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
	OrderKey string `json:"orderKey"`                                            // 排序
	Desc     bool   `json:"desc"`                                                // 排序方式:升序false(默认)|降序true
}

type UpdateApiAuthority struct {
	AuthorityId string       `json:"authorityId"` // 权限id
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}
