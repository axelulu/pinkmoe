/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:51:41
 * @FilePath: /pinkmoe_server/util/upload/obs.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */

package upload

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/pkg/errors"
	"mime/multipart"
	"server/model/response"
)

var HuaWeiObs = new(_obs)

type _obs struct{}

func NewHuaWeiObsClient(uploadConfig response.UploadConfig) (client *obs.ObsClient, err error) {
	return obs.New(uploadConfig.HuaWeiObsConfig.AccessKey, uploadConfig.HuaWeiObsConfig.SecretKey, uploadConfig.HuaWeiObsConfig.Endpoint)
}

func (o *_obs) UploadFile(file *multipart.FileHeader, uploadConfig response.UploadConfig) (filename string, filepath string, err error) {
	var open multipart.File
	open, err = file.Open()
	if err != nil {
		return filename, filepath, err
	}
	filename = file.Filename
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: uploadConfig.HuaWeiObsConfig.Bucket,
				Key:    filename,
			},
			ContentType: file.Header.Get("content-type"),
		},
		Body: open,
	}

	var client *obs.ObsClient
	client, err = NewHuaWeiObsClient(uploadConfig)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "获取华为对象存储对象失败!")
	}

	_, err = client.PutObject(input)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "文件上传失败!")
	}
	filepath = uploadConfig.HuaWeiObsConfig.Path + "/" + filename
	return filepath, filename, err
}

func (o *_obs) DeleteFile(key string, uploadConfig response.UploadConfig) error {
	client, err := NewHuaWeiObsClient(uploadConfig)
	if err != nil {
		return errors.Wrap(err, "获取华为对象存储对象失败!")
	}
	input := &obs.DeleteObjectInput{
		Bucket: uploadConfig.HuaWeiObsConfig.Bucket,
		Key:    key,
	}
	var output *obs.DeleteObjectOutput
	output, err = client.DeleteObject(input)
	if err != nil {
		return errors.Wrapf(err, "删除对象(%s)失败!, output: %v", key, output)
	}
	return nil
}
