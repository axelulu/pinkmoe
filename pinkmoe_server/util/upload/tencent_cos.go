/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:51:49
 * @FilePath: /pinkmoe_server/util/upload/tencent_cos.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */

package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"server/global"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
)

type TencentCOS struct{}

// UploadFile upload file to COS
func (*TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewClient()
	f, openError := file.Open()
	if openError != nil {
		global.XD_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), global.XD_CONFIG.UploadConfig.TencentCOSConfig.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return global.XD_CONFIG.UploadConfig.TencentCOSConfig.BaseURL + "/" + global.XD_CONFIG.UploadConfig.TencentCOSConfig.PathPrefix + "/" + fileKey, fileKey, nil
}

// DeleteFile delete file form COS
func (*TencentCOS) DeleteFile(key string) error {
	client := NewClient()
	name := global.XD_CONFIG.UploadConfig.TencentCOSConfig.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		global.XD_LOG.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

// NewClient init COS client
func NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + global.XD_CONFIG.UploadConfig.TencentCOSConfig.Bucket + ".cos." + global.XD_CONFIG.UploadConfig.TencentCOSConfig.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.XD_CONFIG.UploadConfig.TencentCOSConfig.SecretID,
			SecretKey: global.XD_CONFIG.UploadConfig.TencentCOSConfig.SecretKey,
		},
	})
	return client
}
