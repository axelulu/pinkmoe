/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:53
 * @FilePath: /pinkmoe_server/model/response/xd_upload.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package response

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type XdUploadFile struct {
	global.XD_MODEL
	Name   string    `json:"name" gorm:"comment:文件名"`                 // 文件名
	Url    string    `json:"url" gorm:"comment:文件地址"`                 // 文件地址
	Tag    string    `json:"tag" gorm:"comment:文件标签"`                 // 文件标签
	Key    string    `json:"key" gorm:"comment:编号"`                   // 编号
	PostId string    `json:"postId" gorm:"comment:文章ID;default:null"` // 文章ID
	Type   string    `json:"type" gorm:"comment:图片类型;default:'post'"` // 图片类型
	Uuid   uuid.UUID `json:"uuid" gorm:"comment:用户ID"`                // 用户ID
}

// UploadConfig 上传配置
type UploadConfig struct {
	OssType           string `mapstructure:"ossType" json:"ossType"`
	*AliyunOssConfig  `mapstructure:"aliyunOssConfig" json:"aliyunOssConfig"`
	*HuaWeiObsConfig  `mapstructure:"huaWeiObsConfig" json:"huaWeiObsConfig"`
	*QiniuConfig      `mapstructure:"qiniuConfig" json:"qiniuConfig"`
	*TencentCOSConfig `mapstructure:"tencentCOSConfig" json:"tencentCOSConfig"`
	*LocalConfig      `mapstructure:"localConfig" json:"localConfig"`
}

// AliyunOssConfig oss配置
type AliyunOssConfig struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint"`
	AccessKeyId     string `mapstructure:"accessKeyId" json:"accessKeyId"`
	AccessKeySecret string `mapstructure:"accessKeySecret" json:"accessKeySecret"`
	BucketName      string `mapstructure:"bucketName" json:"bucketName"`
	BucketPoint     string `mapstructure:"bucketPoint" json:"bucketPoint"`
}

// HuaWeiObsConfig oss配置
type HuaWeiObsConfig struct {
	Path      string `mapstructure:"path" json:"path"`
	Bucket    string `mapstructure:"bucket" json:"bucket"`
	Endpoint  string `mapstructure:"endpoint" json:"endpoint"`
	AccessKey string `mapstructure:"accessKey" json:"accessKey"`
	SecretKey string `mapstructure:"secretKey" json:"secretKey"`
}

// QiniuConfig oss配置
type QiniuConfig struct {
	Zone          string `mapstructure:"zone" json:"zone"`
	Bucket        string `mapstructure:"bucket" json:"bucket"`
	ImgPath       string `mapstructure:"imgPath" json:"imgPath"`
	UseHTTPS      bool   `mapstructure:"useHttps" json:"useHttps"`
	AccessKey     string `mapstructure:"accessKey" json:"accessKey"`
	SecretKey     string `mapstructure:"secretKey" json:"secretKey"`
	UseCdnDomains bool   `mapstructure:"useCdnDomains" json:"useCdnDomains"`
}

// TencentCOSConfig oss配置
type TencentCOSConfig struct {
	Bucket     string `mapstructure:"bucket" json:"bucket"`
	Region     string `mapstructure:"region" json:"region"`
	SecretID   string `mapstructure:"secretId" json:"secretId"`
	SecretKey  string `mapstructure:"secretKey" json:"secretKey"`
	BaseURL    string `mapstructure:"baseUrl" json:"baseUrl"`
	PathPrefix string `mapstructure:"pathPrefix" json:"pathPrefix"`
}

// LocalConfig oss配置
type LocalConfig struct {
	Path string `mapstructure:"path" json:"path"`
}
