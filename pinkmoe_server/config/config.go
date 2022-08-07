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
	Name          string `mapstructure:"name"`
	Mode          string `mapstructure:"mode"`
	Version       string `mapstructure:"version"`
	VideoSize     int64  `mapstructure:"video_size"`
	PicSize       int64  `mapstructure:"pic_size"`
	StartTime     string `mapstructure:"start_time"`
	MachineId     int64  `mapstructure:"machine_id"`
	RateLimitTime int64  `mapstructure:"rate_limit_time"`
	RateLimitNum  int64  `mapstructure:"rate_limit_num"`
	Port          int    `mapstructure:"port"`
	OssType       string `mapstructure:"oss_type"`

	*Casbin           `mapstructure:"casbin"`
	*Captcha          `mapstructure:"captcha"`
	*AuthConfig       `mapstructure:"auth"`
	*LogConfig        `mapstructure:"log"`
	*MySqlConfig      `mapstructure:"mysql"`
	*RedisConfig      `mapstructure:"redis"`
	*AliyunOssConfig  `mapstructure:"aliyun_oss"`
	*HuaWeiObsConfig  `mapstructure:"hua_wei_obs"`
	*QiniuConfig      `mapstructure:"qiniu"`
	*TencentCOSConfig `mapstructure:"tencent_cos"`
	*LocalConfig      `mapstructure:"local"`
	*EmailConfig      `mapstructure:"email"`
}

// Casbin jwt配置
type Casbin struct {
	ModelPath string `mapstructure:"model_path"`
}

// Captcha jwt配置
type Captcha struct {
	ImgHeight int `mapstructure:"img_height"`
	ImgWidth  int `mapstructure:"img_width"`
	KeyLong   int `mapstructure:"key_long"`
}

// AuthConfig jwt配置
type AuthConfig struct {
	JwtExpire  int    `mapstructure:"jwt_expire"`
	AuthSecret string `mapstructure:"auth_secret"`
	Issuer     string `mapstructure:"issuer"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level             string `mapstructure:"level"`
	Filename          string `mapstructure:"filename"`
	UserClickFilename string `mapstructure:"userClickFileName"`
	MaxSize           int    `mapstructure:"max_size"`
	MaxAge            int    `mapstructure:"max_age"`
	MaxBackups        int    `mapstructure:"max_backups"`
}

// MySqlConfig mysql配置
type MySqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	Config       string `mapstructure:"config"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// AliyunOssConfig oss配置
type AliyunOssConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	BucketName      string `mapstructure:"bucket_name"`
	BucketPoint     string `mapstructure:"bucket_point"`
}

// HuaWeiObsConfig oss配置
type HuaWeiObsConfig struct {
	Path      string `mapstructure:"path"`
	Bucket    string `mapstructure:"bucket"`
	Endpoint  string `mapstructure:"endpoint"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}

// QiniuConfig oss配置
type QiniuConfig struct {
	Zone          string `mapstructure:"zone"`
	Bucket        string `mapstructure:"bucket"`
	ImgPath       string `mapstructure:"img_path"`
	UseHTTPS      bool   `mapstructure:"use_https"`
	AccessKey     string `mapstructure:"access_key"`
	SecretKey     string `mapstructure:"secret_key"`
	UseCdnDomains bool   `mapstructure:"use_cdn_domains"`
}

// TencentCOSConfig oss配置
type TencentCOSConfig struct {
	Bucket     string `mapstructure:"bucket"`
	Region     string `mapstructure:"region"`
	SecretID   string `mapstructure:"secret_id"`
	SecretKey  string `mapstructure:"secret_key"`
	BaseURL    string `mapstructure:"base_url"`
	PathPrefix string `mapstructure:"path_prefix"`
}

// LocalConfig oss配置
type LocalConfig struct {
	Path string `mapstructure:"path"`
}

// EmailConfig oss配置
type EmailConfig struct {
	User       string `mapstructure:"user"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	IsSSL      bool   `mapstructure:"is_ssl"`
	ReplyEmail string `mapstructure:"reply_email"`
}
