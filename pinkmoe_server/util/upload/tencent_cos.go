/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:51:49
 * @FilePath: /pinkmoe_server/util/upload/tencent_cos.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
	"server/model/response"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
)

type TencentCOS struct{}

// UploadFile upload file to COS
func (*TencentCOS) UploadFile(file *multipart.FileHeader, uploadConfig response.UploadConfig) (string, string, error) {
	client := NewClient(uploadConfig)
	f, openError := file.Open()
	if openError != nil {
		global.XD_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), uploadConfig.TencentCOSConfig.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return uploadConfig.TencentCOSConfig.BaseURL + "/" + uploadConfig.TencentCOSConfig.PathPrefix + "/" + fileKey, fileKey, nil
}

// DeleteFile delete file form COS
func (*TencentCOS) DeleteFile(key string, uploadConfig response.UploadConfig) error {
	client := NewClient(uploadConfig)
	name := uploadConfig.TencentCOSConfig.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		global.XD_LOG.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

// NewClient init COS client
func NewClient(uploadConfig response.UploadConfig) *cos.Client {
	urlStr, _ := url.Parse("https://" + uploadConfig.TencentCOSConfig.Bucket + ".cos." + uploadConfig.TencentCOSConfig.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  uploadConfig.TencentCOSConfig.SecretID,
			SecretKey: uploadConfig.TencentCOSConfig.SecretKey,
		},
	})
	return client
}
