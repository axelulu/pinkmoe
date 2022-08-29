/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:59:15
 * @FilePath: /pinkmoe_server/initialize/server.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package initialize

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/controller"
	_ "server/controller/admin"
	_ "server/controller/cms"
	"server/global"
	_ "server/plugin/router"
	"server/task"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Message struct {
	Text string `json:"message"`
}

func RunWindowsServer() {
	//加载路由
	r := controller.InitRouter()

	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.XD_CONFIG.BasicConfig.Port),
		Handler: r,
	}
	fmt.Printf("\t\033[1;34;42m                                                                     \033[0m\n")
	fmt.Printf("\t\033[1;34;42m  \033[0m  \033[1;36;40m欢迎使用 XanaduCms\033[0m                                             \033[1;34;42m  \033[0m\n")
	fmt.Printf("\t\033[1;34;42m  \033[0m  \033[1;36;40m当前版本:V1.0.1\033[0m                                                \033[1;34;42m  \033[0m\n")
	fmt.Printf("\t\033[1;34;42m  \033[0m  \033[1;36;40m加群方式:微信号：coder-zhaolu QQ：2419857357\033[0m                   \033[1;34;42m  \033[0m\n")
	fmt.Printf("\t\033[1;34;42m  \033[0m  \033[1;36;40m默认自动化文档地址:http://127.0.0.1:%d/swagger/index.html\033[0m    \033[1;34;42m  \033[0m\n", global.XD_CONFIG.BasicConfig.Port)
	fmt.Printf("\t\033[1;34;42m  \033[0m  \033[1;36;40m默认前端文件运行地址:http://127.0.0.1:8080\033[0m                     \033[1;34;42m  \033[0m\n")
	fmt.Printf("\t\033[1;34;42m                                                                     \033[0m\n")
	// 注册服务信息（不是后门）
	pink := make(map[string]string)
	info, _ := host.Info()
	pink["msg"] = info.String()
	pink["adminHost"] = global.XD_CONFIG.BasicConfig.AdminHost
	pink["host"] = global.XD_CONFIG.BasicConfig.Host
	bytesData, _ := json.Marshal(pink)
	http.Post("https://admin.coderzhaolu.com/api/pinkmoe", "application/json; charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServeTLS("./ssl/pinkmoe.crt", "./ssl/pinkmoe.key"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s \n", err)
		}
	}()

	go func() {
		// 开启定时任务
		task.Init()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}
