package cms

import (
	"github.com/gin-gonic/gin"
	"github.com/h2non/bimg"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"server/controller"
	"server/model/request"
	"server/model/response"
)

func init() {
	controller.Register(&Images{})
}

type Images struct{}

func (images *Images) ImageThumbnailGet(c *gin.Context) {
	var p request.ImageParams
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("request.ImageParams with invalid param", zap.Error(err))
		response.FailWithMessage(response.CodeInvalidParam.Msg(), c)
		return
	}
	res, _ := http.Get(p.Image)
	data, _ := ioutil.ReadAll(res.Body)
	newImage, _ := bimg.NewImage(data).SmartCrop(p.Width, p.Height)
	img := bimg.NewImage(newImage).Image()
	c.Writer.WriteString(string(img))
}
