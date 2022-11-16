package response

// EmailConfig oss配置
type EmailConfig struct {
	User       string `json:"user"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	IsSSL      bool   `json:"isSsl"`
	ReplyEmail string `json:"replyEmail"`
}
