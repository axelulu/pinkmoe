/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:57:44
 * @FilePath: /pinkmoe_server/middleware/operation.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"server/dao/mysql"
	"server/global"
	"server/model"
	"server/util"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

var respPool sync.Pool

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId uuid.UUID
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.XD_LOG.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		userId, err := util.GetCurrentUserID(c)
		var record model.XdOperationLog
		if err != nil {
			record = model.XdOperationLog{
				Ip:     c.ClientIP(),
				Method: c.Request.Method,
				Path:   c.Request.URL.Path,
				Agent:  c.Request.UserAgent(),
				Body:   string(body),
			}
		} else {
			record = model.XdOperationLog{
				Ip:     c.ClientIP(),
				Method: c.Request.Method,
				Path:   c.Request.URL.Path,
				Agent:  c.Request.UserAgent(),
				Body:   string(body),
				UserID: userId,
			}
		}

		// 上传文件时候 中间件日志进行裁断操作
		//if strings.Index(c.GetHeader("Content-Type"), "multipart/form-data") > -1 {
		//	if len(record.Body) > 2048 {
		//		// 截断
		//		newBody := respPool.Get().([]byte)
		//		copy(newBody, record.Body)
		//		record.Body = string(newBody)
		//		defer respPool.Put(newBody[:0])
		//	}
		//}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if strings.Index(c.Writer.Header().Get("Pragma"), "public") > -1 ||
			strings.Index(c.Writer.Header().Get("Expires"), "0") > -1 ||
			strings.Index(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/force-download") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/octet-stream") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/download") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Disposition"), "attachment") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") > -1 {
			if len(record.Resp) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Resp)
				record.Body = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		if err := mysql.CreateOperation(&record); err != nil {
			global.XD_LOG.Error("create operation record error:", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
