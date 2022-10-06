/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:04
 * @FilePath: /pinkmoe_server/model/response/code.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package response

import (
	"fmt"
	"math/rand"
	"time"
)

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeValidateCode
	CodeEmailExist
	CodeCurrentUser
	CodeCommentCreateFail
	CodeUserMeta
	CodePostMeta
	CodeContactExist
	CodeUserChat
	CodeCatEmpty
	CodeUserNotExist
	CodeIsCurrentUser
	CodeInvalidPassword
	CodeUserFollowed
	CodeUserUnFollowed
	CodePostLiked
	CodePostUnLiked
	CodePostStared
	CodePostUnStared
	CodePostCoined
	CodeServerBusy
	CodeUploadFileOver

	CodeNeedLogin
	CodeInvalidToken
	CodeErrorAuth
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:           "success",
	CodeInvalidParam:      "请求参数错误",
	CodeUserExist:         "用户名已存在",
	CodeValidateCode:      "验证码错误",
	CodeEmailExist:        "邮箱已存在",
	CodeCurrentUser:       "获取当前用户ID失败",
	CodeCommentCreateFail: "评论创建失败",
	CodeUserMeta:          "获取用户信息错误",
	CodePostMeta:          "获取文章信息错误",
	CodeContactExist:      "该对话已存在",
	CodeUserChat:          "获取用户聊天列表错误",
	CodeCatEmpty:          "分类为空",
	CodeUserNotExist:      "用户名不存在",
	CodeIsCurrentUser:     "目标用户错误",
	CodeInvalidPassword:   "用户名或密码错误",
	CodeUserFollowed:      "用户已关注",
	CodeUserUnFollowed:    "用户已取消关注",
	CodePostLiked:         "已经喜欢了该文章",
	CodePostUnLiked:       "已经取消喜欢了该文章",
	CodePostStared:        "已经收藏了该文章",
	CodePostUnStared:      "已经取消收藏了该文章",
	CodePostCoined:        "已经投币了该文章",
	CodeServerBusy:        "服务繁忙",
	CodeUploadFileOver:    "伤处文件过大",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",
	CodeErrorAuth:    "权限不足",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

func Captcha(n int) string {
	code := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		code = fmt.Sprintf("%s%d", code, rand.Intn(10))
	}
	return code
}
