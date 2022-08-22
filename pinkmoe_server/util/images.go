package util

import (
	"fmt"
	"github.com/h2non/bimg"
	"os"
)

//调整图像大小
func Resize(path string, width, height int) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Resize(width, height)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	_, err = bimg.NewImage(newImage).Size()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("001-resize.jpg", newImage)
}

//指定长宽缩略图像
func SmartCrop(path string, width, height int) (img []byte) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).SmartCrop(width, height)

	img = bimg.NewImage(newImage).Image()
	return img
}

//格式转换
func convert(path string, imageType bimg.ImageType) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Convert(imageType)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("004-convert."+bimg.ImageTypes[imageType], newImage)

}

// 文字水印
func watermarkText(path string) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// 水印的相关数据
	watermark := bimg.Watermark{
		Text:       "GoCN 社区",
		DPI:        150,
		Margin:     300,
		Opacity:    0.25, // 不透明度
		Width:      500,
		Font:       "sans bold 14",
		Background: bimg.Color{255, 255, 255},
	}

	newImage, err := bimg.NewImage(buffer).Watermark(watermark)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("005-watermark.jpg", newImage)
}

// 图片水印
func watermarkImage(path, logo string) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	logoBuffer, err := bimg.Read(logo)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// 水印的相关数据
	watermark := bimg.WatermarkImage{
		Left:    100,
		Top:     200,
		Buf:     logoBuffer,
		Opacity: 1, // 不透明度 0~1 之间的浮点数
	}

	newImage, err := bimg.NewImage(buffer).WatermarkImage(watermark)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("006-watermark-image.jpg", newImage)
}
