/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:02:44
 * @FilePath: /pinkmoe_server/config/config.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package config

// AppConfig 全局配置
type AppConfig struct {
	*BasicConfig  `mapstructure:"basicConfig" json:"basicConfig"`
	*UploadConfig `mapstructure:"uploadConfig" json:"uploadConfig"`
	*MySqlConfig  `mapstructure:"mySqlConfig" json:"mySqlConfig"`
	*RedisConfig  `mapstructure:"redisConfig" json:"redisConfig"`
	*EmailConfig  `mapstructure:"emailConfig" json:"emailConfig"`
}

// BasicConfig 上传配置
type BasicConfig struct {
	Name          string `mapstructure:"name" json:"name"`
	Mode          string `mapstructure:"mode" json:"mode"`
	Host          string `mapstructure:"host" json:"host"`
	AdminHost     string `mapstructure:"adminHost" json:"adminHost"`
	Version       string `mapstructure:"version" json:"version"`
	VideoSize     int64  `mapstructure:"videoSize" json:"videoSize"`
	PicSize       int64  `mapstructure:"picSize" json:"picSize"`
	StartTime     string `mapstructure:"startTime" json:"startTime"`
	MachineId     int64  `mapstructure:"machineId" json:"machineId"`
	RateLimitTime int64  `mapstructure:"rateLimitTime" json:"rateLimitTime"`
	RateLimitNum  int64  `mapstructure:"rateLimitNum" json:"rateLimitNum"`
	Port          int    `mapstructure:"port" json:"port"`
	*Casbin       `mapstructure:"casbin" json:"casbin"`
	*Captcha      `mapstructure:"captcha" json:"captcha"`
	*AuthConfig   `mapstructure:"authConfig" json:"authConfig"`
	*LogConfig    `mapstructure:"logConfig" json:"logConfig"`
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

// Casbin jwt配置
type Casbin struct {
	ModelPath string `mapstructure:"modelPath" json:"modelPath"`
}

// Captcha jwt配置
type Captcha struct {
	ImgHeight int `mapstructure:"imgHeight" json:"imgHeight"`
	ImgWidth  int `mapstructure:"imgWidth" json:"imgWidth"`
	KeyLong   int `mapstructure:"keyLong" json:"keyLong"`
}

// AuthConfig jwt配置
type AuthConfig struct {
	JwtExpire  int    `mapstructure:"jwtExpire" json:"jwtExpire"`
	AuthSecret string `mapstructure:"authSecret" json:"authSecret"`
	Issuer     string `mapstructure:"issuer" json:"issuer"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level" json:"level"`
	Filename   string `mapstructure:"filename" json:"filename"`
	MaxSize    int    `mapstructure:"maxSize" json:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge" json:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups" json:"maxBackups"`
}

// MySqlConfig mysql配置
type MySqlConfig struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         string `mapstructure:"port" json:"port"`
	User         string `mapstructure:"user" json:"user"`
	Password     string `mapstructure:"password" json:"password"`
	Dbname       string `mapstructure:"dbname" json:"dbname"`
	Config       string `mapstructure:"config" json:"config"`
	MaxIdleConns int    `mapstructure:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns" json:"maxOpenConns"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	Db       int    `mapstructure:"db" json:"db"`
	PoolSize int    `mapstructure:"poolSize" json:"poolSize"`
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

// EmailConfig oss配置
type EmailConfig struct {
	User       string `mapstructure:"user" json:"user"`
	Username   string `mapstructure:"username" json:"username"`
	Password   string `mapstructure:"password" json:"password"`
	Host       string `mapstructure:"host" json:"host"`
	Port       int    `mapstructure:"port" json:"port"`
	IsSSL      bool   `mapstructure:"isSsl" json:"isSsl"`
	ReplyEmail string `mapstructure:"replyEmail" json:"replyEmail"`
}
