/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:52:54
 * @FilePath: /pinkmoe_server/plugin/demo/controller/demo.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package demo

import (
	"server/controller"
	"server/model/response"

	"github.com/gin-gonic/gin"
)

//这里每个controller执行init方法都要注册自动路由
func init() {
	controller.Register(&Demo{})
}

type Demo struct{}

// DemoGet 控制器的方法 分页查询
func (api *Demo) DemoGet(c *gin.Context) {
	response.OkWithData("ok ", c)
}
