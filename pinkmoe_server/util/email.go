/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:51:29
 * @FilePath: /pinkmoe_server/util/email.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */

package util

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"
	"server/dao/redis"
	"server/global"
	"server/model"
	"server/model/response"
	"strings"
	"time"

	"github.com/jordan-wright/email"
	"gorm.io/gorm"
)

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func getCaptchaCode(emailType string, emailCode string) (code string) {
	code = GenValidateCode(6)
	global.XD_REDIS.Set(context.Background(), redis.EmailPreKey+emailType+"_"+emailCode, code, redis.EmailExpiration)
	return
}

func getEmailTitle(emailType string) (title string, err error) {
	switch emailType {
	case "reg":
		return "网站注册验证码", nil
	case "forget":
		return "忘记密码验证码", nil
	case "changePwd":
		return "修改密码验证码", nil
	default:
		return "", response.ErrorEmailTitle
	}
}

func getEmailContent(emailType string, emailCode string) (title string, err error) {
	code := getCaptchaCode(emailType, emailCode)
	switch emailType {
	case "reg":
		return "您的网站注册验证码为:" + code, nil
	case "forget":
		return "您的忘记密码验证码为:" + code, nil
	case "changePwd":
		return "您的修改密码验证码为:" + code, nil
	default:
		return "", response.ErrorEmailContent
	}
}

func SendEmailCaptcha(emailCode string, emailType string) (err error) {
	var user model.XdUser
	if emailType == "reg" {
		if !errors.Is(global.XD_DB.Where("email = ?", emailCode).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
			return response.ErrorEmailExist
		}
	}
	title, err := getEmailTitle(emailType)
	if err != nil {
		return err
	}
	content, err := getEmailContent(emailType, emailCode)
	if err != nil {
		return err
	}
	err = SendEmail(emailCode, title, content)
	return
}

func SendEmail(emailCode string, title string, content string) (err error) {
	from := global.XD_CONFIG.EmailConfig.User
	nickname := global.XD_CONFIG.EmailConfig.Username
	secret := global.XD_CONFIG.EmailConfig.Password
	host := global.XD_CONFIG.EmailConfig.Host
	port := global.XD_CONFIG.EmailConfig.Port
	isSSL := global.XD_CONFIG.EmailConfig.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = []string{emailCode}
	e.Subject = title
	e.HTML = []byte(content)
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	if err != nil {
		return response.ErrorEmailSend
	}
	return err
}
