/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:51:21
 * @FilePath: /pinkmoe_server/util/jwt.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */

package util

import (
	"server/global"
	"server/model/request"
	"server/model/response"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID      uuid.UUID `json:"userId"`
	Username    string    `json:"username"`
	AuthorityId string    `json:"authorityId"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(authorityId string, userID uuid.UUID, username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userID,
		username, // 自定义字段
		authorityId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(global.XD_CONFIG.BasicConfig.AuthConfig.JwtExpire) * time.Hour).Unix(), // 过期时间
			Issuer:    global.XD_CONFIG.BasicConfig.Issuer,                                                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(global.XD_CONFIG.BasicConfig.AuthConfig.AuthSecret))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(global.XD_CONFIG.BasicConfig.AuthConfig.AuthSecret), nil
	})
	if err != nil {
		return nil, response.ErrorJwtToken
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, response.ErrorJwtToken
}

// GetCurrentUserID 获取当前用户
func GetCurrentUserID(c *gin.Context) (userID uuid.UUID, err error) {
	uid := c.GetString(request.CtxUserIDKey)
	return uuid.FromString(uid)
}

// GetCurrentUserToken 获取当前用户
func GetCurrentUserToken(c *gin.Context) (token string) {
	uid := c.GetString(request.CtxUserTokenKey)
	return uid
}
