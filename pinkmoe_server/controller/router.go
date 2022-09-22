/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:39
 * @FilePath: /pinkmoe_server/controller/router.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package controller

import (
	"net/http"
	"reflect"
	"server/controller/socket"
	"server/global"
	"server/middleware"
	"server/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

// Route 路由结构体
type Route struct {
	Path       string         //url路径
	HttpMethod string         //http方法 get post
	Method     reflect.Value  //方法路由
	Args       []reflect.Type //参数类型
	Auth       bool           //是否使用授权
}

// Routes 路由集合
var Routes []Route

func InitRouter() *gin.Engine {
	mode := global.XD_CONFIG.BasicConfig.Mode
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	//初始化路由
	r := gin.New()
	r.Use(middleware.Cors())

	//socket
	r.GET("/ws", socket.InitSocket)

	r.StaticFS(global.XD_CONFIG.UploadConfig.LocalConfig.Path, http.Dir(global.XD_CONFIG.UploadConfig.LocalConfig.Path)) // 为用户头像和文件提供静态地址
	r.Use(util.GinLogger(), util.GinRecovery(true), middleware.RateLimitMiddleware(time.Duration(global.XD_CONFIG.BasicConfig.RateLimitTime)*time.Second, global.XD_CONFIG.BasicConfig.RateLimitNum))

	// swagger文档地址
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	Bind(v1)
	v1.Use(middleware.JWTAuthMiddleware())
	v1.Use(middleware.OperationRecord())
	BindAuth(v1)

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{
			"msg": "OK",
		})
	})

	// pprof性能测试
	//pprof.Register(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"msg": "404",
		})
	})
	return r
}

// Register 注册控制器
func Register(controller interface{}) bool {
	ctrlName := reflect.TypeOf(controller).String()
	module := ctrlName
	controllerStr := ctrlName
	if strings.Contains(ctrlName, ".") {
		module = ctrlName[strings.Index(ctrlName, ".")+1:]
		cstr := ctrlName[:strings.Index(ctrlName, ".")][1:]
		controllerStr = strings.ToUpper(cstr[:1]) + cstr[1:]
	}
	v := reflect.ValueOf(controller)
	//遍历方法
	for i := 0; i < v.NumMethod(); i++ {
		isAuth := false
		method := v.Method(i)
		//遍历参数
		params := make([]reflect.Type, 0, v.NumMethod())
		httpMethod := "POST"
		action := v.Type().Method(i).Name
		if len(action) >= 4 && action[:4] == "Auth" {
			isAuth = true
			action = action[4:]
		}
		if action[len(action)-3:] == "Get" {
			httpMethod = "GET"
			action = action[:len(action)-3]
		}
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
		}
		path := "/" + controllerStr + "/" + module + "/" + action
		route := Route{Path: path, Method: method, Args: params, HttpMethod: httpMethod, Auth: isAuth}
		Routes = append(Routes, route)
	}
	return true
}

// Bind 绑定路由 m是方法GET POST等
//绑定基本路由
func Bind(e *gin.RouterGroup) {
	for _, route := range Routes {
		if route.HttpMethod == "GET" && !route.Auth {
			e.GET(route.Path, match(route.Path, route))
		}
		if route.HttpMethod == "POST" && !route.Auth {
			e.POST(route.Path, match(route.Path, route))
		}
	}
}

func BindAuth(e *gin.RouterGroup) {
	for _, route := range Routes {
		if route.HttpMethod == "GET" && route.Auth {
			e.GET(route.Path, match(route.Path, route))
		}
		if route.HttpMethod == "POST" && route.Auth {
			e.POST(route.Path, match(route.Path, route))
		}
	}
}

//根据path匹配对应的方法
func match(path string, route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")
		if len(fields) < 3 {
			return
		}
		if len(Routes) > 0 {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c) // *gin.Context
			route.Method.Call(arguments)
		}
	}
}
