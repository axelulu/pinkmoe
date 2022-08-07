/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:51:41
 * @FilePath: /pinkmoe_server/util/upload/obs.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */

package upload

import (
	"mime/multipart"
	"server/global"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/pkg/errors"
)

var HuaWeiObs = new(_obs)

type _obs struct{}

func NewHuaWeiObsClient() (client *obs.ObsClient, err error) {
	return obs.New(global.XD_CONFIG.HuaWeiObsConfig.AccessKey, global.XD_CONFIG.HuaWeiObsConfig.SecretKey, global.XD_CONFIG.HuaWeiObsConfig.Endpoint)
}

func (o *_obs) UploadFile(file *multipart.FileHeader) (filename string, filepath string, err error) {
	var open multipart.File
	open, err = file.Open()
	if err != nil {
		return filename, filepath, err
	}
	filename = file.Filename
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: global.XD_CONFIG.HuaWeiObsConfig.Bucket,
				Key:    filename,
			},
			ContentType: file.Header.Get("content-type"),
		},
		Body: open,
	}

	var client *obs.ObsClient
	client, err = NewHuaWeiObsClient()
	if err != nil {
		return filepath, filename, errors.Wrap(err, "获取华为对象存储对象失败!")
	}

	_, err = client.PutObject(input)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "文件上传失败!")
	}
	filepath = global.XD_CONFIG.HuaWeiObsConfig.Path + "/" + filename
	return filepath, filename, err
}

func (o *_obs) DeleteFile(key string) error {
	client, err := NewHuaWeiObsClient()
	if err != nil {
		return errors.Wrap(err, "获取华为对象存储对象失败!")
	}
	input := &obs.DeleteObjectInput{
		Bucket: global.XD_CONFIG.HuaWeiObsConfig.Bucket,
		Key:    key,
	}
	var output *obs.DeleteObjectOutput
	output, err = client.DeleteObject(input)
	if err != nil {
		return errors.Wrapf(err, "删除对象(%s)失败!, output: %v", key, output)
	}
	return nil
}
